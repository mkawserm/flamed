package iface

type IndexObjectType int

const (
	JSONMapType IndexObjectType = iota
	BleveClassifierType
	GolangStructType
)

type IFlameConfiguration interface {
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
