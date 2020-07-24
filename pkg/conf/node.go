package conf

import (
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/raftio"
	"sync"
	"time"
)

type NodeConfigurationInput struct {
	NodeHostDir    string `json:"nodeHostDir"`
	WALDir         string `json:"walDir"`
	DeploymentID   uint64 `json:"deploymentId"`
	RTTMillisecond uint64 `json:"rttMillisecond"`

	RaftAddress   string `json:"raftAddress"`
	ListenAddress string `json:"listenAddress"`

	MutualTLS bool   `json:"mutualTLS"`
	CAFile    string `json:"caFile"`
	CertFile  string `json:"certFile"`
	KeyFile   string `json:"keyFile"`

	MaxSendQueueSize    uint64 `json:"maxSendQueueSize"`
	MaxReceiveQueueSize uint64 `json:"maxReceiveQueueSize"`

	EnableMetrics                 bool   `json:"enableMetrics"`
	MaxSnapshotSendBytesPerSecond uint64 `json:"maxSnapshotSendBytesPerSecond"`
	MaxSnapshotRecvBytesPerSecond uint64 `json:"maxSnapshotRecvBytesPerSecond"`

	SystemTickerPrecision time.Duration `json:"systemTickerPrecision"`

	NotifyCommit bool `json:"notifyCommit"`

	LogDB               config.LogDBConfig          `json:"-"`
	LogDBFactory        config.LogDBFactoryFunc     `json:"-"`
	RaftRPCFactory      config.RaftRPCFactoryFunc   `json:"-"`
	RaftEventListener   raftio.IRaftEventListener   `json:"-"`
	SystemEventListener raftio.ISystemEventListener `json:"-"`
}

type NodeConfiguration struct {
	mMutex                 sync.Mutex
	NodeConfigurationInput NodeConfigurationInput
}

/*Dragonboat NodeHostConfig*/
func (n *NodeConfiguration) NodeHostDir() string {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.NodeHostDir
}

func (n *NodeConfiguration) WALDir() string {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.WALDir
}

func (n *NodeConfiguration) DeploymentID() uint64 {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.DeploymentID
}

//WALDir() string
//NodeHostDir() string
func (n *NodeConfiguration) RTTMillisecond() uint64 {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.RTTMillisecond
}

func (n *NodeConfiguration) RaftAddress() string {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.RaftAddress
}

func (n *NodeConfiguration) ListenAddress() string {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.ListenAddress
}

func (n *NodeConfiguration) MutualTLS() bool {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.MutualTLS
}

func (n *NodeConfiguration) CAFile() string {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.CAFile
}

func (n *NodeConfiguration) CertFile() string {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.CertFile
}

func (n *NodeConfiguration) KeyFile() string {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.KeyFile
}

func (n *NodeConfiguration) MaxSendQueueSize() uint64 {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.MaxSendQueueSize
}

func (n *NodeConfiguration) MaxReceiveQueueSize() uint64 {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.MaxReceiveQueueSize
}

func (n *NodeConfiguration) LogDBFactory() config.LogDBFactoryFunc {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.LogDBFactory
}

func (n *NodeConfiguration) RaftRPCFactory() config.RaftRPCFactoryFunc {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.RaftRPCFactory
}

func (n *NodeConfiguration) EnableMetrics() bool {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.EnableMetrics
}

func (n *NodeConfiguration) MaxSnapshotSendBytesPerSecond() uint64 {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.MaxSnapshotSendBytesPerSecond
}

func (n *NodeConfiguration) MaxSnapshotRecvBytesPerSecond() uint64 {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.MaxSnapshotRecvBytesPerSecond
}

func (n *NodeConfiguration) RaftEventListener() raftio.IRaftEventListener {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.RaftEventListener
}

func (n *NodeConfiguration) SystemEventListener() raftio.ISystemEventListener {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.SystemEventListener
}

func (n *NodeConfiguration) SystemTickerPrecision() time.Duration {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.SystemTickerPrecision
}

func (n *NodeConfiguration) NotifyCommit() bool {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.NotifyCommit
}

func (n *NodeConfiguration) LogDB() config.LogDBConfig {
	n.mMutex.Lock()
	defer n.mMutex.Unlock()
	return n.NodeConfigurationInput.LogDB
}
