package iface

type IStoragedConfiguration interface {
	IStorageConfiguration
	//StoragePluginRaftLog() IRaftLogStorage
	//RaftLogStorageCustomConfiguration() interface{}
}

type IStoraged interface {
}
