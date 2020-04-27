package iface

type IndexObjectType int

const (
	JSONMapType IndexObjectType = iota
	BleveClassifierType
	GolangStructType
)

type IStorageConfiguration interface {
	/*Flame config*/
	StoragePath() string
	StorageSecretKey() []byte

	StoragePluginKV() IKVStorage
	StoragePluginIndex() IIndexStorage
	StoragePluginRaftLog() IRaftLogStorage

	KVStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}
	RaftLogStorageCustomConfiguration() interface{}

	KVStorageSnapshotConfiguration() interface{}
	//IndexStorageSnapshotConfiguration() interface{}
	//RaftLogStorageSnapshotConfiguration() interface{}

	IndexObject(namespace []byte, fields []string, value []byte) (IndexObjectType, interface{})
}

type IStorage interface {
}
