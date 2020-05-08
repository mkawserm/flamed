package flamed

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/lni/dragonboat/v3"
	sm "github.com/lni/dragonboat/v3/statemachine"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"time"
)

type StorageManager struct {
	mClusterID          uint64
	mDragonboatNodeHost *dragonboat.NodeHost
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
