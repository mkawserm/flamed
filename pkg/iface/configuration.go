package iface

type IConfiguration interface {
	FlamedPath() string
	StoragePluginKV() string
	StoragePluginIndex() string
	StoragePluginRaftLog() string

	KVStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}
	RaftLogStorageCustomConfiguration() interface{}
}
