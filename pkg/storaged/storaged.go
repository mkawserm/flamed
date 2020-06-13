package storaged

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/storage"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"io"
)

import sm "github.com/lni/dragonboat/v3/statemachine"

type Storaged struct {
	mStorage iface.IStorage

	mNodeId      uint64
	mClusterId   uint64
	mLastApplied uint64

	mStoragedConfiguration iface.IStoragedConfiguration

	//mMutex sync.Mutex
}

func (s *Storaged) Setup(clusterID uint64, nodeID uint64) {
	s.mClusterId = clusterID
	s.mNodeId = nodeID
}

func (s *Storaged) SetStorage(storage iface.IStorage) {
	if s.mStorage == nil {
		s.mStorage = storage
	}
}

func (s *Storaged) SetConfiguration(configuration iface.IStoragedConfiguration) bool {
	if s.mStorage == nil {
		return false
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
	//s.mMutex.Lock()
	//defer s.mMutex.Unlock()

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
	//s.mMutex.Lock()
	//defer s.mMutex.Unlock()

	logger.L("storaged").Debug("storaged closing")
	defer func() {
		logger.L("storaged").Debug("storaged closed")
	}()
	if s.mStorage == nil {
		return nil
	}

	s.mStoragedConfiguration = nil
	return s.mStorage.Close()
}

func (s *Storaged) Update(entries []sm.Entry) ([]sm.Entry, error) {
	logger.L("storaged").Debug("storaged update")
	if s.mStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	logger.L("storaged").Debug("", zap.Int("entriesLength", len(entries)))
	for idx, e := range entries {
		entries[idx].Result = sm.Result{Value: uint64(len(entries[idx].Cmd))}

		pp := &pb.Proposal{}
		if err := proto.Unmarshal(e.Cmd, pp); err != nil {
			logger.L("storaged").Error("proto unmarshal error", zap.Error(err))
			continue
		}
		pr := s.mStorage.ApplyProposal(context.TODO(), pp, e.Index)
		if pr != nil {
			if data, err := proto.Marshal(pr); err == nil {
				entries[idx].Result.Data = data
			} else {
				logger.L("storaged").Error("proto marshal error", zap.Error(err))
			}
		}
	}

	// save the applied indexmeta to the DB.
	idx := entries[len(entries)-1].Index
	err := s.saveAppliedIndex(idx)

	if err != nil {
		return nil, err
	}

	if s.mLastApplied >= entries[len(entries)-1].Index {
		return nil, x.ErrLastIndexIsNotMovingForward
	}

	s.mLastApplied = entries[len(entries)-1].Index

	logger.L("storaged").Debug("storaged update done")
	return entries, nil
}

func (s *Storaged) Lookup(input interface{}) (interface{}, error) {
	logger.L("storaged").Debug("storaged lookup")
	if s.mStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	switch v := input.(type) {
	case *pb.SearchInput:
		r, err := s.mStorage.Search(context.TODO(), v)
		logger.L("storaged").Debug("storaged lookup done with [pb.SearchInput]")
		return r, err
	case *pb.IterateInput:
		r, err := s.mStorage.Iterate(context.TODO(), v)
		logger.L("storaged").Debug("storaged lookup done with [pb.IterateInput]")
		return r, err
	case *pb.RetrieveInput:
		r, err := s.mStorage.Retrieve(context.TODO(), v)
		logger.L("storaged").Debug("storaged lookup done with [pb.RetrieveInput]")
		return r, err
	case *pb.GlobalSearchInput:
		r, err := s.mStorage.GlobalSearch(context.TODO(), v)
		logger.L("storaged").Debug("storaged lookup done with [pb.GlobalSearchInput]")
		return r, err
	case *pb.GlobalIterateInput:
		r, err := s.mStorage.GlobalIterate(context.TODO(), v)
		logger.L("storaged").Debug("storaged lookup done with [pb.GlobalIterateInput]")
		return r, err
	case *pb.GlobalRetrieveInput:
		r, err := s.mStorage.GlobalRetrieve(context.TODO(), v)
		logger.L("storaged").Debug("storaged lookup done with [pb.GlobalRetrieveInput]")
		return r, err
	case *pb.AppliedIndexQuery:
		r, err := s.mStorage.GetAppliedIndex(context.TODO())
		logger.L("storaged").Debug("storaged lookup done with [pb.AppliedIndexQuery]")
		return r, err
	default:
		logger.L("storaged").Debug("storaged lookup done")
		return nil, x.ErrInvalidLookupInput
	}
}

func (s *Storaged) PrepareSnapshot() (interface{}, error) {
	//s.mMutex.Lock()
	//defer s.mMutex.Unlock()
	if s.mStorage == nil {
		return nil, x.ErrStorageIsNotReady
	}

	return s.mStorage.PrepareSnapshot()
}

func (s *Storaged) SaveSnapshot(snapshotContext interface{}, w io.Writer, _ <-chan struct{}) error {
	//s.mMutex.Lock()
	//defer s.mMutex.Unlock()
	if s.mStorage == nil {
		return x.ErrStorageIsNotReady
	}
	return s.mStorage.SaveSnapshot(snapshotContext, w)
}

func (s *Storaged) RecoverFromSnapshot(r io.Reader, _ <-chan struct{}) error {
	//s.mMutex.Lock()
	//defer s.mMutex.Unlock()
	if s.mStorage == nil {
		return x.ErrStorageIsNotReady
	}

	if err := s.mStorage.RecoverFromSnapshot(r); err != nil {
		return err
	}

	// update the last applied indexmeta from the DB.
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
	clusterID uint64,
	nodeID uint64) iface.IStoraged {
	return func(clusterId uint64, nodeId uint64) iface.IStoraged {
		sd := &Storaged{}
		sd.Setup(clusterId, nodeId)
		sd.SetStorage(&storage.Storage{})
		sd.SetConfiguration(configuration)
		return sd
	}
}
