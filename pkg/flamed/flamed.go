package flamed

import (
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/config"
	sm "github.com/lni/dragonboat/v3/statemachine"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"time"
)

type Flamed struct {
	mNodeHost *Node
}

func (f *Flamed) ConfigureNode(nodeConfiguration iface.INodeConfiguration) error {
	return f.mNodeHost.ConfigureNode(nodeConfiguration)
}

func (f *Flamed) StartOnDiskCluster(clusterConfiguration iface.IOnDiskClusterConfiguration,
	storagedConfiguration iface.IStoragedConfiguration,
	raftConfiguration iface.IRaftConfiguration) error {
	return f.mNodeHost.StartOnDiskCluster(clusterConfiguration,
		storagedConfiguration,
		raftConfiguration)
}

func (f *Flamed) StopCluster(clusterID uint64) error {
	return f.mNodeHost.StopCluster(clusterID)
}

func (f *Flamed) IsClusterIDAvailable(clusterID uint64) bool {
	return f.mNodeHost.IsClusterIDAvailable(clusterID)
}

func (f *Flamed) StopNode() {
	f.mNodeHost.StopNode()
}

func (f *Flamed) TotalCluster() int {
	return f.mNodeHost.TotalCluster()
}

func (f *Flamed) ClusterIDList() []uint64 {
	return f.mNodeHost.ClusterIDList()
}

func (f *Flamed) NodeHostConfig() config.NodeHostConfig {
	return f.mNodeHost.NodeHostConfig()
}

func (f *Flamed) RaftAddress() string {
	return f.mNodeHost.RaftAddress()
}

func (f *Flamed) GetNodeHostInfo() *dragonboat.NodeHostInfo {
	return f.mNodeHost.GetNodeHostInfo()
}

func (f *Flamed) NewNodeAdmin(clusterID uint64, timeout time.Duration) *NodeAdmin {
	return f.mNodeHost.NewNodeAdmin(clusterID, timeout)
}

func (f *Flamed) NewAdmin(clusterID uint64, timeout time.Duration) *Admin {
	return &Admin{
		mRW:        f,
		mClusterID: clusterID,
		mTimeout:   timeout,
	}
}

func (f *Flamed) NewGlobalOperation(clusterID uint64, namespace []byte, timeout time.Duration) *GlobalOperation {
	q := &GlobalOperation{}
	if err := q.Setup(clusterID, namespace, f, timeout); err != nil {
		return nil
	}

	return q
}

func (f *Flamed) Read(clusterID uint64, query interface{}, timeout time.Duration) (interface{}, error) {
	return f.mNodeHost.managedSyncRead(clusterID, query, timeout)
}

func (f *Flamed) Write(clusterID uint64, pp *pb.Proposal, timeout time.Duration) (sm.Result, error) {
	return f.mNodeHost.managedSyncApplyProposal(clusterID, pp, timeout)
}

func NewFlamed() *Flamed {
	return &Flamed{mNodeHost: &Node{}}
}
