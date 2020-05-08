package storage

import (
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/variant"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"io"
	"time"
)

type Storage struct {
	mSecretKey []byte

	mKVStoragePath          string
	mKVStorage              iface.IKVStorage
	mKVStorageConfiguration interface{}

	mIndexStoragePath          string
	mIndexStorage              iface.IIndexStorage
	mIndexStorageConfiguration interface{}

	mConfiguration iface.IStorageConfiguration
}

func (s *Storage) SetConfiguration(configuration iface.IStorageConfiguration) bool {
	if s.mConfiguration != nil {
		return false
	}

	s.mConfiguration = configuration

	if s.mConfiguration.StoragePluginKV() == nil {
		return false
	}

	if s.mConfiguration.StoragePluginIndex() == nil {
		return false
	}

	kvStoragePath := s.mConfiguration.StoragePath() + "/kv"
	indexStoragePath := s.mConfiguration.StoragePath() + "/index"

	if !utility.MkPath(kvStoragePath) {
		return false
	}
	if !utility.MkPath(indexStoragePath) {
		return false
	}

	s.mSecretKey = s.mConfiguration.StorageSecretKey()

	s.mKVStorage = s.mConfiguration.StoragePluginKV()
	s.mKVStoragePath = kvStoragePath
	s.mKVStorageConfiguration = s.mConfiguration.KVStorageCustomConfiguration()

	s.mIndexStorage = s.mConfiguration.StoragePluginIndex()
	s.mIndexStoragePath = indexStoragePath
	s.mIndexStorageConfiguration = s.mConfiguration.IndexStorageCustomConfiguration()

	return true
}

func (s *Storage) Open() error {
	if s.mConfiguration == nil {
		return x.ErrInvalidConfiguration
	}

	err1 := s.mKVStorage.Open(
		s.mKVStoragePath,
		s.mSecretKey,
		false,
		s.mKVStorageConfiguration)
	if err1 != nil {
		return err1
	}

	err2 := s.mIndexStorage.Open(
		s.mIndexStoragePath,
		s.mSecretKey,
		s.mIndexStorageConfiguration)
	if err2 != nil {
		return err2
	}

	return nil
}

func (s *Storage) ReadOnlyOpen() error {
	if s.mConfiguration == nil {
		return x.ErrInvalidConfiguration
	}

	return s.mKVStorage.Open(s.mKVStoragePath, s.mSecretKey, true, s.mKVStorageConfiguration)
}

func (s *Storage) Close() error {
	return s.mKVStorage.Close()
}

func (s *Storage) RunGC() {
	s.mKVStorage.RunGC()
}

func (s *Storage) ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) error {
	return s.mKVStorage.ChangeSecretKey(oldSecretKey, newSecretKey)
}

func (s *Storage) PrepareSnapshot() (interface{}, error) {
	return s.mKVStorage.PrepareSnapshot()
}

func (s *Storage) RecoverFromSnapshot(r io.Reader) error {
	return s.mKVStorage.RecoverFromSnapshot(r)
}

func (s *Storage) SaveSnapshot(snapshotContext interface{}, w io.Writer) error {
	return s.mKVStorage.SaveSnapshot(snapshotContext, w)
}

func (s *Storage) SaveAppliedIndex(u uint64) error {
	return s.mKVStorage.Create(
		[]byte(constant.AppliedIndexNamespace),
		[]byte(constant.AppliedIndexKey),
		uidutil.Uint64ToByteSlice(u))
}

func (s *Storage) QueryAppliedIndex() (uint64, error) {
	data, err := s.mKVStorage.Read(
		[]byte(constant.AppliedIndexNamespace),
		[]byte(constant.AppliedIndexKey))

	if err == x.ErrUidDoesNotExists {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}
	return uidutil.ByteSliceToUint64(data), nil
}

func (s *Storage) Lookup(input interface{}, checkNamespaceValidity bool) (interface{}, error) {
	if v, ok := input.([]byte); ok {
		e := &pb.FlameEntry{}
		if err := proto.Unmarshal(v, e); err != nil {
			return nil, x.ErrInvalidLookupInput
		}
		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(e.Namespace) {
				return nil, nil
			}
		}
		return s.mKVStorage.Read(e.Namespace, e.Key)
	}

	if v, ok := input.(*pb.FlameEntry); ok {
		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(v.Namespace) {
				return nil, nil
			}
		}
		return s.mKVStorage.Read(v.Namespace, v.Key)
	}

	if v, ok := input.(pb.FlameEntry); ok {
		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(v.Namespace) {
				return nil, nil
			}
		}

		return s.mKVStorage.Read(v.Namespace, v.Key)
	}

	return nil, x.ErrInvalidLookupInput
}

func (s *Storage) ApplyProposal(pp *pb.FlameProposal, checkNamespaceValidity bool) error {
	if pp.FlameProposalType == pb.FlameProposal_BATCH_ACTION {
		batchAction := &pb.FlameBatchAction{}
		if err := proto.Unmarshal(pp.FlameProposalData, batchAction); err != nil {
			return x.ErrDataUnmarshalError
		}

		if checkNamespaceValidity {
			for idx := range batchAction.FlameActionList {
				if !utility.IsNamespaceValid(batchAction.FlameActionList[idx].FlameEntry.Namespace) {
					return x.ErrInvalidNamespace
				}
			}
		}

		err := s.mKVStorage.ApplyBatchAction(batchAction)
		if err != nil {
			return err
		}

		_ = s.directIndex(batchAction)
	} else if pp.FlameProposalType == pb.FlameProposal_CREATE_INDEX_META {
		indexMeta := &pb.FlameIndexMeta{}
		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
			return x.ErrDataUnmarshalError
		}

		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(indexMeta.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		if err := s.mKVStorage.CreateIndexMeta(indexMeta); err != nil {
			internalLogger.Error("CreateIndexMeta error", zap.Error(err))
			return err
		} else {
			if err := s.mIndexStorage.CreateIndexMeta(indexMeta); err != nil {
				internalLogger.Error("IndexStorage CreateIndexMeta error", zap.Error(err))
				return err
			} else {
				return nil
			}
		}
	} else if pp.FlameProposalType == pb.FlameProposal_UPDATE_INDEX_META {
		indexMeta := &pb.FlameIndexMeta{}
		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
			return x.ErrDataUnmarshalError
		}

		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(indexMeta.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		if err := s.mKVStorage.UpdateIndexMeta(indexMeta); err != nil {
			internalLogger.Error("UpdateIndexMeta error", zap.Error(err))
			return err
		} else {
			if err := s.mIndexStorage.UpdateIndexMeta(indexMeta); err != nil {
				internalLogger.Error("IndexStorage UpdateIndexMeta error", zap.Error(err))
				return err
			} else {
				return nil
			}
		}
	} else if pp.FlameProposalType == pb.FlameProposal_DELETE_INDEX_META {
		indexMeta := &pb.FlameIndexMeta{}
		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
			return x.ErrDataUnmarshalError
		}

		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(indexMeta.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		if err := s.mKVStorage.DeleteIndexMeta(indexMeta); err != nil {
			internalLogger.Error("DeleteIndexMeta error", zap.Error(err))
			return err
		} else {
			if err := s.mIndexStorage.DeleteIndexMeta(indexMeta); err != nil {
				internalLogger.Error("IndexStorage DeleteIndexMeta error", zap.Error(err))
				return err
			} else {
				return nil
			}
		}
	} else if pp.FlameProposalType == pb.FlameProposal_CREATE_ACCESS_CONTROL {
		ac := &pb.FlameAccessControl{}
		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
			return x.ErrDataUnmarshalError
		}

		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(ac.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.mKVStorage.CreateAccessControl(ac)
	} else if pp.FlameProposalType == pb.FlameProposal_UPDATE_ACCESS_CONTROL {
		ac := &pb.FlameAccessControl{}
		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
			return x.ErrDataUnmarshalError
		}

		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(ac.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.mKVStorage.UpdateAccessControl(ac)
	} else if pp.FlameProposalType == pb.FlameProposal_DELETE_ACCESS_CONTROL {
		ac := &pb.FlameAccessControl{}
		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
			return x.ErrDataUnmarshalError
		}

		if checkNamespaceValidity {
			if !utility.IsNamespaceValid(ac.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.mKVStorage.DeleteAccessControl(ac)
	} else if pp.FlameProposalType == pb.FlameProposal_CREATE_USER {
		user := &pb.FlameUser{}
		if err := proto.Unmarshal(pp.FlameProposalData, user); err != nil {
			return x.ErrDataUnmarshalError
		}

		return s.mKVStorage.CreateUser(user)
	} else if pp.FlameProposalType == pb.FlameProposal_UPDATE_USER {
		user := &pb.FlameUser{}
		if err := proto.Unmarshal(pp.FlameProposalData, user); err != nil {
			return x.ErrDataUnmarshalError
		}

		return s.mKVStorage.UpdateUser(user)
	} else if pp.FlameProposalType == pb.FlameProposal_DELETE_USER {
		user := &pb.FlameUser{}
		if err := proto.Unmarshal(pp.FlameProposalData, user); err != nil {
			return x.ErrDataUnmarshalError
		}

		return s.mKVStorage.DeleteUser(user)
	}

	return x.ErrFailedToApplyProposal
}

func (s *Storage) getIndexHolderMap(batchAction *pb.FlameBatchAction) map[string]*internalIndexHolder {
	var indexHolderMap = make(map[string]*internalIndexHolder)

	for idx := range batchAction.FlameActionList {
		flameAction := batchAction.FlameActionList[idx]
		if flameAction == nil {
			continue
		}

		if flameAction.FlameEntry == nil {
			continue
		}

		currentIndexHolder, ok := indexHolderMap[string(flameAction.FlameEntry.Namespace)]
		if !ok {
			currentIndexHolder = &internalIndexHolder{namespace: string(flameAction.FlameEntry.Namespace)}
			currentIndexHolder.indexData = make([]*variant.IndexData, 0, s.mConfiguration.CacheSize())
			indexHolderMap[string(flameAction.FlameEntry.Namespace)] = currentIndexHolder
		}

		data := s.mConfiguration.IndexObject(flameAction.FlameEntry.Namespace, flameAction.FlameEntry.Value)

		if data == nil {
			continue
		}

		id := uidutil.GetUidString(flameAction.FlameEntry.Namespace, flameAction.FlameEntry.Key)

		if flameAction.FlameActionType == pb.FlameAction_CREATE {
			currentIndexHolder.indexData = append(currentIndexHolder.indexData, &variant.IndexData{
				ID:     id,
				Action: variant.CREATE,
				Data:   data,
			})
		}

		if flameAction.FlameActionType == pb.FlameAction_UPDATE {
			currentIndexHolder.indexData = append(currentIndexHolder.indexData, &variant.IndexData{
				ID:     id,
				Action: variant.UPDATE,
				Data:   data,
			})
		}

		if flameAction.FlameActionType == pb.FlameAction_DELETE {
			currentIndexHolder.indexData = append(currentIndexHolder.indexData, &variant.IndexData{
				ID:     id,
				Action: variant.DELETE,
				Data:   data,
			})
		}
	}

	return indexHolderMap
}

func (s *Storage) directIndex(batchAction *pb.FlameBatchAction) error {
	if batchAction == nil {
		return nil
	}

	for k, v := range s.getIndexHolderMap(batchAction) {
		if !s.mIndexStorage.CanIndex(k) {
			flameMeta := &pb.FlameIndexMeta{
				Namespace: []byte(k),
				Version:   1,
				Enabled:   true,
				Default:   true,
				CreatedAt: uint64(time.Now().UnixNano()),
				UpdatedAt: uint64(time.Now().UnixNano()),
			}
			err := s.mIndexStorage.CreateIndexMeta(flameMeta)
			internalLogger.Error("CreateIndexMeta failure",
				zap.Error(err),
				zap.String("namespace", k))
		}

		err := s.mIndexStorage.ApplyIndex(k, v.indexData)

		internalLogger.Error("ApplyIndex failure",
			zap.Error(err),
			zap.String("namespace", k))
	}

	return nil
}

type internalIndexHolder struct {
	namespace string
	indexData []*variant.IndexData
}
