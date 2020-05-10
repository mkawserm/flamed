package flamed

import (
	"github.com/mkawserm/dragonboat/v3"
	"github.com/mkawserm/dragonboat/v3/config"
	"github.com/mkawserm/flamed/pkg/iface"
)

type Flamed struct {
	mNodeHost *NodeHost
}

func (f *Flamed) ConfigureNode(nodeConfiguration iface.INodeConfiguration,
	storagedConfiguration iface.IStoragedConfiguration) error {
	return f.mNodeHost.ConfigureNode(nodeConfiguration, storagedConfiguration)
}

func (f *Flamed) StartCluster(clusterConfiguration iface.IClusterConfiguration) error {
	return f.mNodeHost.StartCluster(clusterConfiguration)
}

func (f *Flamed) StopCluster(clusterID uint64) error {
	return f.mNodeHost.StopCluster(clusterID)
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

func (f *Flamed) NewClusterAdmin(clusterID uint64) *ClusterAdmin {
	return f.mNodeHost.NewClusterAdmin(clusterID)
}

func (f *Flamed) NewAdmin(clusterID uint64) *Admin {
	return f.mNodeHost.NewAdmin(clusterID)
}

func (f *Flamed) NewStorageManager(clusterID uint64) *StorageManager {
	return f.mNodeHost.NewStorageManager(clusterID)
}

func NewFlamed() *Flamed {
	return &Flamed{mNodeHost: &NodeHost{}}
}
