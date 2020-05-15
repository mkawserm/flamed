package flamed

import (
	"context"
	"fmt"
	"github.com/lni/dragonboat/v3"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

type ClusterAdmin struct {
	mClusterID          uint64
	mTimeout            time.Duration
	mStorageTaskQueue   variant.TaskQueue
	mDragonboatNodeHost *dragonboat.NodeHost
}

func (c *ClusterAdmin) UpdateTimeout(timeout time.Duration) {
	c.mTimeout = timeout
}

func (c *ClusterAdmin) GetLeaderID() (uint64, bool, error) {
	return c.mDragonboatNodeHost.GetLeaderID(c.mClusterID)
}

func (c *ClusterAdmin) HasNodeInfo(nodeID uint64) bool {
	return c.mDragonboatNodeHost.HasNodeInfo(c.mClusterID, nodeID)
}

func (c *ClusterAdmin) AddNode(nodeID uint64, address string, configChangeIndex uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.mTimeout)
	err := c.mDragonboatNodeHost.SyncRequestAddNode(ctx,
		c.mClusterID,
		nodeID,
		address,
		configChangeIndex)
	cancel()
	return err
}

func (c *ClusterAdmin) AddObserver(nodeID uint64, address string, configChangeIndex uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.mTimeout)
	err := c.mDragonboatNodeHost.SyncRequestAddObserver(ctx,
		c.mClusterID,
		nodeID,
		address,
		configChangeIndex)
	cancel()
	return err
}

func (c *ClusterAdmin) AddWitness(nodeID uint64, address string, configChangeIndex uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.mTimeout)
	err := c.mDragonboatNodeHost.SyncRequestAddWitness(ctx,
		c.mClusterID,
		nodeID,
		address,
		configChangeIndex)
	cancel()
	return err
}

func (c *ClusterAdmin) DeleteNode(nodeID uint64, configChangeIndex uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.mTimeout)
	err := c.mDragonboatNodeHost.SyncRequestDeleteNode(ctx,
		c.mClusterID,
		nodeID,
		configChangeIndex)
	cancel()
	return err
}

func (c *ClusterAdmin) RequestSnapshot(clusterID uint64,
	compactionOverhead uint64,
	exportPath string,
	exported bool,
	overrideCompactionOverhead bool) (uint64, error) {

	opt := dragonboat.SnapshotOption{
		CompactionOverhead:         compactionOverhead,
		ExportPath:                 exportPath,
		Exported:                   exported,
		OverrideCompactionOverhead: overrideCompactionOverhead,
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.mTimeout)
	num, err := c.mDragonboatNodeHost.SyncRequestSnapshot(ctx, clusterID, opt)
	cancel()
	return num, err
}

func (c *ClusterAdmin) GetAppliedIndex() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.mTimeout)
	request := variant.LookupRequest{
		Query:   pb.AppliedIndexQuery{},
		Context: ctx,
	}
	d, e := c.mDragonboatNodeHost.SyncRead(ctx, c.mClusterID, request)
	cancel()

	if e != nil {
		return 0, e
	}

	if v, ok := d.(uint64); ok {
		return v, nil
	} else {
		return 0, x.ErrUnknownValue
	}
}

func (c *ClusterAdmin) RunStorageGC() {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if c.mStorageTaskQueue == nil {
		internalLogger.Debug("storage task queue is nil")
		return
	}

	c.mStorageTaskQueue <- variant.Task{
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:    "storage-task",
		Command: "gc",
	}
}

func (c *ClusterAdmin) FullIndex() {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if c.mStorageTaskQueue == nil {
		internalLogger.Debug("storage task queue is nil")
		return
	}

	c.mStorageTaskQueue <- variant.Task{
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:    "storage-task",
		Command: "full-index",
	}
}
