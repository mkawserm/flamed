package storaged

import (
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/x"
	"io"
)

import "github.com/mkawserm/flamed/pkg/storage"
import sm "github.com/mkawserm/dragonboat/v3/statemachine"

type Storaged struct {
	mStorage *storage.Storage

	mNodeId      uint64
	mClusterId   uint64
	mLastApplied uint64

	mStoragedConfiguration iface.IStoragedConfiguration
}

func (s *Storaged) GetStorage() iface.IStorage {
	return s.mStorage
}

func (s *Storaged) SetConfiguration(configuration iface.IStoragedConfiguration) bool {
	if s.mStorage != nil {
		return false
	} else {
		s.mStorage = &storage.Storage{}
	}

	s.mStoragedConfiguration = configuration
	b := s.mStorage.SetConfiguration(configuration)

	if !b {
		s.mStoragedConfiguration = nil
	}

	return b
}

func (s *Storaged) saveAppliedIndex(u uint64) error {
	if s.mStorage == nil {
		return x.ErrStorageIsNotReady
	}

	return s.mStorage.SaveAppliedIndex(u)
}

func (s *Storaged) queryAppliedIndex() (uint64, error) {
	if s.mStorage == nil {
		return 0, x.ErrStorageIsNotReady
	}

	return s.mStorage.QueryAppliedIndex()
}

func (s *Storaged) Open(<-chan struct{}) (uint64, error) {
	if s.mStorage == nil {
		return 0, x.ErrStorageIsNotReady
	}

	err := s.mStorage.Open()

	if err != nil {
		return 0, err
	}

	if appliedIndex, err := s.queryAppliedIndex(); err != nil {
		return 0, err
	} else {
		s.mLastApplied = appliedIndex
	}

	return s.mLastApplied, nil
}

func (s *Storaged) Sync() error {
	return nil
}

func (s *Storaged) Close() error {
	if s.mStorage == nil {
		return nil
	}

	s.mStoragedConfiguration = nil
	return s.mStorage.Close()
}

func (s *Storaged) Update(entries []sm.Entry) ([]sm.Entry, error) {
	if s.mStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	for idx, e := range entries {
		pp := &pb.FlameProposal{}

		if err := proto.Unmarshal(e.Cmd, pp); err != nil {
			return nil, err
		}

		if err := s.mStorage.ApplyProposal(pp, false); err == nil {
			entries[idx].Result = sm.Result{Value: uint64(len(entries[idx].Cmd))}
		} else {
			return nil, err
		}
	}

	// save the applied index to the DB.
	idx := entries[len(entries)-1].Index
	err := s.saveAppliedIndex(idx)

	if err != nil {
		return nil, err
	}

	if s.mLastApplied >= entries[len(entries)-1].Index {
		return nil, x.ErrLastIndexIsNotMovingForward
	}

	s.mLastApplied = entries[len(entries)-1].Index

	return entries, nil
}

func (s *Storaged) Lookup(input interface{}) (interface{}, error) {
	if s.mStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	return s.mStorage.Lookup(input, false)
}

func (s *Storaged) PrepareSnapshot() (interface{}, error) {
	if s.mStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	return s.mStorage.PrepareSnapshot()
}

func (s *Storaged) SaveSnapshot(snapshotContext interface{}, w io.Writer, _ <-chan struct{}) error {
	if s.mStorage == nil {
		return x.ErrStorageIsNotReady
	}
	return s.mStorage.SaveSnapshot(snapshotContext, w)
}

func (s *Storaged) RecoverFromSnapshot(r io.Reader, _ <-chan struct{}) error {
	if s.mStorage == nil {
		return x.ErrStorageIsNotReady
	}

	if err := s.mStorage.RecoverFromSnapshot(r); err != nil {
		return err
	}

	// update the last applied index from the DB.
	newLastApplied, err := s.queryAppliedIndex()
	if err != nil || newLastApplied == 0 {
		return x.ErrLastIndexIsNotMovingForward
	}

	if s.mLastApplied > newLastApplied {
		return x.ErrLastIndexIsNotMovingForward
	}

	s.mLastApplied = newLastApplied

	return nil
}

func NewStoraged(configuration iface.IStoragedConfiguration) func(
	clusterId uint64,
	nodeId uint64) sm.IOnDiskStateMachine {
	return func(clusterId uint64, nodeId uint64) sm.IOnDiskStateMachine {
		sd := &Storaged{}
		sd.SetConfiguration(configuration)
		sd.mClusterId = clusterId
		sd.mNodeId = nodeId
		return sd
	}
}
