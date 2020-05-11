package iface

import (
	"io"
)

type IStorageConfiguration interface {
	/*Storage Config*/
	CacheSize() int
	BatchSize() int

	IndexEnable() bool
	AutoIndexMeta() bool

	StateStoragePath() string
	StateStorageSecretKey() []byte

	IndexStoragePath() string
	IndexStorageSecretKey() []byte

	StoragePluginIndex() IIndexStorage
	StoragePluginState() IStateStorage

	StateStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}

	GetTransactionProcessor(family, version string) ITransactionProcessor
}

type IStorage interface {
	SetConfiguration(configuration IStorageConfiguration) bool
	Open() error
	Close() error
	RunGC()

	ChangeSecretKey(path string, oldSecretKey []byte, newSecretKey []byte) error

	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error
	RecoverFromSnapshot(r io.Reader) error

	SaveAppliedIndex(u uint64) error
	QueryAppliedIndex() (uint64, error)

	Lookup(input interface{}) (interface{}, error)
}
