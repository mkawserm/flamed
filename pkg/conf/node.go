package conf

import (
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/plugin/pebble"
)

type NodeConfigurationInput struct {
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

	LogDBFactory   config.LogDBFactoryFunc   `json:"-"`
	RaftRPCFactory config.RaftRPCFactoryFunc `json:"-"`

	EnableMetrics                 bool   `json:"enableMetrics"`
	MaxSnapshotSendBytesPerSecond uint64 `json:"maxSnapshotSendBytesPerSecond"`
	MaxSnapshotRecvBytesPerSecond uint64 `json:"maxSnapshotRecvBytesPerSecond"`
}

type NodeConfiguration struct {
	NodeConfigurationInput NodeConfigurationInput
}

func (n *NodeConfiguration) NodeID() uint64 {
	return n.NodeConfigurationInput.NodeID
}

func (n *NodeConfiguration) NodeHostDir() string {
	return n.NodeConfigurationInput.NodeHostDir
}

func (n *NodeConfiguration) WALDir() string {
	return n.NodeConfigurationInput.WALDir
}

func (n *NodeConfiguration) CheckQuorum() bool {
	return n.NodeConfigurationInput.CheckQuorum
}

func (n *NodeConfiguration) ElectionRTT() uint64 {
	return n.NodeConfigurationInput.ElectionRTT
}

func (n *NodeConfiguration) HeartbeatRTT() uint64 {
	return n.NodeConfigurationInput.HeartbeatRTT
}

func (n *NodeConfiguration) SnapshotEntries() uint64 {
	return n.NodeConfigurationInput.SnapshotEntries
}

func (n *NodeConfiguration) CompactionOverhead() uint64 {
	return n.NodeConfigurationInput.CompactionOverhead
}

func (n *NodeConfiguration) OrderedConfigChange() bool {
	return n.NodeConfigurationInput.OrderedConfigChange
}

func (n *NodeConfiguration) MaxInMemLogSize() uint64 {
	return n.NodeConfigurationInput.MaxInMemLogSize
}

func (n *NodeConfiguration) DisableAutoCompactions() bool {
	return n.NodeConfigurationInput.DisableAutoCompactions
}

func (n *NodeConfiguration) IsObserver() bool {
	return n.NodeConfigurationInput.IsObserver
}

func (n *NodeConfiguration) IsWitness() bool {
	return n.NodeConfigurationInput.IsWitness
}

func (n *NodeConfiguration) Quiesce() bool {
	return n.NodeConfigurationInput.Quiesce
}

/*Dragonboat NodeHostConfig*/
func (n *NodeConfiguration) DeploymentID() uint64 {
	return n.NodeConfigurationInput.DeploymentID
}

//WALDir() string
//NodeHostDir() string
func (n *NodeConfiguration) RTTMillisecond() uint64 {
	return n.NodeConfigurationInput.RTTMillisecond
}

func (n *NodeConfiguration) RaftAddress() string {
	return n.NodeConfigurationInput.RaftAddress
}

func (n *NodeConfiguration) ListenAddress() string {
	return n.NodeConfigurationInput.ListenAddress
}

func (n *NodeConfiguration) MutualTLS() bool {
	return n.NodeConfigurationInput.MutualTLS
}

func (n *NodeConfiguration) CAFile() string {
	return n.NodeConfigurationInput.CAFile
}

func (n *NodeConfiguration) CertFile() string {
	return n.NodeConfigurationInput.CertFile
}

func (n *NodeConfiguration) KeyFile() string {
	return n.NodeConfigurationInput.KeyFile
}

func (n *NodeConfiguration) MaxSendQueueSize() uint64 {
	return n.NodeConfigurationInput.MaxSendQueueSize
}

func (n *NodeConfiguration) MaxReceiveQueueSize() uint64 {
	return n.NodeConfigurationInput.MaxReceiveQueueSize
}

func (n *NodeConfiguration) LogDBFactory() config.LogDBFactoryFunc {
	if n.NodeConfigurationInput.LogDBFactory == nil {
		return pebble.NewLogDB
	} else {
		return n.NodeConfigurationInput.LogDBFactory
	}
}

func (n *NodeConfiguration) RaftRPCFactory() config.RaftRPCFactoryFunc {
	return n.NodeConfigurationInput.RaftRPCFactory
}

func (n *NodeConfiguration) EnableMetrics() bool {
	return n.NodeConfigurationInput.EnableMetrics
}

func (n *NodeConfiguration) MaxSnapshotSendBytesPerSecond() uint64 {
	return n.NodeConfigurationInput.MaxSnapshotSendBytesPerSecond
}

func (n *NodeConfiguration) MaxSnapshotRecvBytesPerSecond() uint64 {
	return n.NodeConfigurationInput.MaxSnapshotRecvBytesPerSecond
}
