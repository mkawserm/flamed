package iface

type IConfiguration interface {
	FlamePath() string
	StoragePluginKV() string
	StoragePluginIndex() string
	StoragePluginRaftLog() string

	KVStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}
	RaftLogStorageCustomConfiguration() interface{}
}
