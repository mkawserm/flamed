package storaged

import (
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/x"
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
	s.mStoragedConfiguration = configuration
	b := s.mStorage.SetConfiguration(configuration)

	if !b {
		s.mStoragedConfiguration = nil
	}

	return b
}

func (s *Storaged) saveAppliedIndex(u uint64) (bool, error) {
	return s.mStorage.Create(
		[]byte(appliedIndexNamespace),
		[]byte(appliedIndexKey),
		uidutil.Uint64ToByteSlice(u))
}

func (s *Storaged) queryAppliedIndex() (uint64, error) {
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
	_, err := s.mStorage.Open()

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
	return s.mStorage.Close()
}

func (s *Storaged) Update(entries []sm.Entry) ([]sm.Entry, error) {
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

//PrepareSnapshot() (interface{}, error)
//SaveSnapshot(interface{}, io.Writer, <-chan struct{}) error
//RecoverFromSnapshot(io.Reader, <-chan struct{}) error

func NewStoraged(configuration iface.IStoragedConfiguration) func(
	clusterId uint64,
	nodeId uint64) sm.IOnDiskStateMachine {

	sd := &Storaged{}
	if sd.SetConfiguration(configuration) {
		return func(clusterId uint64, nodeId uint64) sm.IOnDiskStateMachine {
			sd.mClusterId = clusterId
			sd.mNodeId = nodeId
			return nil
		}
	} else {
		return nil
	}
}
