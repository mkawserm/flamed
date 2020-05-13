package flamed

import (
	"context"
	"github.com/lni/dragonboat/v3"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
	"time"
)

type ClusterAdmin struct {
	mClusterID          uint64
	mDragonboatNodeHost *dragonboat.NodeHost
}

func (c *ClusterAdmin) GetLeaderID() (uint64, bool, error) {
	return c.mDragonboatNodeHost.GetLeaderID(c.mClusterID)
}

func (c *ClusterAdmin) HasNodeInfo(nodeID uint64) bool {
	return c.mDragonboatNodeHost.HasNodeInfo(c.mClusterID, nodeID)
}

func (c *ClusterAdmin) AddNode(nodeID uint64, address string, timeout time.Duration, configChangeIndex uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	err := c.mDragonboatNodeHost.SyncRequestAddNode(ctx,
		c.mClusterID,
		nodeID,
		address,
		configChangeIndex)
	cancel()
	return err
}

func (c *ClusterAdmin) AddObserver(nodeID uint64, address string, timeout time.Duration, configChangeIndex uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	err := c.mDragonboatNodeHost.SyncRequestAddObserver(ctx,
		c.mClusterID,
		nodeID,
		address,
		configChangeIndex)
	cancel()
	return err
}

func (c *ClusterAdmin) AddWitness(nodeID uint64, address string, timeout time.Duration, configChangeIndex uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	err := c.mDragonboatNodeHost.SyncRequestAddWitness(ctx,
		c.mClusterID,
		nodeID,
		address,
		configChangeIndex)
	cancel()
	return err
}

func (c *ClusterAdmin) DeleteNode(nodeID uint64, timeout time.Duration, configChangeIndex uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	err := c.mDragonboatNodeHost.SyncRequestDeleteNode(ctx,
		c.mClusterID,
		nodeID,
		configChangeIndex)
	cancel()
	return err
}

func (c *ClusterAdmin) RequestSnapshot(clusterID uint64,
	timeout time.Duration,
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

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	num, err := c.mDragonboatNodeHost.SyncRequestSnapshot(ctx, clusterID, opt)
	cancel()
	return num, err
}

func (c *ClusterAdmin) AppliedIndex(timeout time.Duration) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	request := variant.LookupRequest{
		Query:   pb.AppliedIndexQuery{},
		Context: ctx,
	}
	d, e := c.mDragonboatNodeHost.SyncRead(ctx, c.mClusterID, request)
	cancel()

	if e != nil {
		return 0, e
	}

	return d.(uint64), e
}
