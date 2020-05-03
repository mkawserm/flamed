package storage

import (
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"io"
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

func (s *Storage) IsExists(namespace []byte, key []byte) bool {
	return s.mKVStorage.IsExists(namespace, key)
}

func (s *Storage) Read(namespace []byte, key []byte) ([]byte, error) {
	d, err := s.mKVStorage.Read(namespace, key)
	if err == x.ErrUidDoesNotExists {
		return nil, nil
	}
	return d, err
}

func (s *Storage) Delete(namespace []byte, key []byte) error {
	return s.mKVStorage.Delete(namespace, key)
}

func (s *Storage) Create(namespace []byte, key []byte, value []byte) error {
	return s.mKVStorage.Create(namespace, key, value)
}

func (s *Storage) Update(namespace []byte, key []byte, value []byte) error {
	return s.mKVStorage.Update(namespace, key, value)
}

func (s *Storage) Append(namespace []byte, key []byte, value []byte) error {
	return s.mKVStorage.Append(namespace, key, value)
}

func (s *Storage) ApplyBatchAction(batch *pb.FlameBatchAction) error {
	return s.mKVStorage.ApplyBatchAction(batch)
}

func (s *Storage) ApplyAction(action *pb.FlameAction) error {
	return s.mKVStorage.ApplyAction(action)
}

func (s *Storage) ReadBatch(batch *pb.FlameBatchRead) error {
	return s.mKVStorage.ReadBatch(batch)
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
	return s.mKVStorage.SaveAppliedIndex(u)
}

func (s *Storage) QueryAppliedIndex() (uint64, error) {
	return s.mKVStorage.QueryAppliedIndex()
}

func (s *Storage) CreateIndexMeta(meta *pb.FlameIndexMeta) error {
	return s.mKVStorage.CreateIndexMeta(meta)
}

func (s *Storage) GetIndexMeta(meta *pb.FlameIndexMeta) error {
	return s.mKVStorage.GetIndexMeta(meta)
}

func (s *Storage) GetAllIndexMeta() ([]*pb.FlameIndexMeta, error) {
	return s.mKVStorage.GetAllIndexMeta()
}

func (s *Storage) UpdateIndexMeta(meta *pb.FlameIndexMeta) error {
	return s.mKVStorage.UpdateIndexMeta(meta)
}

func (s *Storage) DeleteIndexMeta(meta *pb.FlameIndexMeta) error {
	return s.mKVStorage.DeleteIndexMeta(meta)
}

func (s *Storage) CreateUser(user *pb.FlameUser) error {
	return s.mKVStorage.CreateUser(user)
}

func (s *Storage) GetUser(user *pb.FlameUser) error {
	return s.mKVStorage.GetUser(user)
}

func (s *Storage) GetAllUser() ([]*pb.FlameUser, error) {
	return s.mKVStorage.GetAllUser()
}

func (s *Storage) UpdateUser(user *pb.FlameUser) error {
	return s.mKVStorage.UpdateUser(user)
}

func (s *Storage) DeleteUser(user *pb.FlameUser) error {
	return s.mKVStorage.DeleteUser(user)
}

func (s *Storage) CreateAccessControl(ac *pb.FlameAccessControl) error {
	return s.mKVStorage.CreateAccessControl(ac)
}

func (s *Storage) GetAccessControl(ac *pb.FlameAccessControl) error {
	return s.mKVStorage.GetAccessControl(ac)
}

func (s *Storage) GetAllAccessControl() ([]*pb.FlameAccessControl, error) {
	return s.mKVStorage.GetAllAccessControl()
}

func (s *Storage) UpdateAccessControl(ac *pb.FlameAccessControl) error {
	return s.mKVStorage.UpdateAccessControl(ac)
}

func (s *Storage) DeleteAccessControl(ac *pb.FlameAccessControl) error {
	return s.mKVStorage.DeleteAccessControl(ac)
}

func (s *Storage) Lookup(input interface{}, checkValidity bool) (interface{}, error) {
	if v, ok := input.([]byte); ok {
		e := &pb.FlameEntry{}
		if err := proto.Unmarshal(v, e); err != nil {
			return nil, x.ErrInvalidLookupInput
		}
		if checkValidity {
			if !utility.IsNamespaceValid(e.Namespace) {
				return nil, nil
			}
		}
		return s.Read(e.Namespace, e.Key)
	}

	if v, ok := input.(*pb.FlameEntry); ok {
		if checkValidity {
			if !utility.IsNamespaceValid(v.Namespace) {
				return nil, nil
			}
		}
		return s.Read(v.Namespace, v.Key)
	}

	if v, ok := input.(pb.FlameEntry); ok {
		if checkValidity {
			if !utility.IsNamespaceValid(v.Namespace) {
				return nil, nil
			}
		}

		return s.Read(v.Namespace, v.Key)
	}

	return nil, x.ErrInvalidLookupInput
}

func (s *Storage) ApplyProposal(pp *pb.FlameProposal, checkValidity bool) error {
	if pp.FlameProposalType == pb.FlameProposal_BATCH_ACTION {
		batchAction := &pb.FlameBatchAction{}
		if err := proto.Unmarshal(pp.FlameProposalData, batchAction); err != nil {
			return x.ErrFailedToApplyProposal
		}

		if checkValidity {
			for idx := range batchAction.FlameActionList {
				if !utility.IsNamespaceValid(batchAction.FlameActionList[idx].FlameEntry.Namespace) {
					return x.ErrInvalidNamespace
				}
			}
		}

		return s.ApplyBatchAction(batchAction)
	} else if pp.FlameProposalType == pb.FlameProposal_CREATE_INDEX_META {
		indexMeta := &pb.FlameIndexMeta{}
		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
			return x.ErrFailedToApplyProposal
		}

		if checkValidity {
			if !utility.IsNamespaceValid(indexMeta.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.CreateIndexMeta(indexMeta)
	} else if pp.FlameProposalType == pb.FlameProposal_UPDATE_INDEX_META {
		indexMeta := &pb.FlameIndexMeta{}
		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
			return x.ErrFailedToApplyProposal
		}

		if checkValidity {
			if !utility.IsNamespaceValid(indexMeta.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.UpdateIndexMeta(indexMeta)
	} else if pp.FlameProposalType == pb.FlameProposal_DELETE_INDEX_META {
		indexMeta := &pb.FlameIndexMeta{}
		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
			return x.ErrFailedToApplyProposal
		}

		if checkValidity {
			if !utility.IsNamespaceValid(indexMeta.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.DeleteIndexMeta(indexMeta)
	} else if pp.FlameProposalType == pb.FlameProposal_CREATE_ACCESS_CONTROL {
		ac := &pb.FlameAccessControl{}
		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
			return x.ErrFailedToApplyProposal
		}

		if checkValidity {
			if !utility.IsNamespaceValid(ac.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.CreateAccessControl(ac)
	} else if pp.FlameProposalType == pb.FlameProposal_UPDATE_ACCESS_CONTROL {
		ac := &pb.FlameAccessControl{}
		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
			return x.ErrFailedToApplyProposal
		}

		if checkValidity {
			if !utility.IsNamespaceValid(ac.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.UpdateAccessControl(ac)
	} else if pp.FlameProposalType == pb.FlameProposal_DELETE_ACCESS_CONTROL {
		ac := &pb.FlameAccessControl{}
		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
			return x.ErrFailedToApplyProposal
		}

		if checkValidity {
			if !utility.IsNamespaceValid(ac.Namespace) {
				return x.ErrInvalidNamespace
			}
		}

		return s.DeleteAccessControl(ac)
	}

	return x.ErrFailedToApplyProposal
}
