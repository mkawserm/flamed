package iface

import (
	sm "github.com/lni/dragonboat/v3/statemachine"
	"io"
)

type IStoragedConfiguration interface {
	IStorageConfiguration
	//StoragePluginRaftLog() IRaftLogStorage
	//RaftLogStorageCustomConfiguration() interface{}
}

type IStoraged interface {
	SetConfiguration(configuration IStoragedConfiguration) bool
	Open(<-chan struct{}) (uint64, error)
	Sync() error
	Close() error
	Update(entries []sm.Entry) ([]sm.Entry, error)
	Lookup(input interface{}) (interface{}, error)
	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer, done <-chan struct{}) error
	RecoverFromSnapshot(r io.Reader, done <-chan struct{}) error
}
