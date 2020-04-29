package storaged

import (
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/x"
	"io"
)

import "github.com/mkawserm/flamed/pkg/storage"
import sm "github.com/lni/dragonboat/v3/statemachine"

const appliedIndexNamespace string = "_l"
const appliedIndexKey string = "_l"

type Storaged struct {
	mStorage *storage.Storage

	mNodeId      uint64
	mClusterId   uint64
	mLastApplied uint64

	mStoragedConfiguration iface.IStoragedConfiguration
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

func (s *Storaged) saveAppliedIndex(u uint64) (bool, error) {
	if s.mStorage == nil {
		return false, x.ErrStorageIsNotReady
	}

	return s.mStorage.Create(
		[]byte(appliedIndexNamespace),
		[]byte(appliedIndexKey),
		uidutil.Uint64ToByteSlice(u))
}

func (s *Storaged) queryAppliedIndex() (uint64, error) {
	if s.mStorage == nil {
		return 0, x.ErrStorageIsNotReady
	}

	data, err := s.mStorage.Read(
		[]byte(appliedIndexNamespace),
		[]byte(appliedIndexKey))

	if err == x.ErrUidDoesNotExists {
		return 0, nil
	}

	if err == x.ErrFailedToReadDataFromStorage || err != nil {
		return 0, err
	}

	return uidutil.ByteSliceToUint64(data), nil
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
		batch := &pb.FlameBatch{}

		if err := proto.Unmarshal(e.Cmd, batch); err != nil {
			return nil, err
		}

		if b, err := s.mStorage.ApplyBatch(batch); b {
			entries[idx].Result = sm.Result{Value: uint64(len(entries[idx].Cmd))}
		} else {
			return nil, err
		}
	}

	// save the applied index to the DB.
	idx := entries[len(entries)-1].Index
	_, err := s.saveAppliedIndex(idx)

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

	if v, ok := input.(pb.FlameEntry); ok {
		if len(v.Namespace) < 3 {
			return nil, x.ErrInvalidLookupInput
		}

		return s.mStorage.Read(v.Namespace, v.Key)
	}

	if v, ok := input.(*pb.FlameEntry); ok {
		if len(v.Namespace) < 3 {
			return nil, x.ErrInvalidLookupInput
		}

		return s.mStorage.Read(v.Namespace, v.Key)
	}

	return nil, x.ErrInvalidLookupInput
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

	if snapshot, ok := snapshotContext.(iface.IKVStorage); ok {
		return snapshot.SaveSnapshot(w)
	} else {
		return x.ErrInvalidSnapshotContext
	}
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
