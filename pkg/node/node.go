package node

import (
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/config"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/x"
)

type Node struct {
	mIsNodeReady bool

	mNodeHost              *dragonboat.NodeHost
	mRaftConfiguration     config.Config
	mNodeHostConfiguration config.NodeHostConfig
	mNodeConfiguration     iface.INodeConfiguration

	mStoragedConfiguration iface.IStoragedConfiguration
}

func (n *Node) ConfigureNode(nodeConfiguration iface.INodeConfiguration,
	storagedConfiguration iface.IStoragedConfiguration) (bool, error) {

	if n.mIsNodeReady {
		return false, x.ErrNodeAlreadyConfigured
	}

	n.mNodeConfiguration = nodeConfiguration
	n.mStoragedConfiguration = storagedConfiguration

	return false, nil
}

func (n *Node) StartCluster(clusterConfiguration iface.IClusterConfiguration) error {
	if n.mNodeHost == nil {
		return x.ErrNodeIsNotReady
	}

	return nil
}

func (n *Node) StopCluster(clusterId uint64) error {
	if n.mNodeHost == nil {
		return x.ErrNodeIsNotReady
	}

	if err := n.mNodeHost.StopCluster(clusterId); err != nil {
		return x.ErrFailedToStopCluster
	}
	return nil
}

func (n *Node) StopNode() {
	if !n.mIsNodeReady {
		return
	}

	n.mNodeHost.Stop()

	n.mIsNodeReady = false

	n.mNodeHost = nil
	n.mNodeConfiguration = nil
	n.mStoragedConfiguration = nil
	n.mRaftConfiguration = config.Config{}
	n.mNodeHostConfiguration = config.NodeHostConfig{}
}
