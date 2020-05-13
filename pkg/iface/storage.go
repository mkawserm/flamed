package iface

import (
	"io"
	"time"
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

	StoragePluginState() IStateStorage
	StoragePluginIndex() IIndexStorage

	StateStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}

	IsTransactionProcessorExists(family, version string) bool
	AddTransactionProcessor(tp ITransactionProcessor)
	GetTransactionProcessor(family, version string) ITransactionProcessor
}

type IStorage interface {
	SetConfiguration(configuration IStorageConfiguration) bool
	Open() error
	Close() error
	RunGC()

	ChangeSecretKey(path string,
		oldSecretKey []byte,
		newSecretKey []byte,
		encryptionKeyRotationDuration time.Duration) error

	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error
	RecoverFromSnapshot(r io.Reader) error

	SaveAppliedIndex(u uint64) error
	QueryAppliedIndex() (uint64, error)
}
