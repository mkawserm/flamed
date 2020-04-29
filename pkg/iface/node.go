package iface

import "github.com/lni/dragonboat/v3/config"

type INodeConfiguration interface {
	NodeID() uint64
	//NodePath() string

	CheckQuorum() bool
	ElectionRTT() uint64
	HeartbeatRTT() uint64
	SnapshotEntries() uint64
	CompactionOverhead() uint64
	OrderedConfigChange() bool
	MaxInMemLogSize() uint64
	//SnapshotCompressionType() string
	//EntryCompressionType() string
	DisableAutoCompactions() bool

	IsObserver() bool
	IsWitness() bool
	Quiesce() bool

	/*Dragonboat NodeHostConfig*/
	DeploymentID() uint64
	WALDir() string
	NodeHostDir() string
	RTTMillisecond() uint64
	RaftAddress() string
	ListenAddress() string
	MutualTLS() bool
	CAFile() string
	CertFile() string
	KeyFile() string
	MaxSendQueueSize() uint64
	MaxReceiveQueueSize() uint64

	LogDBFactory() config.LogDBFactoryFunc
	RaftRPCFactory() config.RaftRPCFactoryFunc

	EnableMetrics() bool
	//RaftEventListener
	MaxSnapshotSendBytesPerSecond() uint64
	MaxSnapshotRecvBytesPerSecond() uint64
	//FS
	//SystemEventListener
	//SystemTickerPrecision
}
