package flamed

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/raftpb"
	sm "github.com/lni/dragonboat/v3/statemachine"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/variant"
	"github.com/mkawserm/flamed/pkg/x"
	"sync"
	"time"
)

type Node struct {
	mMutex       sync.Mutex
	mIsNodeReady bool

	mNodeHost              *dragonboat.NodeHost
	mNodeHostConfiguration config.NodeHostConfig
	mNodeConfiguration     iface.INodeConfiguration

	mClusterMap              map[uint64]string
	mClusterStorageTaskQueue map[uint64]variant.TaskQueue
}

func (n *Node) isStoragedConfigurationOk(storagedConfiguration iface.IStoragedConfiguration) bool {
	if storagedConfiguration.StateStoragePath() == "" {
		return false
	}
	if storagedConfiguration.StoragePluginState() == nil {
		return false
	}

	if storagedConfiguration.IndexEnable() {
		if storagedConfiguration.IndexStoragePath() == "" {
			return false
		}
		if storagedConfiguration.StoragePluginIndex() == nil {
			return false
		}
	}

	return true
}

func (n *Node) ConfigureNode(nodeConfiguration iface.INodeConfiguration) error {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	if n.mIsNodeReady {
		return x.ErrNodeAlreadyConfigured
	}

	if !utility.MkPath(nodeConfiguration.NodeHostDir()) {
		return x.ErrFailedToCreateNodeHostDir
	}

	if !utility.MkPath(nodeConfiguration.WALDir()) {
		return x.ErrFailedToCreateWALDir
	}

	n.mNodeConfiguration = nodeConfiguration

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
		EnableMetrics:                 nodeConfiguration.EnableMetrics(),
		MaxSnapshotSendBytesPerSecond: nodeConfiguration.MaxSnapshotSendBytesPerSecond(),
		MaxSnapshotRecvBytesPerSecond: nodeConfiguration.MaxSnapshotRecvBytesPerSecond(),

		NotifyCommit: nodeConfiguration.NotifyCommit(),

		LogDB:        nodeConfiguration.LogDB(),
		LogDBFactory: nodeConfiguration.LogDBFactory(),

		RaftRPCFactory:        nodeConfiguration.RaftRPCFactory(),
		RaftEventListener:     nodeConfiguration.RaftEventListener(),
		SystemEventListener:   nodeConfiguration.SystemEventListener(),
		SystemTickerPrecision: nodeConfiguration.SystemTickerPrecision(),
	}

	if nh, err := dragonboat.NewNodeHost(n.mNodeHostConfiguration); err != nil {
		return err
	} else {
		n.mNodeHost = nh
		n.mIsNodeReady = true
	}

	n.mClusterMap = make(map[uint64]string)
	n.mClusterStorageTaskQueue = make(map[uint64]variant.TaskQueue)

	return nil
}

func (n *Node) StartOnDiskCluster(clusterConfiguration iface.IOnDiskClusterConfiguration,
	storagedConfiguration iface.IStoragedConfiguration,
	raftConfiguration iface.IRaftConfiguration) error {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	if n.mNodeHost == nil {
		return x.ErrNodeIsNotReady
	}

	if !n.isStoragedConfigurationOk(storagedConfiguration) {
		return x.ErrInvalidStoragedConfiguration
	}
	//n.mStoragedConfiguration = storagedConfiguration

	rc := config.Config{
		NodeID:                  raftConfiguration.NodeID(),
		CheckQuorum:             raftConfiguration.CheckQuorum(),
		ElectionRTT:             raftConfiguration.ElectionRTT(),
		HeartbeatRTT:            raftConfiguration.HeartbeatRTT(),
		SnapshotEntries:         raftConfiguration.SnapshotEntries(),
		CompactionOverhead:      raftConfiguration.CompactionOverhead(),
		OrderedConfigChange:     raftConfiguration.OrderedConfigChange(),
		MaxInMemLogSize:         raftConfiguration.MaxInMemLogSize(),
		SnapshotCompressionType: raftpb.Snappy,
		EntryCompressionType:    raftpb.Snappy,
		DisableAutoCompactions:  raftConfiguration.DisableAutoCompactions(),
		IsObserver:              raftConfiguration.IsObserver(),
		IsWitness:               raftConfiguration.IsWitness(),
		Quiesce:                 raftConfiguration.Quiesce(),
	}

	rc.ClusterID = clusterConfiguration.ClusterID()

	err := n.mNodeHost.StartOnDiskCluster(clusterConfiguration.InitialMembers(),
		clusterConfiguration.Join(), func(clusterId uint64, nodeId uint64) sm.IOnDiskStateMachine {
			return clusterConfiguration.StateMachine(storagedConfiguration)(clusterId, nodeId)
		}, rc)

	if err != nil {
		return err
	}

	n.mClusterMap[rc.ClusterID] = clusterConfiguration.ClusterName()
	n.mClusterStorageTaskQueue[rc.ClusterID] = storagedConfiguration.StorageTaskQueue()
	//n.mNodeHost.GetNoOPSession(clusterConfiguration.ClusterID())

	return nil
}

func (n *Node) StopCluster(clusterID uint64) error {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	if n.mNodeHost == nil {
		return x.ErrNodeIsNotReady
	}

	if err := n.mNodeHost.StopCluster(clusterID); err != nil {
		return x.ErrFailedToStopCluster
	}

	delete(n.mClusterMap, clusterID)
	delete(n.mClusterStorageTaskQueue, clusterID)

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
	n.mNodeHostConfiguration = config.NodeHostConfig{}

	n.mClusterMap = make(map[uint64]string)
}

func (n *Node) TotalCluster() int {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	return len(n.mClusterMap)
}

func (n *Node) IsClusterIDAvailable(clusterID uint64) bool {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	_, found := n.mClusterMap[clusterID]
	return found
}

func (n *Node) ClusterIDList() []uint64 {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	var ids []uint64
	for k := range n.mClusterMap {
		ids = append(ids, k)
	}

	return ids
}

func (n *Node) NodeHostConfig() config.NodeHostConfig {
	return n.mNodeHost.NodeHostConfig()
}

func (n *Node) RaftAddress() string {
	return n.mNodeHost.RaftAddress()
}

func (n *Node) GetNodeHostInfo() *dragonboat.NodeHostInfo {
	return n.mNodeHost.GetNodeHostInfo(dragonboat.NodeHostInfoOption{SkipLogInfo: false})
}

func (n *Node) NewNodeAdmin(clusterID uint64, timeout time.Duration) *NodeAdmin {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	if _, ok := n.mClusterMap[clusterID]; !ok {
		return nil
	}

	return &NodeAdmin{
		mTimeout:            timeout,
		mClusterID:          clusterID,
		mDragonboatNodeHost: n.mNodeHost,
		mStorageTaskQueue:   n.mClusterStorageTaskQueue[clusterID],
	}
}

func (n *Node) managedSyncRead(clusterID uint64, query interface{}, timeout time.Duration) (interface{}, error) {
	if !n.IsClusterIDAvailable(clusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	d, e := n.mNodeHost.SyncRead(ctx, clusterID, query)
	cancel()

	return d, e
}

func (n *Node) managedSyncApplyProposal(clusterID uint64,
	pp *pb.Proposal,
	timeout time.Duration) (sm.Result, error) {
	if !n.IsClusterIDAvailable(clusterID) {
		return sm.Result{}, x.ErrClusterIsNotAvailable
	}

	cmd, err := proto.Marshal(pp)
	if err != nil {
		return sm.Result{}, err
	}

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	session := n.mNodeHost.GetNoOPSession(clusterID)
	r, err := n.mNodeHost.SyncPropose(ctx, session, cmd)
	cancel()

	_ = n.mNodeHost.SyncCloseSession(context.TODO(), session)

	return r, err
}
