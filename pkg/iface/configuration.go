package iface

type IndexObjectType int

const (
	JSONMapType IndexObjectType = iota
	BleveClassifierType
)

type IConfiguration interface {

	/*Dragonboat Raft config*/
	InitialMembers() map[uint64]string
	Join() bool
	NodeID() uint64
	ClusterID() uint64
	ClusterName() []byte

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
	//WALDir() string
	//NodeHostDir() string
	RTTMillisecond() string
	RaftAddress() string
	ListenAddress() string
	MutualTLS() bool
	CAFile() string
	CertFile() string
	KeyFile() string
	MaxSendQueueSize() uint64
	MaxReceiveQueueSize() uint64
	//LogDBFactory
	//RaftRPCFactory
	EnableMetrics() bool
	//RaftEventListener
	MaxSnapshotSendBytesPerSecond() uint64
	MaxSnapshotRecvBytesPerSecond() uint64
	//FS
	//SystemEventListener
	//SystemTickerPrecision

	/*Flame config*/
	FlamePath() string
	FlameSecretKey() []byte

	StoragePluginKV() IKVStorage
	StoragePluginIndex() IIndexStorage
	StoragePluginRaftLog() IRaftLogStorage

	KVStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}
	RaftLogStorageCustomConfiguration() interface{}

	KVStorageSnapshotConfiguration() interface{}
	//IndexStorageSnapshotConfiguration() interface{}
	//RaftLogStorageSnapshotConfiguration() interface{}

	IndexObject(namespace []byte, keys []string, value []byte) (IndexObjectType, interface{})
}
