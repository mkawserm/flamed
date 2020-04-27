package iface

type IndexObjectType int

const (
	JSONMapType IndexObjectType = iota
	BleveClassifierType
	GolangStructType
)

type IStorageConfiguration interface {
	/*Storage Config*/
	StoragePath() string
	StorageSecretKey() []byte

	StoragePluginKV() IKVStorage
	StoragePluginIndex() IIndexStorage
	StoragePluginRaftLog() IRaftLogStorage

	KVStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}
	RaftLogStorageCustomConfiguration() interface{}

	IndexObject(namespace []byte, fields []string, value []byte) (IndexObjectType, interface{})
}

type IStorage interface {
}
