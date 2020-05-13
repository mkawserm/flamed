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

type StateContext struct {
	mReadOnly      bool
	mStorage       *Storage
	mIndexDataList []*variant.IndexData
	mTxn           iface.IStateStorageTransaction
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

	return s.mStorage.UpsertIndexMeta(meta)
}

func (s *StateContext) DeleteIndexMeta(meta *pb.IndexMeta) error {
	if s.mReadOnly {
		return nil
	}

	return s.mStorage.DeleteIndexMeta(meta)
}

func (s *StateContext) DefaultIndexMeta(namespace string) error {
	if s.mReadOnly {
		return nil
	}

	return s.mStorage.DefaultIndexMeta(namespace)
}

func (s *StateContext) ApplyIndex(namespace string, data []*variant.IndexData) error {
	if s.mReadOnly {
		return nil
	}

	return s.mStorage.ApplyIndex(namespace, data)
}

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

func (s *Storage) ApplyProposal(ctx context.Context, proposal *pb.Proposal) *pb.ProposalResponse {
	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()

	var indexDataContainer = make(IndexDataContainer)

	pr := pb.NewProposalResponse(0)
	for _, t := range proposal.Transactions {
		tp := s.mConfiguration.GetTransactionProcessor(t.FamilyName, t.FamilyVersion)
		if tp == nil {
			pr.Status = 0
			pr.ErrorCode = 0
			pr.ErrorText = "Transaction family name:" + t.FamilyName + " family version:" + t.FamilyVersion + " not found"
			return pr
		}

		stateContext := &StateContext{
			mReadOnly:      false,
			mStorage:       s,
			mTxn:           txn,
			mIndexDataList: make([]*variant.IndexData, 0),
		}
		tpr := tp.Apply(ctx, stateContext, t)
		pr.Append(tpr)

		if tpr.Status == 0 {
			pr.ErrorText = "proposal rejected"
			return pr
		} else {
			indexDataContainer[string(t.Namespace)] = stateContext.mIndexDataList
		}
	}

	if err := txn.Commit(); err == nil {
		if s.mConfiguration.IndexEnable() {
			if len(indexDataContainer) != 0 {
				/* TODO: SEND INDEX DATA TO BE PROCESSED */

			}
		}

		pr.Status = 1
		return pr
	} else {
		indexDataContainer = nil
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

//func (s *Storage) UpsertIndexMeta(meta *pb.FlameIndexMeta) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	data, err := proto.Marshal(meta)
//	if err != nil {
//		internalLogger.Error("marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	//uid := uidutil.GetUid([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	err = s.Create([]byte(constant.IndexMetaNamespace), meta.Namespace, data)
//
//	if err != nil {
//		internalLogger.Error("create error", zap.Error(err))
//		return x.ErrFailedToCreateIndexMeta
//	}
//
//	return nil
//}

//func (s *Storage) IsIndexMetaExists(meta *pb.FlameIndexMeta) bool {
//	return s.mStateMachineStorage.IsExists([]byte(constant.IndexMetaNamespace), meta.Namespace)
//}

//func (s *Storage) GetIndexMeta(meta *pb.FlameIndexMeta) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	//uid := uidutil.GetUid([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	data, err := s.Read([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	if err != nil {
//		internalLogger.Error("read error", zap.Error(err))
//		return x.ErrFailedToGetIndexMeta
//	}
//
//	err = proto.Unmarshal(data, meta)
//	if err != nil {
//		internalLogger.Error("storage proto unmarshal error", zap.Error(err))
//		return x.ErrFailedToGetIndexMeta
//	}
//
//	return nil
//}

//GetAllIndexMeta() ([]*pb.FlameIndexMeta, error)
//func (s *Storage) UpdateIndexMeta(meta *pb.FlameIndexMeta) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	data, err := proto.Marshal(meta)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	//uid := uidutil.GetUid([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	err = s.Create([]byte(constant.IndexMetaNamespace), meta.Namespace, data)
//
//	if err != nil {
//		internalLogger.Error("update error", zap.Error(err))
//		return x.ErrFailedToUpdateIndexMeta
//	}
//
//	return nil
//}

//func (s *Storage) DeleteIndexMeta(meta *pb.FlameIndexMeta) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	//uid := uidutil.GetUid([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	err := s.mStateMachineStorage.Delete([]byte(constant.IndexMetaNamespace), meta.Namespace)
//
//	if err != nil {
//		internalLogger.Error("sm storage delete error", zap.Error(err))
//		return x.ErrFailedToDeleteIndexMeta
//	}
//
//	return nil
//}

//func (s *Storage) CreateUser(user *pb.FlameUser) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	data, err := proto.Marshal(user)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	//uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(user.Username))
//	err = s.mStateMachineStorage.Create([]byte(constant.UserNamespace), []byte(user.Username), data)
//
//	if err != nil {
//		internalLogger.Error("sm storage create error", zap.Error(err))
//		return x.ErrFailedToCreateUser
//	}
//
//	return nil
//}

//func (s *Storage) IsUserExists(user *pb.FlameUser) bool {
//	return s.mStateMachineStorage.IsExists([]byte(constant.UserNamespace), []byte(user.Username))
//}

//func (s *Storage) GetUser(user *pb.FlameUser) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	//uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(user.Username))
//	data, err := s.mStateMachineStorage.Read([]byte(constant.UserNamespace), []byte(user.Username))
//	if err != nil {
//		internalLogger.Error("sm storage read error", zap.Error(err))
//		return x.ErrFailedToGetUser
//	}
//
//	err = proto.Unmarshal(data, user)
//	if err != nil {
//		internalLogger.Error("unmarshal error", zap.Error(err))
//		return x.ErrFailedToGetUser
//	}
//
//	return nil
//}

//GetAllUser() ([]*pb.FlameUser, error)
//func (s *Storage) UpdateUser(user *pb.FlameUser) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	data, err := proto.Marshal(user)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	//uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(user.Username))
//	err = s.mStateMachineStorage.Update([]byte(constant.UserNamespace), []byte(user.Username), data)
//
//	if err != nil {
//		internalLogger.Error("sm storage update error", zap.Error(err))
//		return x.ErrFailedToUpdateUser
//	}
//
//	return nil
//}

//func (s *Storage) DeleteUser(user *pb.FlameUser) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	//uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(user.Username))
//	err := s.mStateMachineStorage.Delete([]byte(constant.UserNamespace), []byte(user.Username))
//
//	if err != nil {
//		internalLogger.Error("sm storage delete error", zap.Error(err))
//		return x.ErrFailedToDeleteUser
//	}
//
//	return nil
//}

//func (s *Storage) CreateAccessControl(ac *pb.FlameAccessControl) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	data, err := proto.Marshal(ac)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	//uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
//	//	uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//
//	err = s.mStateMachineStorage.Create([]byte(constant.AccessControlNamespace),
//		uidutil.GetUid([]byte(ac.Username), ac.Namespace), data)
//
//	if err != nil {
//		internalLogger.Error("sm storage create error", zap.Error(err))
//		return x.ErrFailedToCreateAccessControl
//	}
//
//	return nil
//}

//func (s *Storage) IsAccessControlExists(ac *pb.FlameAccessControl) bool {
//	return s.mStateMachineStorage.IsExists([]byte(constant.UserNamespace), uidutil.GetUid([]byte(ac.Username), ac.Namespace))
//}

//func (s *Storage) GetAccessControl(ac *pb.FlameAccessControl) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	//uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
//	//	uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//	data, err := s.mStateMachineStorage.Read([]byte(constant.AccessControlNamespace),
//		uidutil.GetUid([]byte(ac.Username), ac.Namespace))
//
//	if err != nil {
//		internalLogger.Error("sm storage read error", zap.Error(err))
//		return x.ErrFailedToGetAccessControl
//	}
//
//	err = proto.Unmarshal(data, ac)
//	if err != nil {
//		internalLogger.Error("unmarshal error", zap.Error(err))
//		return x.ErrFailedToGetAccessControl
//	}
//
//	return nil
//}

//GetAllAccessControl() ([]*pb.FlameAccessControl, error)
//func (s *Storage) UpdateAccessControl(ac *pb.FlameAccessControl) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	data, err := proto.Marshal(ac)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	//uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
//	//	uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//
//	err = s.mStateMachineStorage.Update([]byte(constant.AccessControlNamespace),
//		uidutil.GetUid([]byte(ac.Username), ac.Namespace), data)
//
//	if err != nil {
//		internalLogger.Error("sm storage update error", zap.Error(err))
//		return x.ErrFailedToUpdateAccessControl
//	}
//
//	return nil
//}

//func (s *Storage) DeleteAccessControl(ac *pb.FlameAccessControl) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	//uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
//	//	uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//	err := s.mStateMachineStorage.Delete([]byte(constant.AccessControlNamespace),
//		uidutil.GetUid([]byte(ac.Username), ac.Namespace))
//
//	if err != nil {
//		internalLogger.Error("sm storage delete error", zap.Error(err))
//		return x.ErrFailedToDeleteAccessControl
//	}
//
//	return nil
//}

//func (s *Storage) executeCommand(cmd *Command) error {
//	if cmd.CommandID == SyncCompleteIndexUpdate {
//		return s.FullIndex()
//	}
//
//	if cmd.CommandID == SyncPartialIndexUpdate {
//		if n, ok := cmd.Data.([]byte); ok {
//			return s.UpdateIndex(n)
//		} else {
//			internalLogger.Error("unknown data for SyncPartialIndexUpdate")
//		}
//	}
//
//	if cmd.CommandID == SyncRunGC {
//		s.RunGC()
//	}
//
//	return nil
//}

//func (s *Storage) Lookup(input interface{}, checkNamespaceValidity bool) (interface{}, error) {
//if v, ok := input.(*Iterator); ok {
//	err := s.mStateMachineStorage.Iterate(v.Seek, v.Prefix, v.Limit, v.Receiver)
//	if err != nil {
//		return nil, err
//	}
//	return nil, nil
//}
//
//if v, ok := input.(*pb.FlameBatchRead); ok {
//	err := s.mStateMachineStorage.ReadBatch(v)
//	if err != nil {
//		return nil, err
//	}
//	return nil, nil
//}
//
//if v, ok := input.(*AppliedIndexQuery); ok {
//	o, err := s.QueryAppliedIndex()
//	if err != nil {
//		return nil, err
//	}
//	v.GetAppliedIndex = o
//	return v, nil
//}
//
//if v, ok := input.(*pb.FlameUser); ok {
//	return v, s.GetUser(v)
//}
//
//if v, ok := input.(*pb.FlameIndexMeta); ok {
//	if checkNamespaceValidity {
//		if !utility.IsNamespaceValid(v.Namespace) {
//			return nil, nil
//		}
//	}
//	return v, s.GetIndexMeta(v)
//}
//
//if v, ok := input.(*pb.FlameAccessControl); ok {
//	if checkNamespaceValidity {
//		if !utility.IsNamespaceValid(v.Namespace) {
//			return nil, nil
//		}
//	}
//	return v, s.GetAccessControl(v)
//}
//
//if v, ok := input.(*pb.FlameEntry); ok {
//	if checkNamespaceValidity {
//		if !utility.IsNamespaceValid(v.Namespace) {
//			return nil, nil
//		}
//	}
//	return v, s.GetFlameEntry(v)
//}
//
//if v, ok := input.(*Command); ok {
//	return nil, s.executeCommand(v)
//}

//	return nil, x.ErrInvalidLookupInput
//}

//func (s *Storage) ApplyProposal(pp *pb.FlameProposal, checkNamespaceValidity bool) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if pp.FlameProposalType == pb.FlameProposal_BATCH_ACTION {
//		batchAction := &pb.FlameBatchAction{}
//		if err := proto.Unmarshal(pp.FlameProposalData, batchAction); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		if checkNamespaceValidity {
//			for idx := range batchAction.FlameActionList {
//				if !utility.IsNamespaceValid(batchAction.FlameActionList[idx].FlameEntry.Namespace) {
//					return x.ErrInvalidNamespace
//				}
//			}
//		}
//
//		err := s.mStateMachineStorage.ApplyBatchAction(batchAction)
//		if err != nil {
//			return err
//		}
//
//		if !s.mConfiguration.IndexEnable() {
//			return nil
//		}
//		err = s.directIndex(batchAction)
//		if err != nil {
//			internalLogger.Error("batch action direct index error", zap.Error(err))
//		}
//
//		return nil
//	} else if pp.FlameProposalType == pb.FlameProposal_CREATE_INDEX_META {
//		indexMeta := &pb.FlameIndexMeta{}
//		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		if checkNamespaceValidity {
//			if !utility.IsNamespaceValid(indexMeta.Namespace) {
//				return x.ErrInvalidNamespace
//			}
//		}
//
//		if err := s.UpsertIndexMeta(indexMeta); err != nil {
//			internalLogger.Error("UpsertIndexMeta error", zap.Error(err))
//			return err
//		} else {
//			if !s.mConfiguration.IndexEnable() {
//				return nil
//			}
//			if err := s.mIndexStorage.UpsertIndexMeta(indexMeta); err != nil {
//				internalLogger.Error("IndexStorage UpsertIndexMeta error", zap.Error(err))
//				return err
//			} else {
//				return nil
//			}
//		}
//	} else if pp.FlameProposalType == pb.FlameProposal_UPDATE_INDEX_META {
//		indexMeta := &pb.FlameIndexMeta{}
//		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		if checkNamespaceValidity {
//			if !utility.IsNamespaceValid(indexMeta.Namespace) {
//				return x.ErrInvalidNamespace
//			}
//		}
//
//		if err := s.UpdateIndexMeta(indexMeta); err != nil {
//			internalLogger.Error("UpdateIndexMeta error", zap.Error(err))
//			return err
//		} else {
//			if !s.mConfiguration.IndexEnable() {
//				return nil
//			}
//
//			if err := s.mIndexStorage.UpdateIndexMeta(indexMeta); err != nil {
//				internalLogger.Error("IndexStorage UpdateIndexMeta error", zap.Error(err))
//				return err
//			} else {
//				err := s.UpdateIndex(indexMeta.Namespace)
//				if err != nil {
//					internalLogger.Error("UpdateIndex error", zap.Error(err))
//				}
//
//				return nil
//			}
//		}
//	} else if pp.FlameProposalType == pb.FlameProposal_DELETE_INDEX_META {
//		indexMeta := &pb.FlameIndexMeta{}
//		if err := proto.Unmarshal(pp.FlameProposalData, indexMeta); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		if checkNamespaceValidity {
//			if !utility.IsNamespaceValid(indexMeta.Namespace) {
//				return x.ErrInvalidNamespace
//			}
//		}
//
//		if err := s.DeleteIndexMeta(indexMeta); err != nil {
//			internalLogger.Error("DeleteIndexMeta error", zap.Error(err))
//			return err
//		} else {
//			if !s.mConfiguration.IndexEnable() {
//				return nil
//			}
//			if err := s.mIndexStorage.DeleteIndexMeta(indexMeta); err != nil {
//				internalLogger.Error("IndexStorage DeleteIndexMeta error", zap.Error(err))
//				return err
//			} else {
//				return nil
//			}
//		}
//	} else if pp.FlameProposalType == pb.FlameProposal_CREATE_ACCESS_CONTROL {
//		ac := &pb.FlameAccessControl{}
//		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		if checkNamespaceValidity {
//			if !utility.IsNamespaceValid(ac.Namespace) {
//				return x.ErrInvalidNamespace
//			}
//		}
//
//		return s.CreateAccessControl(ac)
//	} else if pp.FlameProposalType == pb.FlameProposal_UPDATE_ACCESS_CONTROL {
//		ac := &pb.FlameAccessControl{}
//		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		if checkNamespaceValidity {
//			if !utility.IsNamespaceValid(ac.Namespace) {
//				return x.ErrInvalidNamespace
//			}
//		}
//
//		return s.UpdateAccessControl(ac)
//	} else if pp.FlameProposalType == pb.FlameProposal_DELETE_ACCESS_CONTROL {
//		ac := &pb.FlameAccessControl{}
//		if err := proto.Unmarshal(pp.FlameProposalData, ac); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		if checkNamespaceValidity {
//			if !utility.IsNamespaceValid(ac.Namespace) {
//				return x.ErrInvalidNamespace
//			}
//		}
//
//		return s.DeleteAccessControl(ac)
//	} else if pp.FlameProposalType == pb.FlameProposal_CREATE_USER {
//		user := &pb.FlameUser{}
//		if err := proto.Unmarshal(pp.FlameProposalData, user); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		return s.CreateUser(user)
//	} else if pp.FlameProposalType == pb.FlameProposal_UPDATE_USER {
//		user := &pb.FlameUser{}
//		if err := proto.Unmarshal(pp.FlameProposalData, user); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		return s.UpdateUser(user)
//	} else if pp.FlameProposalType == pb.FlameProposal_DELETE_USER {
//		user := &pb.FlameUser{}
//		if err := proto.Unmarshal(pp.FlameProposalData, user); err != nil {
//			return x.ErrDataUnmarshalError
//		}
//
//		return s.DeleteUser(user)
//	}
//
//	return x.ErrFailedToApplyProposal
//}

//func (s *Storage) getIndexHolderMap(batchAction *pb.FlameBatchAction) map[string]*internalIndexHolder {
//	var indexHolderMap = make(map[string]*internalIndexHolder)
//
//	for idx := range batchAction.FlameActionList {
//		flameAction := batchAction.FlameActionList[idx]
//		if flameAction == nil {
//			continue
//		}
//
//		if flameAction.FlameEntry == nil {
//			continue
//		}
//
//		currentIndexHolder, ok := indexHolderMap[string(flameAction.FlameEntry.Namespace)]
//		if !ok {
//			currentIndexHolder = &internalIndexHolder{namespace: string(flameAction.FlameEntry.Namespace)}
//			currentIndexHolder.indexData = make([]*variant.IndexData, 0, s.mConfiguration.CacheSize())
//			indexHolderMap[string(flameAction.FlameEntry.Namespace)] = currentIndexHolder
//		}
//
//		data := s.mConfiguration.IndexObject(flameAction.FlameEntry.Namespace, flameAction.FlameEntry.Value)
//
//		if data == nil {
//			continue
//		}
//
//		id := uidutil.GetUidString(flameAction.FlameEntry.Namespace, flameAction.FlameEntry.Key)
//
//		if flameAction.FlameActionType == pb.FlameAction_CREATE {
//			currentIndexHolder.indexData = append(currentIndexHolder.indexData, &variant.IndexData{
//				ID:     id,
//				Action: variant.CREATE,
//				Data:   data,
//			})
//		}
//
//		if flameAction.FlameActionType == pb.FlameAction_UPDATE {
//			currentIndexHolder.indexData = append(currentIndexHolder.indexData, &variant.IndexData{
//				ID:     id,
//				Action: variant.UPDATE,
//				Data:   data,
//			})
//		}
//
//		if flameAction.FlameActionType == pb.FlameAction_DELETE {
//			currentIndexHolder.indexData = append(currentIndexHolder.indexData, &variant.IndexData{
//				ID:     id,
//				Action: variant.DELETE,
//				Data:   data,
//			})
//		}
//	}
//
//	return indexHolderMap
//}

//func (s *Storage) directIndex(batchAction *pb.FlameBatchAction) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if batchAction == nil {
//		return nil
//	}
//
//	for k, v := range s.getIndexHolderMap(batchAction) {
//		if !s.mIndexStorage.CanIndex(k) && s.mConfiguration.AutoIndexMeta() {
//			//internalLogger.Info("no index found, creating new one", zap.String("namespace",k))
//			flameMeta := &pb.FlameIndexMeta{
//				Namespace: []byte(k),
//				FamilyVersion:   1,
//				Enabled:   true,
//				Default:   true,
//				CreatedAt: uint64(time.Now().UnixNano()),
//				UpdatedAt: uint64(time.Now().UnixNano()),
//			}
//			err := s.mIndexStorage.UpsertIndexMeta(flameMeta)
//			if err != nil {
//				internalLogger.Error("UpsertIndexMeta failure",
//					zap.Error(err),
//					zap.String("namespace", k))
//			}
//		}
//
//		if s.mIndexStorage.CanIndex(k) {
//			err := s.mIndexStorage.ApplyIndex(k, v.indexData)
//			if err != nil {
//				internalLogger.Error("ApplyIndex failure",
//					zap.Error(err),
//					zap.String("namespace", k))
//			}
//		}
//	}
//
//	return nil
//}
//
//func (s *Storage) FullIndex() error {
//	errFound := false
//	err := s.mStateMachineStorage.Iterate([]byte(constant.IndexMetaNamespace), []byte(""), 0, func(entry *pb.FlameEntry) bool {
//		indexMeta := &pb.FlameIndexMeta{}
//		if err := proto.Unmarshal(entry.Value, indexMeta); err != nil {
//			errFound = true
//			internalLogger.Error("proto unmarshal error", zap.Error(err))
//			return false
//		}
//
//		if err := s.mIndexStorage.UpsertIndexMeta(indexMeta); err != nil {
//			internalLogger.Error("UpsertIndexMeta failure", zap.Error(err))
//			errFound = true
//			return false
//		}
//		return true
//	})
//
//	if errFound {
//		return x.ErrFailedToApplyIndex
//	}
//
//	if err != nil {
//		return err
//	}
//
//	idx := 0
//	batchAction := &pb.FlameBatchAction{
//		FlameActionList: make([]*pb.FlameAction, 0, 100),
//	}
//
//	err = s.mStateMachineStorage.Iterate([]byte("A"), []byte(""), 0, func(entry *pb.FlameEntry) bool {
//		action := &pb.FlameAction{
//			FlameActionType: pb.FlameAction_CREATE,
//			FlameEntry:      entry,
//		}
//		batchAction.FlameActionList[idx] = action
//
//		idx = idx + 1
//
//		if idx == 100 {
//			idx = 0
//			err := s.directIndex(batchAction)
//			if err != nil {
//				errFound = true
//				internalLogger.Error("direct index failed",
//					zap.Error(err))
//				return false
//			}
//		}
//
//		return true
//	})
//
//	if errFound {
//		return x.ErrFailedToApplyIndex
//	}
//
//	if idx < 100 {
//		batchAction.FlameActionList = batchAction.FlameActionList[0:idx]
//		err := s.directIndex(batchAction)
//		return err
//	} else {
//		return err
//	}
//}
//
//func (s *Storage) UpdateIndex(namespace []byte) error {
//	indexMeta := &pb.FlameIndexMeta{Namespace: namespace}
//
//	err := s.GetIndexMeta(indexMeta)
//
//	if err != nil {
//		if s.mConfiguration.AutoIndexMeta() {
//			indexMeta = &pb.FlameIndexMeta{
//				Namespace: namespace,
//				FamilyVersion:   1,
//				Enabled:   true,
//				Default:   true,
//				CreatedAt: uint64(time.Now().UnixNano()),
//				UpdatedAt: uint64(time.Now().UnixNano()),
//			}
//		} else {
//			return err
//		}
//	}
//
//	err = s.mIndexStorage.UpdateIndexMeta(indexMeta)
//	if err != nil {
//		return err
//	}
//
//	idx := 0
//	errDirectIndex := false
//	batchAction := &pb.FlameBatchAction{
//		FlameActionList: make([]*pb.FlameAction, 0, 100),
//	}
//
//	err = s.mStateMachineStorage.Iterate(namespace, namespace, 0, func(entry *pb.FlameEntry) bool {
//		action := &pb.FlameAction{
//			FlameActionType: pb.FlameAction_CREATE,
//			FlameEntry:      entry,
//		}
//		batchAction.FlameActionList[idx] = action
//
//		idx = idx + 1
//
//		if idx == 100 {
//			idx = 0
//			err := s.directIndex(batchAction)
//			if err != nil {
//				errDirectIndex = true
//				internalLogger.Error("direct index failed",
//					zap.Error(err),
//					zap.String("namespace", string(namespace)))
//				return false
//			}
//		}
//
//		return true
//	})
//
//	if errDirectIndex {
//		return x.ErrFailedToApplyIndex
//	}
//
//	if idx < 100 {
//		batchAction.FlameActionList = batchAction.FlameActionList[0:idx]
//		err := s.directIndex(batchAction)
//		return err
//	} else {
//		return err
//	}
//}
//

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
