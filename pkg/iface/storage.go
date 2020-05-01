package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"io"
)

type IndexObjectType int

const (
	JSONMapType IndexObjectType = iota
	BleveClassifierType
	GolangStructType
)

type IStorageConfiguration interface {
	/*Storage Config*/
	StoragePath() string
	StorageSecretKey() []byte

	StoragePluginKV() IKVStorage
	StoragePluginIndex() IIndexStorage

	KVStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}

	IndexObject(namespace []byte, fields []string, value []byte) (IndexObjectType, interface{})
}

type IStorage interface {
	SetConfiguration(configuration IStorageConfiguration) bool
	Open() error
	ReadOnlyOpen() error
	Close() error
	RunGC()
	ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) error
	IsExists(namespace []byte, key []byte) bool
	Read(namespace []byte, key []byte) ([]byte, error)
	Delete(namespace []byte, key []byte) error
	Create(namespace []byte, key []byte, value []byte) error
	Update(namespace []byte, key []byte, value []byte) error
	Append(namespace []byte, key []byte, value []byte) error
	ApplyBatchAction(batch *pb.FlameBatchAction) error
	ApplyAction(action *pb.FlameAction) error

	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error
	RecoverFromSnapshot(r io.Reader) error
}
