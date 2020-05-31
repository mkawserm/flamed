package storage

import (
	"bytes"
	"context"
	"fmt"
	badgerDb "github.com/dgraph-io/badger/v2"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
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
		Action: pb.Action_UPSERT,
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
		Action: pb.Action_DELETE,
	})
	return nil
}

//func (s *StateContext) AutoIndexMeta() bool {
//	return s.mStorage.AutoIndexMeta()
//}

//func (s *StateContext) CanIndex(namespace string) bool {
//	if !s.mStorage.IndexEnable() {
//		return false
//	}
//
//	return s.mStorage.CanIndex(namespace)
//}

func (s *StateContext) UpsertIndexMeta(meta *pb.IndexMeta) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    pb.Action_UPSERT,
		IndexMeta: meta,
	}
	s.mIndexMetaActionList = append(s.mIndexMetaActionList, m)
	return nil
}

func (s *StateContext) DeleteIndexMeta(meta *pb.IndexMeta) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    pb.Action_DELETE,
		IndexMeta: meta,
	}
	s.mIndexMetaActionList = append(s.mIndexMetaActionList, m)
	return nil
}

func (s *StateContext) DefaultIndexMeta(namespace string) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    pb.Action_DEFAULT,
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

	mIndexTaskQueue variant.TaskQueue
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
	defer func() {
		_ = logger.L("storage").Sync()
	}()
	logger.L("storage").Debug("Opening storage")

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

		s.mIndexTaskQueue = make(variant.TaskQueue, s.mConfiguration.QueueSize())
		go s.indexTaskQueueHandler()
	}

	go s.storageTaskQueueHandler()

	logger.L("storage").Debug("Storage opened")
	return nil
}

func (s *Storage) Close() error {
	err1 := s.mStateStorage.Close()

	if err1 != nil {
		return err1
	}

	if s.mConfiguration.IndexEnable() {
		s.mIndexTaskQueue <- variant.Task{
			ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
			Name:    "index-task",
			Command: "done",
		}
		close(s.mIndexTaskQueue)
		s.mIndexTaskQueue = nil
		err2 := s.mIndexStorage.Close()
		return err2
	}

	if s.mConfiguration.StorageTaskQueue() != nil {
		s.mConfiguration.StorageTaskQueue() <- variant.Task{
			ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
			Name:    "storage-task",
			Command: "done",
		}
	}

	s.mConfiguration = nil
	return nil
}

func (s *Storage) RunGC() {
	defer func() {
		_ = logger.L("storage").Sync()
	}()

	logger.L("storage").Info("running storage gc")
	s.mStateStorage.RunGC()
	if s.mConfiguration.IndexEnable() {
		if s.mIndexStorage != nil {
			s.mIndexStorage.RunGC()
		}
	}
	logger.L("storage").Info("storage gc finished")
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
		_ = logger.L("storage").Sync()
	}()

	if s.mStateStorage == nil {
		return x.ErrStorageIsNotReady
	}

	uid := crypto.GetStateAddress(namespace, key)

	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()
	if err := txn.Set(uid, value); err != nil {
		logger.L("storage").Error("set failure", zap.Error(err))
		return x.ErrFailedToCreateDataToStorage
	}

	if err := txn.Commit(); err != nil {
		logger.L("storage").Error("txn commit failure", zap.Error(err))
		return x.ErrFailedToCreateDataToStorage
	}

	return nil
}

func (s *Storage) Read(namespace []byte, key []byte) ([]byte, error) {
	defer func() {
		_ = logger.L("storage").Sync()
	}()

	if s.mStateStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	uid := crypto.GetStateAddress(namespace, key)
	txn := s.mStateStorage.NewReadOnlyTransaction()
	defer txn.Discard()

	if val, err := txn.Get(uid); err != nil {
		if err == x.ErrAddressNotFound {
			return nil, x.ErrStateNotFound
		} else {
			logger.L("storage").Error("read failure", zap.Error(err))
			return nil, x.ErrFailedToReadDataFromStorage
		}
	} else {
		return val, nil
	}
}

func (s *Storage) Delete(namespace []byte, key []byte) error {
	defer func() {
		_ = logger.L("storage").Sync()
	}()

	if s.mStateStorage == nil {
		return x.ErrStorageIsNotReady
	}

	uid := crypto.GetStateAddress(namespace, key)

	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()

	if err := txn.Delete(uid); err != nil {
		logger.L("storage").Error("deletion failure", zap.Error(err))
		return x.ErrFailedToDeleteDataFromStorage
	}

	if err := txn.Commit(); err != nil {
		logger.L("storage").Error("txn commit failure", zap.Error(err))
		return x.ErrFailedToCreateDataToStorage
	}

	return nil
}

func (s *Storage) SaveAppliedIndex(u uint64) error {
	entry := &pb.StateEntry{
		Payload:   crypto.Uint64ToByteSlice(u),
		Namespace: []byte(constant.AppliedIndexNamespace),
	}

	if data, err := proto.Marshal(entry); err != nil {
		return err
	} else {
		return s.Create(
			[]byte(constant.AppliedIndexNamespace),
			[]byte(constant.AppliedIndexKey),
			data)
	}
}

func (s *Storage) QueryAppliedIndex() (uint64, error) {
	data, err := s.Read(
		[]byte(constant.AppliedIndexNamespace),
		[]byte(constant.AppliedIndexKey))

	if err == x.ErrStateNotFound {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	entry := &pb.StateEntry{}

	if err := proto.Unmarshal(data, entry); err != nil {
		return 0, err
	}
	return crypto.ByteSliceToUint64(entry.Payload), nil
}

func (s *Storage) GetAppliedIndex(_ context.Context) (interface{}, error) {
	return s.QueryAppliedIndex()
}

func (s *Storage) Search(ctx context.Context, searchInput *pb.SearchInput) (interface{}, error) {
	if len(searchInput.FamilyName) != 0 && len(searchInput.FamilyVersion) != 0 {
		tp := s.mConfiguration.GetTransactionProcessor(searchInput.FamilyName, searchInput.FamilyVersion)
		if tp == nil {
			return nil, x.ErrTPNotFound
		}
		readOnlyTxn := s.mStateStorage.NewReadOnlyTransaction()
		defer readOnlyTxn.Discard()
		readOnlyStateContext := &StateContext{
			mReadOnly: true,
			mStorage:  s,
			mTxn:      readOnlyTxn,
		}

		return tp.Search(ctx, readOnlyStateContext, searchInput)
	}

	return nil, x.ErrTPNotFound
}

func (s *Storage) Iterate(ctx context.Context, iterateInput *pb.IterateInput) (interface{}, error) {
	if len(iterateInput.FamilyName) != 0 && len(iterateInput.FamilyVersion) != 0 {
		tp := s.mConfiguration.GetTransactionProcessor(
			iterateInput.FamilyName,
			iterateInput.FamilyVersion)
		if tp == nil {
			return nil, x.ErrTPNotFound
		}
		readOnlyTxn := s.mStateStorage.NewReadOnlyTransaction()
		defer readOnlyTxn.Discard()
		readOnlyStateContext := &StateContext{
			mReadOnly: true,
			mStorage:  s,
			mTxn:      readOnlyTxn,
		}

		return tp.Iterate(ctx, readOnlyStateContext, iterateInput)
	}

	return nil, x.ErrTPNotFound
}

func (s *Storage) Retrieve(ctx context.Context, retrieveInput *pb.RetrieveInput) (interface{}, error) {
	if len(retrieveInput.FamilyName) != 0 && len(retrieveInput.FamilyVersion) != 0 {
		tp := s.mConfiguration.GetTransactionProcessor(
			retrieveInput.FamilyName,
			retrieveInput.FamilyVersion)
		if tp == nil {
			return nil, x.ErrTPNotFound
		}
		readOnlyTxn := s.mStateStorage.NewReadOnlyTransaction()
		defer readOnlyTxn.Discard()
		readOnlyStateContext := &StateContext{
			mReadOnly: true,
			mStorage:  s,
			mTxn:      readOnlyTxn,
		}

		return tp.Retrieve(ctx, readOnlyStateContext, retrieveInput)
	}

	return nil, x.ErrTPNotFound
}

func (s *Storage) GlobalSearch(ctx context.Context, input *pb.GlobalSearchInput) (interface{}, error) {
	if s.mIndexStorage == nil {
		return nil, x.ErrIndexStorageIsNotReady
	}

	return s.mIndexStorage.GlobalSearch(ctx, input)
}

func (s *Storage) GlobalIterate(_ context.Context, globalIterate *pb.GlobalIterateInput) (interface{}, error) {
	if len(globalIterate.Namespace) == 0 {
		return nil, x.ErrInvalidNamespace
	}

	if globalIterate.Limit == 0 {
		return nil, nil
	}

	if s.mStateStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	readOnlyTxn := s.mStateStorage.NewReadOnlyTransaction()
	defer readOnlyTxn.Discard()
	readOnlyStateContext := &StateContext{
		mReadOnly: true,
		mStorage:  s,
		mTxn:      readOnlyTxn,
	}

	if len(globalIterate.From) == 0 {
		globalIterate.From = globalIterate.Namespace
	}

	if len(globalIterate.Prefix) == 0 {
		globalIterate.Prefix = globalIterate.Namespace
	}

	if !bytes.HasPrefix(globalIterate.From, globalIterate.Namespace) {
		return nil, x.ErrAccessViolation
	}
	if !bytes.HasPrefix(globalIterate.Prefix, globalIterate.Namespace) {
		return nil, x.ErrAccessViolation
	}

	logger.L("storage").Debug("global iteration in progress")
	logger.L("storage").Debug("iteration input", zap.String("input", globalIterate.String()))
	itr := readOnlyStateContext.GetForwardIterator()
	defer itr.Close()

	stateEntryResponses := make([]*pb.StateEntryResponse, 0, globalIterate.Limit)
	iterCounter := uint64(0)
	for itr.Seek(globalIterate.From); itr.ValidForPrefix(globalIterate.Prefix); itr.Next() {
		iterCounter = iterCounter + 1
		sts := itr.StateSnapshot()
		if sts != nil {
			stateEntryResponse := &pb.StateEntryResponse{}
			stateEntryResponse.StateAvailable = true
			stateEntryResponse.Address = crypto.StateAddressByteSliceToHexString(sts.Address)
			stateEntryResponse.StateEntry = sts.ToStateEntry()
			stateEntryResponses = append(stateEntryResponses, stateEntryResponse)
		}
		if iterCounter == globalIterate.Limit {
			break
		}
	}
	logger.L("storage").Debug("global iteration done")
	return stateEntryResponses, nil
}

func (s *Storage) GlobalRetrieve(_ context.Context, globalRetrieveInput *pb.GlobalRetrieveInput) (interface{}, error) {
	if len(globalRetrieveInput.Addresses) == 0 {
		return nil, nil
	}

	if s.mStateStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	readOnlyTxn := s.mStateStorage.NewReadOnlyTransaction()
	defer readOnlyTxn.Discard()
	readOnlyStateContext := &StateContext{
		mReadOnly: true,
		mStorage:  s,
		mTxn:      readOnlyTxn,
	}

	stateEntryResponses := make([]*pb.StateEntryResponse, 0, len(globalRetrieveInput.Addresses))
	for _, sa := range globalRetrieveInput.Addresses {
		if bytes.HasPrefix(sa, globalRetrieveInput.Namespace) {
			stateEntryResponse := &pb.StateEntryResponse{}
			stateEntryResponse.StateAvailable = true
			stateEntryResponse.Address = crypto.StateAddressByteSliceToHexString(sa)
			state, err := readOnlyStateContext.GetState(sa)
			if err != nil {
				stateEntryResponse.StateAvailable = false
			}
			stateEntryResponse.StateEntry = state
			stateEntryResponses = append(stateEntryResponses, stateEntryResponse)
		} else {
			stateEntryResponse := &pb.StateEntryResponse{}
			stateEntryResponse.StateAvailable = false
			stateEntryResponse.Address = crypto.StateAddressByteSliceToHexString(sa)
			stateEntryResponses = append(stateEntryResponses, stateEntryResponse)
		}
	}
	return stateEntryResponses, nil
}

func (s *Storage) ApplyProposal(ctx context.Context, proposal *pb.Proposal, entryIndex uint64) *pb.ProposalResponse {
	defer func() {
		_ = logger.L("storage").Sync()
	}()

	logger.L("storage").Info("raft entry index", zap.Uint64("entryIndex", entryIndex))

	pr := pb.NewProposalResponse(pb.Status_REJECTED)
	pr.Uuid = proposal.Uuid

	if len(proposal.Uuid) == 0 {
		pr.Status = pb.Status_REJECTED
		pr.ErrorCode = 0
		pr.ErrorText = "uuid can not be empty"
		return pr
	}

	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()

	var indexDataContainer = make(IndexDataContainer)
	var indexMetaActionContainer = make(IndexMetaActionContainer)

	for _, t := range proposal.Transactions {
		tp := s.mConfiguration.GetTransactionProcessor(t.FamilyName, t.FamilyVersion)
		if tp == nil {
			pr.Status = pb.Status_REJECTED
			pr.ErrorCode = 0
			pr.ErrorText = "Transaction family name:" +
				t.FamilyName + " family version:" +
				t.FamilyVersion + " not found"
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

		if tpr.Status == pb.Status_REJECTED {
			stateContext.mIndexDataList = nil
			stateContext.mIndexMetaActionList = nil
			pr.Status = pb.Status_REJECTED
			pr.ErrorCode = 0
			pr.ErrorText = "proposal rejected"
			return pr
		} else {
			if len(stateContext.mIndexDataList) > 0 {
				indexDataContainer[string(t.Namespace)] = stateContext.mIndexDataList
			}
			if len(stateContext.mIndexMetaActionList) > 0 {
				indexMetaActionContainer[string(t.Namespace)] = stateContext.mIndexMetaActionList
			}
		}
	}

	if err := txn.Commit(); err == nil {
		s.mConfiguration.ProposalReceiver(proposal, pb.Status_ACCEPTED)
		if s.mConfiguration.IndexEnable() {
			//NOTE: update indexmeta meta
			s.updateIndexMetaOfIndexStorage(indexMetaActionContainer)

			//NOTE: index data
			s.updateIndexOfIndexStorage(indexDataContainer)
		}

		pr.Status = pb.Status_ACCEPTED
		return pr
	} else {
		s.mConfiguration.ProposalReceiver(proposal, pb.Status_REJECTED)
		indexDataContainer = nil
		indexMetaActionContainer = nil
		pr.Status = pb.Status_REJECTED
		pr.ErrorCode = 0
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
	/* NOTE: INDEX DATA TO BE PROCESSED USING GO CHANNEL */
	s.mIndexTaskQueue <- variant.Task{
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:    "index-task",
		Command: "index-data-container",
		Payload: indexDataContainer,
	}

	//err := s.directIndex(indexDataContainer)
	//
	//if err != nil {
	//	logger.L("storage").Error("directIndex error", zap.Error(err))
	//}
}

func (s *Storage) updateIndexMetaOfIndexStorage(indexMetaActionContainer IndexMetaActionContainer) {
	if len(indexMetaActionContainer) == 0 {
		return
	}

	for _, v := range indexMetaActionContainer {
		for _, v2 := range v {
			if v2.Action == pb.Action_UPSERT {
				err := s.mIndexStorage.UpsertIndexMeta(v2.IndexMeta)
				if err != nil {
					logger.L("storage").Error("upsert indexmeta error", zap.Error(err))
				}

				if s.mConfiguration.AutoBuildIndex() {
					s.BuildIndexByNamespace(v2.IndexMeta.Namespace)
				}
			}
			if v2.Action == pb.Action_DELETE {
				err := s.mIndexStorage.DeleteIndexMeta(v2.IndexMeta)
				if err != nil {
					logger.L("storage").Error("delete indexmeta error", zap.Error(err))
				}
			}

			if v2.Action == pb.Action_DEFAULT {
				err := s.mIndexStorage.DefaultIndexMeta(string(v2.IndexMeta.Namespace))
				if err != nil {
					logger.L("storage").Error("default indexmeta error", zap.Error(err))
				}

				if s.mConfiguration.AutoBuildIndex() {
					s.BuildIndexByNamespace(v2.IndexMeta.Namespace)
				}
			}
		}
	}
}

func (s *Storage) PrepareSnapshot() (interface{}, error) {
	defer func() {
		_ = logger.L("storage").Sync()
	}()

	if s.mStateStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	logger.L("storage").Debug("snapshot prepared")
	return s.mStateStorage.NewTransaction(), nil
}

func (s *Storage) RecoverFromSnapshot(r io.Reader) error {
	defer func() {
		_ = logger.L("storage").Sync()
	}()

	logger.L("storage").Debug("recovering from snapshot")

	if s.mStateStorage == nil {
		return x.ErrStorageIsNotReady
	}

	sz := make([]byte, 8)

	//if _, err := io.ReadFull(r, sz); err != nil {
	//	logger.L("storage").Error("read error", zap.Error(err))
	//	return x.ErrFailedToRecoverFromSnapshot
	//}
	//
	//total := crypto.ByteSliceToUint64(sz)

	txn := s.mStateStorage.NewTransaction()
	defer txn.Discard()

	//indexDataContainer := make(IndexDataContainer)

	for i := uint64(0); ; i++ {
		if _, err := io.ReadFull(r, sz); err == io.ErrUnexpectedEOF || err == io.EOF {
			break
		} else if err != nil {
			logger.L("storage").Error("sm read error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		toRead := crypto.ByteSliceToUint64(sz)
		data := make([]byte, toRead)
		if _, err := io.ReadFull(r, data); err != nil {
			logger.L("storage").Error("sm read error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		snap := &pb.StateSnapshot{}
		if err := proto.Unmarshal(data, snap); err != nil {
			logger.L("storage").Error("sm unmarshal error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		//entry := &pb.StateEntry{}
		//if err := proto.Unmarshal(snap.Data, entry); err == nil {
		//	if entry.FamilyName != "" {
		//		if entry.FamilyName == "IndexMeta" {
		//			meta := &pb.IndexMeta{}
		//			if err := proto.Unmarshal(entry.Payload, meta); err!=nil {
		//
		//			}
		//		} else {
		//			tp := s.mConfiguration.GetTransactionProcessor(entry.FamilyName,
		//				entry.FamilyVersion)
		//			if tp != nil {
		//				obj := tp.IndexObject(entry.Payload)
		//				if obj != nil {
		//
		//				}
		//			}
		//		}
		//	}
		//}

		if err := txn.Set(snap.Address, snap.Data); err == badgerDb.ErrTxnTooBig {
			if err := txn.Commit(); err != nil {
				logger.L("storage").Error("txn commit error", zap.Error(err))
				return x.ErrFailedToRecoverFromSnapshot
			}

			txn = s.mStateStorage.NewTransaction()

			if err := txn.Set(snap.Address, snap.Data); err != nil {
				logger.L("storage").Error("txn set error", zap.Error(err))
				return x.ErrFailedToRecoverFromSnapshot
			}

		} else if err != nil {
			logger.L("storage").Error("txn set error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}
	}

	if err := txn.Commit(); err != nil {
		logger.L("storage").Error("txn commit error", zap.Error(err))
		return x.ErrFailedToRecoverFromSnapshot
	}

	logger.L("storage").Debug("storage recovered from snapshot")
	if s.mConfiguration.AutoBuildIndex() {
		s.BuildIndex()
	}
	return nil
}

func (s *Storage) SaveSnapshot(snapshotContext interface{}, w io.Writer) error {
	defer func() {
		_ = logger.L("storage").Sync()
	}()

	logger.L("storage").Debug("saving snapshot")
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

	//total := uint64(0)
	//it := txn.KeyOnlyForwardIterator()
	//for it.Rewind(); it.Valid(); it.Next() {
	//	total = total + 1
	//}
	//it.Close()
	//
	//if _, err := w.Write(crypto.Uint64ToByteSlice(total)); err != nil {
	//	logger.L("storage").Error("storage write error", zap.Error(err))
	//	return x.ErrFailedToSaveSnapshot
	//}

	it := txn.ForwardIterator()
	defer it.Close()

	for it.Rewind(); it.Valid(); it.Next() {
		item := it.StateSnapshot()
		if data, err := proto.Marshal(item); err == nil {
			dataLength := uint64(len(data))
			if _, err := w.Write(crypto.Uint64ToByteSlice(dataLength)); err != nil {
				logger.L("storage").Error("storage write error", zap.Error(err))
				return x.ErrFailedToSaveSnapshot
			}
			if _, err := w.Write(data); err != nil {
				return x.ErrFailedToSaveSnapshot
			}
		} else {
			logger.L("storage").Error("state snapshot marshal error", zap.Error(err))
			return x.ErrFailedToSaveSnapshot
		}
	}

	logger.L("storage").Debug("storage snapshot saved")
	return nil
}

func (s *Storage) BuildIndex() {
	if !s.mConfiguration.IndexEnable() {
		return
	}

	s.mIndexTaskQueue <- variant.Task{
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:    "index-task",
		Command: "build-index",
	}
}

func (s *Storage) BuildIndexByNamespace(namespace []byte) {
	if !s.mConfiguration.IndexEnable() {
		return
	}

	s.mIndexTaskQueue <- variant.Task{
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:    "index-task",
		Command: "build-index-by-namespace",
		Payload: namespace,
	}
}

func (s *Storage) buildIndex() error {
	if !s.mConfiguration.IndexEnable() {
		return nil
	}

	txn := s.mStateStorage.NewReadOnlyTransaction()
	it := txn.ForwardIterator()
	defer it.Close()

	for it.Seek([]byte(constant.IndexMetaNamespace)); it.ValidForPrefix([]byte(constant.IndexMetaNamespace)); it.Next() {
		snap := it.StateSnapshot()
		if snap != nil {
			entry := &pb.StateEntry{}
			if err := proto.Unmarshal(snap.Data, entry); err != nil {
				return err
			}

			indexMeta := &pb.IndexMeta{}
			if err := proto.Unmarshal(entry.Payload, indexMeta); err != nil {
				return err
			}

			if err := s.mIndexStorage.UpsertIndexMeta(indexMeta); err != nil {
				logger.L("storage").Error("UpsertIndexMeta failure", zap.Error(err))
				return err
			}
		}
	}

	indexDataList := make([]*variant.IndexData, 0, 100)
	oldNamespace := ""
	for it.Seek([]byte("A")); it.Valid(); it.Next() {
		snap := it.StateSnapshot()
		if snap != nil {
			entry := &pb.StateEntry{}
			if err := proto.Unmarshal(snap.Data, entry); err != nil {
				return err
			}
			tp := s.mConfiguration.GetTransactionProcessor(entry.FamilyName, entry.FamilyVersion)
			if tp == nil {
				continue
			}

			indexData := tp.IndexObject(entry.Payload)
			if indexData == nil {
				continue
			}

			indexDataList = append(indexDataList, &variant.IndexData{
				ID:     crypto.StateAddressByteSliceToHexString(snap.Address),
				Data:   indexData,
				Action: pb.Action_UPSERT,
			})

			currentNamespace := string(entry.Namespace)
			if oldNamespace != currentNamespace {
				if len(indexDataList) != 0 {
					indexDataContainer := IndexDataContainer{oldNamespace: indexDataList}
					err := s.directIndex(indexDataContainer)
					if err != nil {
						logger.L("storage").Error("indexing error", zap.Error(err))
					}
					indexDataList = make([]*variant.IndexData, 0, 100)
				}

				oldNamespace = currentNamespace
			}

			if len(indexDataList) == 100 {
				indexDataContainer := IndexDataContainer{oldNamespace: indexDataList}
				err := s.directIndex(indexDataContainer)
				if err != nil {
					logger.L("storage").Error("indexing error", zap.Error(err))
				}
				indexDataList = make([]*variant.IndexData, 0, 100)
			}
		}
	}

	if len(indexDataList) != 0 {
		indexDataContainer := IndexDataContainer{oldNamespace: indexDataList}
		err := s.directIndex(indexDataContainer)
		if err != nil {
			logger.L("storage").Error("indexing error", zap.Error(err))
		}
	}

	return nil
}

func (s *Storage) buildIndexByNamespace(namespace []byte) error {
	if !s.mConfiguration.IndexEnable() {
		return nil
	}

	txn := s.mStateStorage.NewReadOnlyTransaction()
	it := txn.ForwardIterator()
	defer it.Close()

	data, err := txn.Get(crypto.GetStateAddress([]byte(constant.IndexMetaNamespace), namespace))

	if err != x.ErrAddressNotFound && err != nil {
		return err
	}

	if data != nil {
		entry := &pb.StateEntry{}
		if err := proto.Unmarshal(data, entry); err != nil {
			return err
		}

		indexMeta := &pb.IndexMeta{}
		if err := proto.Unmarshal(entry.Payload, indexMeta); err != nil {
			return err
		}

		if err := s.mIndexStorage.UpsertIndexMeta(indexMeta); err != nil {
			logger.L("storage").Error("UpsertIndexMeta failure", zap.Error(err))
			return err
		}
	}

	indexDataList := make([]*variant.IndexData, 0, 100)
	oldNamespace := ""
	for it.Seek(namespace); it.ValidForPrefix(namespace); it.Next() {
		snap := it.StateSnapshot()
		if snap != nil {
			entry := &pb.StateEntry{}
			if err := proto.Unmarshal(snap.Data, entry); err != nil {
				return err
			}
			tp := s.mConfiguration.GetTransactionProcessor(entry.FamilyName, entry.FamilyVersion)
			if tp == nil {
				continue
			}

			indexData := tp.IndexObject(entry.Payload)
			if indexData == nil {
				continue
			}

			indexDataList = append(indexDataList, &variant.IndexData{
				ID:     crypto.StateAddressByteSliceToHexString(snap.Address),
				Data:   indexData,
				Action: pb.Action_UPSERT,
			})

			currentNamespace := string(entry.Namespace)
			if oldNamespace != currentNamespace {
				if len(indexDataList) != 0 {
					indexDataContainer := IndexDataContainer{oldNamespace: indexDataList}
					err := s.directIndex(indexDataContainer)
					if err != nil {
						logger.L("storage").Error("indexing error", zap.Error(err))
					}
					indexDataList = make([]*variant.IndexData, 0, 100)
				}

				oldNamespace = currentNamespace
			}

			if len(indexDataList) == 100 {
				indexDataContainer := IndexDataContainer{oldNamespace: indexDataList}
				err := s.directIndex(indexDataContainer)
				if err != nil {
					logger.L("storage").Error("indexing error", zap.Error(err))
				}
				indexDataList = make([]*variant.IndexData, 0, 100)
			}
		}
	}

	if len(indexDataList) != 0 {
		indexDataContainer := IndexDataContainer{oldNamespace: indexDataList}
		err := s.directIndex(indexDataContainer)
		if err != nil {
			logger.L("storage").Error("indexing error", zap.Error(err))
		}
	}

	return nil
}

func (s *Storage) directIndex(indexDataContainer IndexDataContainer) error {
	defer func() {
		_ = logger.L("storage").Sync()
	}()

	if indexDataContainer == nil {
		return nil
	}

	for k, v := range indexDataContainer {
		if !s.mIndexStorage.CanIndex(k) && s.mConfiguration.AutoIndexMeta() {
			//logger.L("storage").Info("no indexmeta found, creating new one", zap.String("namespace",k))
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
				logger.L("storage").Error("UpsertIndexMeta failure",
					zap.Error(err),
					zap.String("namespace", k))
			}
		}

		if s.mIndexStorage.CanIndex(k) {
			err := s.mIndexStorage.ApplyIndex(k, v)
			if err != nil {
				logger.L("storage").Error("ApplyIndex failure",
					zap.Error(err),
					zap.String("namespace", k))
			}
		}
	}

	return nil
}

func (s *Storage) storageTaskQueueHandler() {
	defer func() {
		_ = logger.L("storage").Sync()
	}()
	logger.L("storage").Info("storage task queue handler started")

	q := s.mConfiguration.StorageTaskQueue()
	if q == nil {
		return
	}

	logger.L("storage").Info("entering into forever loop")
	for {
		task := <-q
		logger.L("storage").Info("executing task",
			zap.String("id", task.ID),
			zap.String("name", task.Name),
			zap.String("command", task.Command),
		)
		_ = logger.L("storage").Sync()

		switch task.Command {
		case "gc":
			s.RunGC()

		case "build-index":
			s.BuildIndex()

		case "build-index-by-namespace":
			if v, ok := task.Payload.([]byte); ok {
				s.BuildIndexByNamespace(v)
			}

		case "done":
			logger.L("storage").Info("storage task queue handler finished")
			break
		}
	}
}

func (s *Storage) indexTaskQueueHandler() {
	defer func() {
		logger.L("storage").Info("index task queue handler exiting")
		_ = logger.L("storage").Sync()
	}()
	logger.L("storage").Info("index task queue handler started")

	if s.mIndexTaskQueue == nil {
		return
	}

	logger.L("storage").Info("entering into forever loop")

	for {
		task := <-s.mIndexTaskQueue
		logger.L("storage").Info("executing task",
			zap.String("id", task.ID),
			zap.String("name", task.Name),
			zap.String("command", task.Command),
		)
		_ = logger.L("storage").Sync()

		switch task.Command {

		case "index-data-container":
			logger.L("storage").Debug("running direct index")
			if indexDataContainer, ok := task.Payload.(IndexDataContainer); ok {
				err := s.directIndex(indexDataContainer)
				if err != nil {
					logger.L("storage").Error("index error", zap.Error(err))
				}
			}
			logger.L("storage").Debug("direct index finished")
		case "build-index":
			logger.L("storage").Debug("build index started")
			err := s.buildIndex()
			if err != nil {
				logger.L("storage").Error("index error", zap.Error(err))
			}
			logger.L("storage").Debug("build index finished")
		case "build-index-by-namespace":
			logger.L("storage").Debug("build index by namespace started")
			if v, ok := task.Payload.([]byte); ok {
				err := s.buildIndexByNamespace(v)
				if err != nil {
					logger.L("storage").Error("index error", zap.Error(err))
				}
			}
			logger.L("storage").Debug("build index by namespace finished")

		case "done":
			logger.L("storage").Info("index task queue handler finished")
			break
		}
	}
}
