package flamed

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/lni/dragonboat/v3"
	sm "github.com/lni/dragonboat/v3/statemachine"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/storage"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Batch struct {
	mNamespace  []byte
	mMutex      *sync.Mutex
	mActionList []*pb.FlameAction
	mEntryList  []*pb.FlameEntry
	mCacheSize  int
}

func (b *Batch) Reset() {
	b.mMutex.Lock()
	defer b.mMutex.Unlock()

	if len(b.mActionList) != 0 {
		b.mActionList = make([]*pb.FlameAction, 0, b.mCacheSize)
	}

	if len(b.mEntryList) != 0 {
		b.mEntryList = make([]*pb.FlameEntry, 0, b.mCacheSize)
	}
}

func (b *Batch) Read(key []byte) {
	b.mMutex.Lock()
	defer b.mMutex.Unlock()

	entry := &pb.FlameEntry{
		Namespace: b.mNamespace,
		Key:       key,
	}

	b.mEntryList = append(b.mEntryList, entry)
}

func (b *Batch) Add(key, value []byte) {
	b.Create(key, value)
}

func (b *Batch) Create(key, value []byte) {
	b.mMutex.Lock()
	defer b.mMutex.Unlock()

	action := &pb.FlameAction{
		FlameEntry: &pb.FlameEntry{
			Namespace: b.mNamespace,
			Key:       key,
			Value:     value,
		},
		FlameActionType: pb.FlameAction_CREATE,
	}
	b.mActionList = append(b.mActionList, action)
}

func (b *Batch) Update(key, value []byte) {
	b.mMutex.Lock()
	defer b.mMutex.Unlock()

	action := &pb.FlameAction{
		FlameEntry: &pb.FlameEntry{
			Namespace: b.mNamespace,
			Key:       key,
			Value:     value,
		},
		FlameActionType: pb.FlameAction_UPDATE,
	}
	b.mActionList = append(b.mActionList, action)
}

func (b *Batch) Delete(key []byte) {
	b.mMutex.Lock()
	defer b.mMutex.Unlock()

	action := &pb.FlameAction{
		FlameEntry: &pb.FlameEntry{
			Namespace: b.mNamespace,
			Key:       key,
		},
		FlameActionType: pb.FlameAction_DELETE,
	}
	b.mActionList = append(b.mActionList, action)
}

func (b *Batch) NewBatch(namespace string) *Batch {
	ns := []byte(namespace)
	if !utility.IsNamespaceValid(ns) {
		return nil
	}

	return &Batch{
		mNamespace:  ns,
		mMutex:      b.mMutex,
		mEntryList:  b.mEntryList,
		mActionList: b.mActionList,
		mCacheSize:  b.mCacheSize,
	}
}

type StorageManager struct {
	mCacheSize          int
	mClusterID          uint64
	mDragonboatNodeHost *dragonboat.NodeHost
}

func (m *StorageManager) NewBatch(namespace string) *Batch {
	ns := []byte(namespace)
	if !utility.IsNamespaceValid(ns) {
		return nil
	}

	return &Batch{
		mNamespace:  ns,
		mMutex:      &sync.Mutex{},
		mActionList: make([]*pb.FlameAction, 0, m.mCacheSize),
		mEntryList:  make([]*pb.FlameEntry, 0, m.mCacheSize),
		mCacheSize:  m.mCacheSize,
	}
}

func (m *StorageManager) GetUser(username string, timeout time.Duration) *pb.FlameUser {
	user := &pb.FlameUser{Username: username}
	_, err := m.managedSyncRead(m.mClusterID, user, timeout)
	if err != nil {
		internalLogger.Error("failed to get user", zap.Error(err))
		return nil
	}
	return user
}

func (m *StorageManager) GetAccessControl(namespace, username string, timeout time.Duration) *pb.FlameAccessControl {
	if !utility.IsUsernameValid(username) {
		return nil
	}

	if !utility.IsNamespaceValid([]byte(namespace)) {
		return nil
	}

	ac := &pb.FlameAccessControl{Username: username, Namespace: []byte(namespace)}
	_, err := m.managedSyncRead(m.mClusterID, ac, timeout)
	if err != nil {
		internalLogger.Error("failed to get access control", zap.Error(err))
		return nil
	}

	return ac
}

func (m *StorageManager) Iterate(seek *pb.FlameEntry, limit int, timeout time.Duration) ([]*pb.FlameEntry, error) {
	if !utility.IsNamespaceValid(seek.Namespace) {
		return nil, x.ErrInvalidNamespace
	}

	allocationLength := 100
	if limit != 0 {
		allocationLength = limit
	}

	newLimit := limit
	if len(seek.Key) != 0 {
		if newLimit != 0 {
			newLimit = newLimit + 1
		}
	}

	data := make([]*pb.FlameEntry, 0, allocationLength)
	uid := uidutil.GetUid(seek.Namespace, seek.Key)

	skipFirstEntry := true
	if len(seek.Key) == 0 {
		skipFirstEntry = false
	}

	itr := &storage.Iterator{
		Seek:   uid,
		Prefix: seek.Namespace,
		Limit:  newLimit,
		Receiver: func(entry *pb.FlameEntry) bool {
			if skipFirstEntry {
				skipFirstEntry = false
				return true
			}

			data = append(data, entry)
			return true
		},
	}

	if _, err := m.managedSyncRead(m.mClusterID, itr, timeout); err != nil {
		return nil, x.ErrFailedToIterate
	} else {
		return data, nil
	}
}

func (m *StorageManager) Get(entry *pb.FlameEntry, timeout time.Duration) error {
	return m.Read(entry, timeout)
}

func (m *StorageManager) Read(entry *pb.FlameEntry, timeout time.Duration) error {
	_, err := m.managedSyncRead(m.mClusterID, entry, timeout)
	return err
}

func (m *StorageManager) Create(entry *pb.FlameEntry, timeout time.Duration) error {
	if !utility.IsNamespaceValid(entry.Namespace) {
		return x.ErrInvalidNamespace
	}

	batch := &pb.FlameBatchAction{
		FlameActionList: []*pb.FlameAction{
			{
				FlameEntry:      entry,
				FlameActionType: pb.FlameAction_CREATE,
			},
		},
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_BATCH_ACTION,
	}

	if data, err := proto.Marshal(batch); err == nil {
		pp.FlameProposalData = data
	} else {
		internalLogger.Error("data marshal error", zap.Error(err))
		return x.ErrDataMarshalError
	}

	r, err := m.managedSyncApplyProposal(m.mClusterID, pp, timeout)

	if err != nil {
		internalLogger.Error("proposal apply error", zap.Error(err))
		return x.ErrFailedToApplyProposal
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToApplyProposal
	}
}

func (m *StorageManager) Update(entry *pb.FlameEntry, timeout time.Duration) error {
	if !utility.IsNamespaceValid(entry.Namespace) {
		return x.ErrInvalidNamespace
	}

	batch := &pb.FlameBatchAction{
		FlameActionList: []*pb.FlameAction{
			{
				FlameEntry:      entry,
				FlameActionType: pb.FlameAction_UPDATE,
			},
		},
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_BATCH_ACTION,
	}

	if data, err := proto.Marshal(batch); err == nil {
		pp.FlameProposalData = data
	} else {
		internalLogger.Error("data marshal error", zap.Error(err))
		return x.ErrDataMarshalError
	}

	r, err := m.managedSyncApplyProposal(m.mClusterID, pp, timeout)

	if err != nil {
		internalLogger.Error("proposal apply error", zap.Error(err))
		return x.ErrFailedToApplyProposal
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToApplyProposal
	}
}

func (m *StorageManager) Delete(entry *pb.FlameEntry, timeout time.Duration) error {
	if !utility.IsNamespaceValid(entry.Namespace) {
		return x.ErrInvalidNamespace
	}

	batch := &pb.FlameBatchAction{
		FlameActionList: []*pb.FlameAction{
			{
				FlameEntry:      entry,
				FlameActionType: pb.FlameAction_DELETE,
			},
		},
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_BATCH_ACTION,
	}

	if data, err := proto.Marshal(batch); err == nil {
		pp.FlameProposalData = data
	} else {
		internalLogger.Error("data marshal error", zap.Error(err))
		return x.ErrDataMarshalError
	}

	r, err := m.managedSyncApplyProposal(m.mClusterID, pp, timeout)

	if err != nil {
		internalLogger.Error("proposal apply error", zap.Error(err))
		return x.ErrFailedToApplyProposal
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToApplyProposal
	}
}

func (m *StorageManager) ReadBatch(batch *Batch, timeout time.Duration) (*pb.FlameBatchRead, error) {
	if len(batch.mEntryList) == 0 {
		return nil, x.ErrEmptyBatch
	}

	rb := &pb.FlameBatchRead{
		FlameEntryList: batch.mEntryList,
	}

	_, err := m.managedSyncRead(m.mClusterID, rb, timeout)
	batch.Reset()

	if err != nil {
		return nil, err
	}

	return rb, nil
}

func (m *StorageManager) ApplyBatch(batch *Batch, timeout time.Duration) error {
	if len(batch.mActionList) == 0 {
		return x.ErrEmptyBatch
	}
	actionList := &pb.FlameBatchAction{FlameActionList: batch.mActionList}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_BATCH_ACTION,
	}

	if data, err := proto.Marshal(actionList); err == nil {
		pp.FlameProposalData = data
	} else {
		internalLogger.Error("data marshal error", zap.Error(err))
		return x.ErrDataMarshalError
	}

	r, err := m.managedSyncApplyProposal(m.mClusterID, pp, timeout)

	if err != nil {
		internalLogger.Error("proposal apply error", zap.Error(err))
		return x.ErrFailedToApplyProposal
	}

	if r.Value > 0 {
		batch.Reset()
		return nil
	} else {
		return x.ErrFailedToApplyProposal
	}
}

func (m *StorageManager) managedSyncApplyProposal(clusterID uint64,
	pp *pb.FlameProposal,
	timeout time.Duration) (sm.Result, error) {
	cmd, err := proto.Marshal(pp)
	if err != nil {
		return sm.Result{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	session := m.mDragonboatNodeHost.GetNoOPSession(clusterID)
	r, err := m.mDragonboatNodeHost.SyncPropose(ctx, session, cmd)
	cancel()

	_ = m.mDragonboatNodeHost.SyncCloseSession(context.Background(), session)

	return r, err
}

func (m *StorageManager) managedSyncRead(clusterID uint64, query interface{}, timeout time.Duration) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	d, e := m.mDragonboatNodeHost.SyncRead(ctx, clusterID, query)
	cancel()

	return d, e
}
