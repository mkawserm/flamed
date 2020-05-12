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
	"github.com/mkawserm/flamed/pkg/x"
	"sync"
	"time"
)

type NodeHost struct {
	mMutex       sync.Mutex
	mIsNodeReady bool

	mNodeHost              *dragonboat.NodeHost
	mRaftConfiguration     config.Config
	mNodeHostConfiguration config.NodeHostConfig
	mNodeConfiguration     iface.INodeConfiguration

	mStoragedConfiguration iface.IStoragedConfiguration

	mClusterMap map[uint64]string
}

func (n *NodeHost) isStoragedConfigurationOk(storagedConfiguration iface.IStoragedConfiguration) bool {
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

func (n *NodeHost) ConfigureNode(nodeConfiguration iface.INodeConfiguration,
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
		EnableMetrics:                 nodeConfiguration.EnableMetrics(),
		MaxSnapshotSendBytesPerSecond: nodeConfiguration.MaxSnapshotSendBytesPerSecond(),
		MaxSnapshotRecvBytesPerSecond: nodeConfiguration.MaxSnapshotRecvBytesPerSecond(),

		LogDBFactory:          nodeConfiguration.LogDBFactory(),
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

	return nil
}

func (n *NodeHost) StartCluster(clusterConfiguration iface.IClusterConfiguration) error {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	if n.mNodeHost == nil {
		return x.ErrNodeIsNotReady
	}

	n.mRaftConfiguration.ClusterID = clusterConfiguration.ClusterID()

	err := n.mNodeHost.StartOnDiskCluster(clusterConfiguration.InitialMembers(),
		clusterConfiguration.Join(),
		clusterConfiguration.StateMachine(n.mStoragedConfiguration), n.mRaftConfiguration)

	if err != nil {
		n.mRaftConfiguration.ClusterID = 0
		return err
	}

	n.mClusterMap[n.mRaftConfiguration.ClusterID] = clusterConfiguration.ClusterName()
	//n.mNodeHost.GetNoOPSession(clusterConfiguration.ClusterID())

	return nil
}

func (n *NodeHost) StopCluster(clusterID uint64) error {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	if n.mNodeHost == nil {
		return x.ErrNodeIsNotReady
	}

	if err := n.mNodeHost.StopCluster(clusterID); err != nil {
		return x.ErrFailedToStopCluster
	}

	delete(n.mClusterMap, clusterID)

	return nil
}

func (n *NodeHost) StopNode() {
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
}

func (n *NodeHost) TotalCluster() int {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	return len(n.mClusterMap)
}

func (n *NodeHost) IsClusterIDExists(clusterID uint64) bool {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	_, found := n.mClusterMap[clusterID]
	return found
}

func (n *NodeHost) ClusterIDList() []uint64 {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()

	var ids []uint64
	for k := range n.mClusterMap {
		ids = append(ids, k)
	}

	return ids
}

func (n *NodeHost) NodeHostConfig() config.NodeHostConfig {
	return n.mNodeHost.NodeHostConfig()
}

func (n *NodeHost) RaftAddress() string {
	return n.mNodeHost.RaftAddress()
}

func (n *NodeHost) GetNodeHostInfo() *dragonboat.NodeHostInfo {
	return n.mNodeHost.GetNodeHostInfo(dragonboat.NodeHostInfoOption{SkipLogInfo: false})
}

func (n *NodeHost) NewClusterAdmin(clusterID uint64) *ClusterAdmin {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	if _, ok := n.mClusterMap[clusterID]; !ok {
		return nil
	}

	return &ClusterAdmin{
		mClusterID:          clusterID,
		mDragonboatNodeHost: n.mNodeHost,
	}
}

func (n *NodeHost) managedSyncRead(clusterID uint64, query interface{}, timeout time.Duration) (interface{}, error) {
	if !n.IsClusterIDExists(clusterID) {
		return nil, x.ErrClusterNotFound
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	d, e := n.mNodeHost.SyncRead(ctx, clusterID, query)
	cancel()

	return d, e
}

func (n *NodeHost) managedSyncApplyProposal(clusterID uint64,
	pp *pb.Proposal,
	timeout time.Duration) (sm.Result, error) {
	if !n.IsClusterIDExists(clusterID) {
		return sm.Result{}, x.ErrClusterNotFound
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

//func (n *NodeHost) GetLeaderID(clusterID uint64) (uint64, bool, error) {
//	return n.mNodeHost.GetLeaderID(clusterID)
//}

//func (n *NodeHost) HasNodeInfo(clusterID uint64, nodeID uint64) bool {
//	return n.mNodeHost.HasNodeInfo(clusterID, nodeID)
//}

//func (n *NodeHost) ManagedSyncRequestAddNode(clusterID uint64,
//	nodeID uint64,
//	address string,
//	timeout time.Duration) error {
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	err := n.mNodeHost.SyncRequestAddNode(ctx, clusterID, nodeID, address, 0)
//	cancel()
//	return err
//}

//func (n *NodeHost) ManagedSyncRequestAddObserver(clusterID uint64, nodeID uint64,
//	address string, timeout time.Duration) error {
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	err := n.mNodeHost.SyncRequestAddObserver(ctx, clusterID, nodeID, address, 0)
//	cancel()
//	return err
//}

//func (n *NodeHost) ManagedSyncRequestAddWitness(clusterID uint64, nodeID uint64,
//	address string, timeout time.Duration) error {
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	err := n.mNodeHost.SyncRequestAddWitness(ctx, clusterID, nodeID, address, 0)
//	cancel()
//	return err
//}

//func (n *NodeHost) ManagedSyncRequestDeleteNode(clusterID uint64,
//	nodeID uint64,
//	timeout time.Duration) error {
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	err := n.mNodeHost.SyncRequestDeleteNode(ctx, clusterID, nodeID, 0)
//	cancel()
//	return err
//}

//func (n *NodeHost) ManagedSyncRequestSnapshot(clusterID uint64,
//	opt dragonboat.SnapshotOption,
//	timeout time.Duration) (uint64, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	num, err := n.mNodeHost.SyncRequestSnapshot(ctx, clusterID, opt)
//	cancel()
//	return num, err
//}

//func (n *NodeHost) GetDragonboatNodeHost() *dragonboat.NodeHost {
//	n.mMutex.Lock()
//	defer n.mMutex.Unlock()
//
//	return n.mNodeHost
//}

//func (n *NodeHost) SyncGetClusterMembership(ctx context.Context, clusterID uint64) (*dragonboat.Membership, error) {
//	return n.mNodeHost.SyncGetClusterMembership(ctx, clusterID)
//}

//func (n *NodeHost) SyncRead(ctx context.Context, clusterID uint64, query interface{}) (interface{}, error) {
//	return n.mNodeHost.SyncRead(ctx, clusterID, query)
//}

//func (n *NodeHost) SyncApplyProposal(ctx context.Context, session *client.Session,
//	pp *pb.FlameProposal) (sm.Result, error) {
//	cmd, err := proto.Marshal(pp)
//	if err != nil {
//		return sm.Result{}, err
//	}
//	return n.mNodeHost.SyncPropose(ctx, session, cmd)
//}

//func (n *NodeHost) GetNoOPSession(clusterID uint64) *client.Session {
//	return n.mNodeHost.GetNoOPSession(clusterID)
//}
//
//func (n *NodeHost) SyncCloseSession(ctx context.Context, cs *client.Session) error {
//	return n.mNodeHost.SyncCloseSession(ctx, cs)
//}
//
//func (n *NodeHost) SyncRequestDeleteNode(ctx context.Context,
//	clusterID uint64, nodeID uint64, configChangeIndex uint64) error {
//	return n.mNodeHost.SyncRequestDeleteNode(ctx, clusterID, nodeID, configChangeIndex)
//}
//
//func (n *NodeHost) SyncRequestAddNode(ctx context.Context, clusterID uint64, nodeID uint64,
//	address string, configChangeIndex uint64) error {
//	return n.mNodeHost.SyncRequestAddNode(ctx, clusterID, nodeID, address, configChangeIndex)
//}
//
//func (n *NodeHost) SyncRequestAddObserver(ctx context.Context,
//	clusterID uint64, nodeID uint64,
//	address string, configChangeIndex uint64) error {
//	return n.mNodeHost.SyncRequestAddObserver(ctx, clusterID, nodeID, address, configChangeIndex)
//}
//
//func (n *NodeHost) SyncRequestAddWitness(ctx context.Context,
//	clusterID uint64, nodeID uint64,
//	address string, configChangeIndex uint64) error {
//	return n.mNodeHost.SyncRequestAddWitness(ctx, clusterID, nodeID, address, configChangeIndex)
//}
//
//func (n *NodeHost) SyncRequestSnapshot(ctx context.Context,
//	clusterID uint64,
//	opt dragonboat.SnapshotOption) (uint64, error) {
//	return n.mNodeHost.SyncRequestSnapshot(ctx, clusterID, opt)
//}
