package conf

import (
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/raftio"
	"time"
)

type NodeHostConfigurationInput struct {
	NodeID      uint64 `json:"nodeId"`
	NodeHostDir string `json:"nodeHostDir"`
	WALDir      string `json:"walDir"`

	CheckQuorum  bool   `json:"checkQuorum"`
	ElectionRTT  uint64 `json:"electionRTT"`
	HeartbeatRTT uint64 `json:"heartbeatRTT"`

	SnapshotEntries     uint64 `json:"snapshotEntries"`
	CompactionOverhead  uint64 `json:"compactionOverhead"`
	OrderedConfigChange bool   `json:"orderedConfigChange"`
	MaxInMemLogSize     uint64 `json:"maxInMemLogSize"`

	DisableAutoCompactions bool `json:"disableAutoCompactions"`

	IsObserver bool `json:"isObserver"`
	IsWitness  bool `json:"isWitness"`
	Quiesce    bool `json:"quiesce"`

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

	LogDBConfig         config.LogDBConfig          `json:"-"`
	LogDBFactory        config.LogDBFactoryFunc     `json:"-"`
	RaftRPCFactory      config.RaftRPCFactoryFunc   `json:"-"`
	RaftEventListener   raftio.IRaftEventListener   `json:"-"`
	SystemEventListener raftio.ISystemEventListener `json:"-"`
}

type NodeHostConfiguration struct {
	NodeHostConfigurationInput NodeHostConfigurationInput
}

func (n *NodeHostConfiguration) NodeID() uint64 {
	return n.NodeHostConfigurationInput.NodeID
}

func (n *NodeHostConfiguration) NodeHostDir() string {
	return n.NodeHostConfigurationInput.NodeHostDir
}

func (n *NodeHostConfiguration) WALDir() string {
	return n.NodeHostConfigurationInput.WALDir
}

func (n *NodeHostConfiguration) CheckQuorum() bool {
	return n.NodeHostConfigurationInput.CheckQuorum
}

func (n *NodeHostConfiguration) ElectionRTT() uint64 {
	return n.NodeHostConfigurationInput.ElectionRTT
}

func (n *NodeHostConfiguration) HeartbeatRTT() uint64 {
	return n.NodeHostConfigurationInput.HeartbeatRTT
}

func (n *NodeHostConfiguration) SnapshotEntries() uint64 {
	return n.NodeHostConfigurationInput.SnapshotEntries
}

func (n *NodeHostConfiguration) CompactionOverhead() uint64 {
	return n.NodeHostConfigurationInput.CompactionOverhead
}

func (n *NodeHostConfiguration) OrderedConfigChange() bool {
	return n.NodeHostConfigurationInput.OrderedConfigChange
}

func (n *NodeHostConfiguration) MaxInMemLogSize() uint64 {
	return n.NodeHostConfigurationInput.MaxInMemLogSize
}

func (n *NodeHostConfiguration) DisableAutoCompactions() bool {
	return n.NodeHostConfigurationInput.DisableAutoCompactions
}

func (n *NodeHostConfiguration) IsObserver() bool {
	return n.NodeHostConfigurationInput.IsObserver
}

func (n *NodeHostConfiguration) IsWitness() bool {
	return n.NodeHostConfigurationInput.IsWitness
}

func (n *NodeHostConfiguration) Quiesce() bool {
	return n.NodeHostConfigurationInput.Quiesce
}

/*Dragonboat NodeHostConfig*/
func (n *NodeHostConfiguration) DeploymentID() uint64 {
	return n.NodeHostConfigurationInput.DeploymentID
}

//WALDir() string
//NodeHostDir() string
func (n *NodeHostConfiguration) RTTMillisecond() uint64 {
	return n.NodeHostConfigurationInput.RTTMillisecond
}

func (n *NodeHostConfiguration) RaftAddress() string {
	return n.NodeHostConfigurationInput.RaftAddress
}

func (n *NodeHostConfiguration) ListenAddress() string {
	return n.NodeHostConfigurationInput.ListenAddress
}

func (n *NodeHostConfiguration) MutualTLS() bool {
	return n.NodeHostConfigurationInput.MutualTLS
}

func (n *NodeHostConfiguration) CAFile() string {
	return n.NodeHostConfigurationInput.CAFile
}

func (n *NodeHostConfiguration) CertFile() string {
	return n.NodeHostConfigurationInput.CertFile
}

func (n *NodeHostConfiguration) KeyFile() string {
	return n.NodeHostConfigurationInput.KeyFile
}

func (n *NodeHostConfiguration) MaxSendQueueSize() uint64 {
	return n.NodeHostConfigurationInput.MaxSendQueueSize
}

func (n *NodeHostConfiguration) MaxReceiveQueueSize() uint64 {
	return n.NodeHostConfigurationInput.MaxReceiveQueueSize
}

func (n *NodeHostConfiguration) LogDBFactory() config.LogDBFactoryFunc {
	return n.NodeHostConfigurationInput.LogDBFactory
}

func (n *NodeHostConfiguration) RaftRPCFactory() config.RaftRPCFactoryFunc {
	return n.NodeHostConfigurationInput.RaftRPCFactory
}

func (n *NodeHostConfiguration) EnableMetrics() bool {
	return n.NodeHostConfigurationInput.EnableMetrics
}

func (n *NodeHostConfiguration) MaxSnapshotSendBytesPerSecond() uint64 {
	return n.NodeHostConfigurationInput.MaxSnapshotSendBytesPerSecond
}

func (n *NodeHostConfiguration) MaxSnapshotRecvBytesPerSecond() uint64 {
	return n.NodeHostConfigurationInput.MaxSnapshotRecvBytesPerSecond
}

func (n *NodeHostConfiguration) RaftEventListener() raftio.IRaftEventListener {
	return n.NodeHostConfigurationInput.RaftEventListener
}

func (n *NodeHostConfiguration) SystemEventListener() raftio.ISystemEventListener {
	return n.NodeHostConfigurationInput.SystemEventListener
}

func (n *NodeHostConfiguration) SystemTickerPrecision() time.Duration {
	return n.NodeHostConfigurationInput.SystemTickerPrecision
}

func (n *NodeHostConfiguration) NotifyCommit() bool {
	return n.NodeHostConfigurationInput.NotifyCommit
}

func (n *NodeHostConfiguration) LogDBConfig() config.LogDBConfig {
	return n.NodeHostConfigurationInput.LogDBConfig
}
