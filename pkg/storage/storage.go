package storage

import (
	"context"
	badgerDb "github.com/dgraph-io/badger/v2"
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

type IndexDataContainer map[string][]*variant.IndexData
type IndexMetaActionContainer map[string][]*variant.IndexMetaAction

type StateContext struct {
	mReadOnly            bool
	mStorage             *Storage
	mIndexDataList       []*variant.IndexData
	mIndexMetaActionList []*variant.IndexMetaAction
	mTxn                 iface.IStateStorageTransaction
}

func (s *StateContext) GetForwardIterator() iface.IStateIterator {
	return s.mTxn.ForwardIterator()
}

func (s *StateContext) GetReverseIterator() iface.IStateIterator {
	return s.mTxn.ReverseIterator()
}

func (s *StateContext) GetKeyOnlyForwardIterator() iface.IStateIterator {
	return s.mTxn.KeyOnlyForwardIterator()
}

func (s *StateContext) GetKeyOnlyReverseIterator() iface.IStateIterator {
	return s.mTxn.KeyOnlyReverseIterator()
}

func (s *StateContext) GetState(key []byte) (*pb.StateEntry, error) {
	if data, err := s.mTxn.Get(key); err == nil {
		entry := &pb.StateEntry{}
		if err := proto.Unmarshal(data, entry); err == nil {
			return entry, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (s *StateContext) UpsertState(key []byte, entry *pb.StateEntry) error {
	if s.mReadOnly {
		return nil
	}

	if data, err := proto.Marshal(entry); err == nil {
		return s.mTxn.Set(key, data)
	} else {
		return err
	}
}

func (s *StateContext) DeleteState(key []byte) error {
	if s.mReadOnly {
		return nil
	}

	return s.mTxn.Delete(key)
}

func (s *StateContext) UpsertIndex(id string, data interface{}) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	s.mIndexDataList = append(s.mIndexDataList, &variant.IndexData{
		ID:     id,
		Data:   data,
		Action: constant.UPSERT,
	})
	return nil
}

func (s *StateContext) DeleteIndex(id string) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	s.mIndexDataList = append(s.mIndexDataList, &variant.IndexData{
		ID:     id,
		Action: constant.DELETE,
	})
	return nil
}

func (s *StateContext) AutoIndexMeta() bool {
	return s.mStorage.AutoIndexMeta()
}

func (s *StateContext) CanIndex(namespace string) bool {
	if !s.mStorage.IndexEnable() {
		return false
	}

	return s.mStorage.CanIndex(namespace)
}

func (s *StateContext) UpsertIndexMeta(meta *pb.IndexMeta) error {
	if s.mReadOnly {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    constant.UPSERT,
		IndexMeta: meta,
	}
	s.mIndexMetaActionList = append(s.mIndexMetaActionList, m)
	return nil
}

func (s *StateContext) DeleteIndexMeta(meta *pb.IndexMeta) error {
	if s.mReadOnly {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    constant.DELETE,
		IndexMeta: meta,
	}
	s.mIndexMetaActionList = append(s.mIndexMetaActionList, m)
	return nil
}

func (s *StateContext) DefaultIndexMeta(namespace string) error {
	if s.mReadOnly {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    constant.DEFAULT,
		IndexMeta: &pb.IndexMeta{Namespace: []byte(namespace)},
	}
	s.mIndexMetaActionList = append(s.mIndexMetaActionList, m)
	return nil
}

//func (s *StateContext) ApplyIndex(namespace string, data []*variant.IndexData) error {
//	if s.mReadOnly {
//		return nil
//	}
//
//	return s.mStorage.ApplyIndex(namespace, data)
//}

type Storage struct {
	mStateStoragePath          string
	mStateStorageSecretKey     []byte
	mStateStorage              iface.IStateStorage
	mStateStorageConfiguration interface{}

	mIndexStoragePath          string
	mIndexStorageSecretKey     []byte
	mIndexStorage              iface.IIndexStorage
	mIndexStorageConfiguration interface{}

	mConfiguration iface.IStorageConfiguration
}

func (s *Storage) SetConfiguration(configuration iface.IStorageConfiguration) bool {
	if s.mConfiguration != nil {
		return false
	}

	s.mConfiguration = configuration

	if s.mConfiguration.StoragePluginState() == nil {
		return false
	}

	if s.mConfiguration.IndexEnable() {
		if s.mConfiguration.StoragePluginIndex() == nil {
			return false
		}
	}

	stateStoragePath := s.mConfiguration.StateStoragePath()
	if !utility.MkPath(stateStoragePath) {
		return false
	}

	s.mStateStoragePath = stateStoragePath
	s.mStateStorage = s.mConfiguration.StoragePluginState()
	s.mStateStorageSecretKey = s.mConfiguration.StateStorageSecretKey()
	s.mStateStorageConfiguration = s.mConfiguration.StateStorageCustomConfiguration()

	if s.mConfiguration.IndexEnable() {
		indexStoragePath := s.mConfiguration.IndexStoragePath()
		if !utility.MkPath(indexStoragePath) {
			return false
		}
		s.mIndexStoragePath = indexStoragePath
		s.mIndexStorage = s.mConfiguration.StoragePluginIndex()
		s.mIndexStorageSecretKey = s.mConfiguration.IndexStorageSecretKey()
		s.mIndexStorageConfiguration = s.mConfiguration.IndexStorageCustomConfiguration()
	}

	return true
}

func (s *Storage) Open() error {
	if s.mConfiguration == nil {
		return x.ErrInvalidConfiguration
	}

	s.mStateStorage.Setup(s.mStateStoragePath, s.mStateStorageSecretKey, s.mStateStorageConfiguration)

	err1 := s.mStateStorage.Open()
	if err1 != nil {
		return err1
	}

	if s.mConfiguration.IndexEnable() {
		err2 := s.mIndexStorage.Open(
			s.mIndexStoragePath,
			s.mIndexStorageSecretKey,
			s.mIndexStorageConfiguration)
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func (s *Storage) Close() error {
	err1 := s.mStateStorage.Close()

	if err1 != nil {
		return err1
	}

	if s.mConfiguration.IndexEnable() {
		err2 := s.mIndexStorage.Close()
		return err2
	}

	return nil
}

func (s *Storage) RunGC() {
	s.mStateStorage.RunGC()
	if s.mConfiguration.IndexEnable() {
		if s.mIndexStorage != nil {
			s.mIndexStorage.RunGC()
		}
	}
}

func (s *Storage) ChangeSecretKey(path string,
	oldSecretKey []byte,
	newSecretKey []byte,
	encryptionKeyRotationDuration time.Duration) error {

	return s.mStateStorage.ChangeSecretKey(path,
		oldSecretKey,
		newSecretKey,
		encryptionKeyRotationDuration)
}

func (s *Storage) Create(namespace []byte, key []byte, value []byte) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if s.mStateStorage == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()
	if err := txn.Set(uid, value); err != nil {
		internalLogger.Error("set failure", zap.Error(err))
		return x.ErrFailedToCreateDataToStorage
	}

	if err := txn.Commit(); err != nil {
		internalLogger.Error("txn commit failure", zap.Error(err))
		return x.ErrFailedToCreateDataToStorage
	}

	return nil
}

func (s *Storage) Read(namespace []byte, key []byte) ([]byte, error) {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if s.mStateStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)
	txn := s.mStateStorage.NewReadOnlyTransaction()
	defer txn.Discard()

	if val, err := txn.Get(uid); err != nil {
		if err == x.ErrKeyDoesNotExists {
			return nil, x.ErrKeyDoesNotExists
		} else {
			internalLogger.Error("read failure", zap.Error(err))
			return nil, x.ErrFailedToReadDataFromStorage
		}
	} else {
		return val, nil
	}
}

func (s *Storage) Delete(namespace []byte, key []byte) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if s.mStateStorage == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()

	if err := txn.Delete(uid); err != nil {
		internalLogger.Error("deletion failure", zap.Error(err))
		return x.ErrFailedToDeleteDataFromStorage
	}

	if err := txn.Commit(); err != nil {
		internalLogger.Error("txn commit failure", zap.Error(err))
		return x.ErrFailedToCreateDataToStorage
	}

	return nil
}

func (s *Storage) SaveAppliedIndex(u uint64) error {
	return s.Create(
		[]byte(constant.AppliedIndexNamespace),
		[]byte(constant.AppliedIndexKey),
		uidutil.Uint64ToByteSlice(u))
}

func (s *Storage) QueryAppliedIndex() (uint64, error) {
	data, err := s.Read(
		[]byte(constant.AppliedIndexNamespace),
		[]byte(constant.AppliedIndexKey))

	if err == x.ErrKeyDoesNotExists {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}
	return uidutil.ByteSliceToUint64(data), nil
}

func (s *Storage) Lookup(request variant.LookupRequest) (interface{}, error) {
	if len(request.FamilyName) != 0 && len(request.FamilyVersion) != 0 {
		tp := s.mConfiguration.GetTransactionProcessor(request.FamilyName, request.FamilyVersion)
		if tp == nil {
			return nil, x.ErrTPNotFound
		} else {
			readOnlyTxn := s.mStateStorage.NewReadOnlyTransaction()
			defer readOnlyTxn.Discard()
			readOnlyStateContext := &StateContext{
				mReadOnly: true,
				mStorage:  s,
				mTxn:      readOnlyTxn,
			}
			return tp.Lookup(request.Context, readOnlyStateContext, request.Query)
		}
	} else {
		if _, ok := request.Query.(pb.AppliedIndexQuery); ok {
			return s.QueryAppliedIndex()
		}
		return nil, x.ErrTPNotFound
	}
}

func (s *Storage) Search(_ variant.SearchRequest) (interface{}, error) {
	return nil, nil
}

func (s *Storage) ApplyProposal(ctx context.Context, proposal *pb.Proposal, entryIndex uint64) *pb.ProposalResponse {
	defer func() {
		_ = internalLogger.Sync()
	}()

	internalLogger.Info("entry indexmeta", zap.Uint64("entryIndex", entryIndex))
	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()

	var indexDataContainer = make(IndexDataContainer)
	var indexMetaActionContainer = make(IndexMetaActionContainer)

	pr := pb.NewProposalResponse(0)
	pr.Uuid = proposal.Uuid

	for _, t := range proposal.Transactions {
		tp := s.mConfiguration.GetTransactionProcessor(t.FamilyName, t.FamilyVersion)
		if tp == nil {
			pr.Status = 0
			pr.ErrorCode = 0
			pr.ErrorText = "Transaction family name:" + t.FamilyName + " family version:" + t.FamilyVersion + " not found"
			return pr
		}

		stateContext := &StateContext{
			mReadOnly:            false,
			mStorage:             s,
			mTxn:                 txn,
			mIndexDataList:       make([]*variant.IndexData, 0),
			mIndexMetaActionList: make([]*variant.IndexMetaAction, 0),
		}
		tpr := tp.Apply(ctx, stateContext, t)
		pr.Append(tpr)

		if tpr.Status == 0 {
			pr.ErrorText = "proposal rejected"
			return pr
		} else {
			indexDataContainer[string(t.Namespace)] = stateContext.mIndexDataList
			indexMetaActionContainer[string(t.Namespace)] = stateContext.mIndexMetaActionList
		}
	}

	if err := txn.Commit(); err == nil {
		if s.mConfiguration.IndexEnable() {
			//NOTE: update indexmeta meta
			s.updateIndexMetaOfIndexStorage(indexMetaActionContainer)
			//NOTE: indexmeta data
			s.updateIndexOfIndexStorage(indexDataContainer)
		}

		pr.Status = 1
		return pr
	} else {
		indexDataContainer = nil
		indexMetaActionContainer = nil
		pr.Status = 0
		pr.ErrorText = "proposal rejected because of commit failure"
		return pr
	}
}

func (s *Storage) UpsertIndexMeta(meta *pb.IndexMeta) error {
	if s.mIndexStorage == nil {
		return nil
	}

	return s.mIndexStorage.UpsertIndexMeta(meta)
}

func (s *Storage) DeleteIndexMeta(meta *pb.IndexMeta) error {
	if s.mIndexStorage == nil {
		return nil
	}

	return s.mIndexStorage.DeleteIndexMeta(meta)
}

func (s *Storage) CanIndex(namespace string) bool {
	if s.mIndexStorage == nil {
		return false
	}

	return s.mIndexStorage.CanIndex(namespace)
}

func (s *Storage) IndexEnable() bool {
	return s.mConfiguration.IndexEnable()
}

func (s *Storage) AutoIndexMeta() bool {
	return s.mConfiguration.AutoIndexMeta()
}

func (s *Storage) ApplyIndex(namespace string, data []*variant.IndexData) error {
	if s.mIndexStorage == nil {
		return nil
	}
	return s.mIndexStorage.ApplyIndex(namespace, data)
}

func (s *Storage) DefaultIndexMeta(namespace string) error {
	if s.mIndexStorage == nil {
		return nil
	}

	return s.mIndexStorage.DefaultIndexMeta(namespace)
}

func (s *Storage) updateIndexOfIndexStorage(indexDataContainer IndexDataContainer) {
	if len(indexDataContainer) == 0 {
		return
	}

	/* TODO: SEND INDEX DATA TO BE PROCESSED USING GO CHANNEL */
	err := s.directIndex(indexDataContainer)

	if err != nil {
		internalLogger.Error("directIndex error", zap.Error(err))
	}
}

func (s *Storage) updateIndexMetaOfIndexStorage(indexMetaActionContainer IndexMetaActionContainer) {
	if len(indexMetaActionContainer) == 0 {
		return
	}

	for _, v := range indexMetaActionContainer {
		for _, v2 := range v {
			if v2.Action == constant.UPSERT {
				err := s.mIndexStorage.UpsertIndexMeta(v2.IndexMeta)
				if err != nil {
					internalLogger.Error("upsert indexmeta error", zap.Error(err))
				}
			}
			if v2.Action == constant.DELETE {
				err := s.mIndexStorage.DeleteIndexMeta(v2.IndexMeta)
				if err != nil {
					internalLogger.Error("delete indexmeta error", zap.Error(err))
				}
			}

			if v2.Action == constant.DEFAULT {
				err := s.mIndexStorage.DefaultIndexMeta(string(v2.IndexMeta.Namespace))
				if err != nil {
					internalLogger.Error("default indexmeta error", zap.Error(err))
				}
			}
		}
	}
}

func (s *Storage) PrepareSnapshot() (interface{}, error) {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if s.mStateStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	internalLogger.Debug("snapshot prepared")
	return s.mStateStorage.NewTransaction(), nil
}

func (s *Storage) RecoverFromSnapshot(r io.Reader) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	internalLogger.Debug("recovering from snapshot")

	if s.mStateStorage == nil {
		return x.ErrStorageIsNotReady
	}

	sz := make([]byte, 8)
	if _, err := io.ReadFull(r, sz); err != nil {
		internalLogger.Error("read error", zap.Error(err))
		return x.ErrFailedToRecoverFromSnapshot
	}

	total := uidutil.ByteSliceToUint64(sz)

	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()

	for i := uint64(0); i < total; i++ {
		if _, err := io.ReadFull(r, sz); err != nil {
			internalLogger.Error("sm read error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		toRead := uidutil.ByteSliceToUint64(sz)
		data := make([]byte, toRead)
		if _, err := io.ReadFull(r, data); err != nil {
			internalLogger.Error("sm read error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		entry := &pb.StateSnapshot{}
		if err := proto.Unmarshal(data, entry); err != nil {
			internalLogger.Error("sm unmarshal error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		if err := txn.Set(entry.Uid, entry.Data); err == badgerDb.ErrTxnTooBig {
			if err := txn.Commit(); err != nil {
				internalLogger.Error("txn commit error", zap.Error(err))
				return x.ErrFailedToRecoverFromSnapshot
			}

			txn = s.mStateStorage.NewTransaction()

			if err := txn.Set(entry.Uid, entry.Data); err != nil {
				internalLogger.Error("txn set error", zap.Error(err))
				return x.ErrFailedToRecoverFromSnapshot
			}

		} else if err != nil {
			internalLogger.Error("txn set error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}
	}

	if err := txn.Commit(); err != nil {
		internalLogger.Error("txn commit error", zap.Error(err))
		return x.ErrFailedToRecoverFromSnapshot
	}

	internalLogger.Debug("storage recovered from snapshot")
	return nil
}

func (s *Storage) SaveSnapshot(snapshotContext interface{}, w io.Writer) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	internalLogger.Debug("saving snapshot")
	if s.mStateStorage == nil {
		return x.ErrStorageIsNotReady
	}

	if snapshotContext == nil {
		return x.ErrFailedToSaveSnapshot
	}

	var txn iface.IStateStorageTransaction
	if v, ok := snapshotContext.(iface.IStateStorageTransaction); ok {
		txn = v
	} else {
		return x.ErrFailedToSaveSnapshot
	}

	defer txn.Discard()

	total := uint64(0)
	it := txn.KeyOnlyForwardIterator()
	for it.Rewind(); it.Valid(); it.Next() {
		total = total + 1
	}
	it.Close()

	if _, err := w.Write(uidutil.Uint64ToByteSlice(total)); err != nil {
		internalLogger.Error("storage write error", zap.Error(err))
		return x.ErrFailedToSaveSnapshot
	}

	it = txn.ForwardIterator()
	defer it.Close()

	for it.Rewind(); it.Valid(); it.Next() {
		item := it.StateSnapshot()
		if data, err := proto.Marshal(item); err == nil {
			dataLength := uint64(len(data))
			if _, err := w.Write(uidutil.Uint64ToByteSlice(dataLength)); err != nil {
				internalLogger.Error("storage write error", zap.Error(err))
				return x.ErrFailedToSaveSnapshot
			}
			if _, err := w.Write(data); err != nil {
				return x.ErrFailedToSaveSnapshot
			}
		} else {
			internalLogger.Error("state snapshot marshal error", zap.Error(err))
			return x.ErrFailedToSaveSnapshot
		}
	}

	internalLogger.Debug("storage snapshot saved")
	return nil
}

func (s *Storage) directIndex(indexDataContainer IndexDataContainer) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if indexDataContainer == nil {
		return nil
	}

	for k, v := range indexDataContainer {
		if !s.mIndexStorage.CanIndex(k) && s.mConfiguration.AutoIndexMeta() {
			//internalLogger.Info("no indexmeta found, creating new one", zap.String("namespace",k))
			indexMeta := &pb.IndexMeta{
				Namespace: []byte(k),
				Version:   1,
				Enabled:   true,
				Default:   true,
				CreatedAt: uint64(time.Now().UnixNano()),
				UpdatedAt: uint64(time.Now().UnixNano()),
			}
			err := s.mIndexStorage.UpsertIndexMeta(indexMeta)
			if err != nil {
				internalLogger.Error("UpsertIndexMeta failure",
					zap.Error(err),
					zap.String("namespace", k))
			}
		}

		if s.mIndexStorage.CanIndex(k) {
			err := s.mIndexStorage.ApplyIndex(k, v)
			if err != nil {
				internalLogger.Error("ApplyIndex failure",
					zap.Error(err),
					zap.String("namespace", k))
			}
		}
	}

	return nil
}
