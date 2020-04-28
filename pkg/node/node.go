package node

import (
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/plugin/pebble"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/x"
)

type Node struct {
	mNodeHost              *dragonboat.NodeHost
	mRaftConfiguration     config.Config
	mNodeHostConfiguration config.NodeHostConfig
	mNodeConfiguration     iface.INodeConfiguration
}

func (n *Node) ConfigureNode(nodeConfiguration iface.INodeConfiguration) (bool, error) {
	if n.mNodeHost != nil {
		return false, x.ErrNodeAlreadyConfigured
	}
	n.mNodeConfiguration = nodeConfiguration

	n.mNodeHostConfiguration = config.NodeHostConfig{
		DeploymentID:                  0,
		WALDir:                        "",
		NodeHostDir:                   "",
		RTTMillisecond:                0,
		RaftAddress:                   "",
		ListenAddress:                 "",
		MutualTLS:                     false,
		CAFile:                        "",
		CertFile:                      "",
		KeyFile:                       "",
		MaxSendQueueSize:              0,
		MaxReceiveQueueSize:           0,
		LogDBFactory:                  pebble.NewLogDB,
		RaftRPCFactory:                nil,
		EnableMetrics:                 false,
		RaftEventListener:             nil,
		MaxSnapshotSendBytesPerSecond: 0,
		MaxSnapshotRecvBytesPerSecond: 0,
		FS:                            nil,
		SystemEventListener:           nil,
		SystemTickerPrecision:         0,
	}

	//n.mNodeHostConfiguration.LogDBFactory = pebble.NewLogDB

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
	if n.mNodeHost == nil {
		return
	}

	n.mNodeHost.Stop()

	n.mNodeHost = nil
	n.mNodeConfiguration = nil
}
