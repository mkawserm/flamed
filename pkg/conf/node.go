package conf

import "github.com/lni/dragonboat/v3/config"

type NodeConfigurationInput struct {
}

type NodeConfiguration struct {
}

func (n *NodeConfiguration) NodeID() uint64 {
	return 0
}

func (n *NodeConfiguration) NodePath() string {
	return ""
}

func (n *NodeConfiguration) CheckQuorum() bool {
	return false
}

func (n *NodeConfiguration) ElectionRTT() uint64 {
	return 0
}

func (n *NodeConfiguration) HeartbeatRTT() uint64 {
	return 0
}

func (n *NodeConfiguration) SnapshotEntries() uint64 {
	return 0
}

func (n *NodeConfiguration) CompactionOverhead() uint64 {
	return 0
}

func (n *NodeConfiguration) OrderedConfigChange() bool {
	return false
}

func (n *NodeConfiguration) MaxInMemLogSize() uint64 {
	return 0
}

func (n *NodeConfiguration) DisableAutoCompactions() bool {
	return false
}

func (n *NodeConfiguration) IsObserver() bool {
	return false
}

func (n *NodeConfiguration) IsWitness() bool {
	return false
}

func (n *NodeConfiguration) Quiesce() bool {
	return false
}

/*Dragonboat NodeHostConfig*/
func (n *NodeConfiguration) DeploymentID() uint64 {
	return 0
}

//WALDir() string
//NodeHostDir() string
func (n *NodeConfiguration) RTTMillisecond() string {
	return ""
}

func (n *NodeConfiguration) RaftAddress() string {
	return ""
}

func (n *NodeConfiguration) ListenAddress() string {
	return ""
}

func (n *NodeConfiguration) MutualTLS() bool {
	return false
}

func (n *NodeConfiguration) CAFile() string {
	return ""
}

func (n *NodeConfiguration) CertFile() string {
	return ""
}

func (n *NodeConfiguration) KeyFile() string {
	return ""
}

func (n *NodeConfiguration) MaxSendQueueSize() uint64 {
	return 0
}

func (n *NodeConfiguration) MaxReceiveQueueSize() uint64 {
	return 0
}

func (n *NodeConfiguration) LogDBFactory() config.LogDBFactoryFunc {
	return nil
}

func (n *NodeConfiguration) RaftRPCFactory() config.RaftRPCFactoryFunc {
	return nil
}

func (n *NodeConfiguration) EnableMetrics() bool {
	return false
}

func (n *NodeConfiguration) MaxSnapshotSendBytesPerSecond() uint64 {
	return 0
}

func (n *NodeConfiguration) MaxSnapshotRecvBytesPerSecond() uint64 {
	return 0
}
