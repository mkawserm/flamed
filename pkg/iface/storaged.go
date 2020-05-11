package iface

import (
	sm "github.com/lni/dragonboat/v3/statemachine"
	"io"
)

type IStoragedConfiguration interface {
	IStorageConfiguration
}

type IStoraged interface {
	Sync() error
	Close() error
	Open(<-chan struct{}) (uint64, error)

	Update(entries []sm.Entry) ([]sm.Entry, error)
	Lookup(input interface{}) (interface{}, error)

	SetConfiguration(configuration IStoragedConfiguration) bool

	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer, done <-chan struct{}) error
	RecoverFromSnapshot(r io.Reader, done <-chan struct{}) error
}
