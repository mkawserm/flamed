package iface

type IConfiguration interface {
	StoragePluginKV() string
	StoragePluginIndex() string
	StoragePluginRaftLog() string
}
