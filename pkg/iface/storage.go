package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"io"
)

type IStorageConfiguration interface {
	/*Storage Config*/
	CacheSize() int
	BatchSize() int

	StoragePath() string
	StorageSecretKey() []byte

	StoragePluginKV() IKVStorage
	StoragePluginIndex() IIndexStorage

	KVStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}

	/*IndexObject*/
	IndexObject
}

type IStorage interface {
	SetConfiguration(configuration IStorageConfiguration) bool
	Open() error
	Close() error
	RunGC()

	ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) error

	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error
	RecoverFromSnapshot(r io.Reader) error

	SaveAppliedIndex(u uint64) error
	QueryAppliedIndex() (uint64, error)

	Lookup(input interface{}, checkValidity bool) (interface{}, error)
	ApplyProposal(pp *pb.FlameProposal, checkValidity bool) error
}
