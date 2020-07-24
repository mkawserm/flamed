package iface

import (
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/raftio"
	"time"
)

type INodeConfiguration interface {
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

	EnableMetrics() bool
	MaxSnapshotSendBytesPerSecond() uint64
	MaxSnapshotRecvBytesPerSecond() uint64

	LogDBFactory() config.LogDBFactoryFunc
	RaftRPCFactory() config.RaftRPCFactoryFunc

	NotifyCommit() bool
	LogDB() config.LogDBConfig

	SystemTickerPrecision() time.Duration
	RaftEventListener() raftio.IRaftEventListener
	SystemEventListener() raftio.ISystemEventListener
}
