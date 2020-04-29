package node

import (
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/client"
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/raftpb"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"sync"
)

type Node struct {
	mMutex       sync.Mutex
	mIsNodeReady bool

	mNodeHost              *dragonboat.NodeHost
	mRaftConfiguration     config.Config
	mNodeHostConfiguration config.NodeHostConfig
	mNodeConfiguration     iface.INodeConfiguration

	mStoragedConfiguration iface.IStoragedConfiguration

	mClusterMap        map[uint64]string
	mClusterSessionMap map[uint64]*client.Session
}

func (n *Node) isStoragedConfigurationOk(storagedConfiguration iface.IStoragedConfiguration) bool {
	if storagedConfiguration.StoragePath() == "" {
		return false
	}

	if storagedConfiguration.StoragePluginIndex() == nil {
		return false
	}

	if storagedConfiguration.StoragePluginKV() == nil {
		return false
	}

	return true
}

func (n *Node) ConfigureNode(nodeConfiguration iface.INodeConfiguration,
	storagedConfiguration iface.IStoragedConfiguration) error {

	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	if n.mIsNodeReady {
		return x.ErrNodeAlreadyConfigured
	}

	if !n.isStoragedConfigurationOk(storagedConfiguration) {
		return x.ErrInvalidStoragedConfiguration
	}

	if !utility.MkPath(nodeConfiguration.NodeHostDir()) {
		return x.ErrFailedToCreateNodeHostDir
	}

	if !utility.MkPath(nodeConfiguration.WALDir()) {
		return x.ErrFailedToCreateWALDir
	}

	n.mNodeConfiguration = nodeConfiguration
	n.mStoragedConfiguration = storagedConfiguration

	n.mRaftConfiguration = config.Config{
		NodeID:                  nodeConfiguration.NodeID(),
		CheckQuorum:             nodeConfiguration.CheckQuorum(),
		ElectionRTT:             nodeConfiguration.ElectionRTT(),
		HeartbeatRTT:            nodeConfiguration.HeartbeatRTT(),
		SnapshotEntries:         nodeConfiguration.SnapshotEntries(),
		CompactionOverhead:      nodeConfiguration.CompactionOverhead(),
		OrderedConfigChange:     nodeConfiguration.OrderedConfigChange(),
		MaxInMemLogSize:         nodeConfiguration.MaxInMemLogSize(),
		SnapshotCompressionType: raftpb.Snappy,
		EntryCompressionType:    raftpb.Snappy,
		DisableAutoCompactions:  nodeConfiguration.DisableAutoCompactions(),
		IsObserver:              nodeConfiguration.IsObserver(),
		IsWitness:               nodeConfiguration.IsWitness(),
		Quiesce:                 nodeConfiguration.Quiesce(),
	}

	n.mNodeHostConfiguration = config.NodeHostConfig{
		DeploymentID:                  nodeConfiguration.DeploymentID(),
		WALDir:                        nodeConfiguration.WALDir(),
		NodeHostDir:                   nodeConfiguration.NodeHostDir(),
		RTTMillisecond:                nodeConfiguration.RTTMillisecond(),
		RaftAddress:                   nodeConfiguration.RaftAddress(),
		ListenAddress:                 nodeConfiguration.ListenAddress(),
		MutualTLS:                     nodeConfiguration.MutualTLS(),
		CAFile:                        nodeConfiguration.CAFile(),
		CertFile:                      nodeConfiguration.CertFile(),
		KeyFile:                       nodeConfiguration.KeyFile(),
		MaxSendQueueSize:              nodeConfiguration.MaxSendQueueSize(),
		MaxReceiveQueueSize:           nodeConfiguration.MaxReceiveQueueSize(),
		LogDBFactory:                  nodeConfiguration.LogDBFactory(),
		RaftRPCFactory:                nodeConfiguration.RaftRPCFactory(),
		EnableMetrics:                 nodeConfiguration.EnableMetrics(),
		MaxSnapshotSendBytesPerSecond: nodeConfiguration.MaxSnapshotSendBytesPerSecond(),
		MaxSnapshotRecvBytesPerSecond: nodeConfiguration.MaxSnapshotRecvBytesPerSecond(),
	}

	if nh, err := dragonboat.NewNodeHost(n.mNodeHostConfiguration); err != nil {
		return err
	} else {
		n.mNodeHost = nh
		n.mIsNodeReady = true
	}

	n.mClusterMap = make(map[uint64]string)
	n.mClusterSessionMap = make(map[uint64]*client.Session)

	return nil
}

func (n *Node) StartCluster(clusterConfiguration iface.IClusterConfiguration) error {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	if n.mNodeHost == nil {
		return x.ErrNodeIsNotReady
	}

	n.mRaftConfiguration.ClusterID = clusterConfiguration.ClusterId()

	err := n.mNodeHost.StartOnDiskCluster(clusterConfiguration.InitialMembers(),
		clusterConfiguration.Join(),
		clusterConfiguration.StateMachine(n.mStoragedConfiguration), n.mRaftConfiguration)

	if err != nil {
		n.mRaftConfiguration.ClusterID = 0
		return err
	}

	n.mClusterMap[n.mRaftConfiguration.ClusterID] = clusterConfiguration.ClusterName()
	n.mClusterSessionMap[n.mRaftConfiguration.ClusterID] = n.mNodeHost.GetNoOPSession(clusterConfiguration.ClusterId())

	return nil
}

func (n *Node) StopCluster(clusterId uint64) error {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	if n.mNodeHost == nil {
		return x.ErrNodeIsNotReady
	}

	if err := n.mNodeHost.StopCluster(clusterId); err != nil {
		return x.ErrFailedToStopCluster
	}

	delete(n.mClusterMap, clusterId)
	delete(n.mClusterSessionMap, clusterId)

	return nil
}

func (n *Node) StopNode() {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

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

	n.mClusterMap = make(map[uint64]string)
	n.mClusterSessionMap = make(map[uint64]*client.Session)
}

func (n *Node) TotalCluster() int {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return len(n.mClusterMap)
}

func (n *Node) GetDragonboatNodeHost() *dragonboat.NodeHost {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.mNodeHost
}
