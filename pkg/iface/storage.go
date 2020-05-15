package iface

import (
	"github.com/mkawserm/flamed/pkg/variant"
	"io"
	"time"
)

type IStorageConfiguration interface {
	/*Storage Config*/
	CacheSize() int
	BatchSize() int

	IndexEnable() bool
	AutoIndexMeta() bool

	StorageTaskQueue() variant.TaskQueue

	StateStoragePath() string
	StateStorageSecretKey() []byte

	IndexStoragePath() string
	IndexStorageSecretKey() []byte

	StoragePluginState() IStateStorage
	StoragePluginIndex() IIndexStorage

	StateStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}

	AddTransactionProcessor(tp ITransactionProcessor)
	IsTransactionProcessorExists(family, version string) bool
	GetTransactionProcessor(family, version string) ITransactionProcessor
}

type IStorage interface {
	RunGC()
	Open() error
	Close() error

	SetConfiguration(configuration IStorageConfiguration) bool

	ChangeSecretKey(path string,
		oldSecretKey []byte,
		newSecretKey []byte,
		encryptionKeyRotationDuration time.Duration) error

	PrepareSnapshot() (interface{}, error)
	RecoverFromSnapshot(r io.Reader) error
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error

	SaveAppliedIndex(u uint64) error
	QueryAppliedIndex() (uint64, error)
}
