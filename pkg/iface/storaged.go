package iface

type IStoragedConfiguration interface {
	StoragePluginRaftLog() IRaftLogStorage
	RaftLogStorageCustomConfiguration() interface{}
}

type IStoraged interface {
}
