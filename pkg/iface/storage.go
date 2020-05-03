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

	ApplyAction(action *pb.FlameAction) error
	ApplyBatchAction(batch *pb.FlameBatchAction) error

	ReadBatch(batch *pb.FlameBatchRead) error

	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error
	RecoverFromSnapshot(r io.Reader) error

	SaveAppliedIndex(u uint64) error
	QueryAppliedIndex() (uint64, error)

	CreateIndexMeta(meta *pb.FlameIndexMeta) error
	GetIndexMeta(meta *pb.FlameIndexMeta) error
	GetAllIndexMeta() ([]*pb.FlameIndexMeta, error)
	UpdateIndexMeta(meta *pb.FlameIndexMeta) error
	DeleteIndexMeta(meta *pb.FlameIndexMeta) error

	CreateUser(user *pb.FlameUser) error
	GetUser(user *pb.FlameUser) error
	GetAllUser() ([]*pb.FlameUser, error)
	UpdateUser(user *pb.FlameUser) error
	DeleteUser(user *pb.FlameUser) error

	CreateAccessControl(ac *pb.FlameAccessControl) error
	GetAccessControl(ac *pb.FlameAccessControl) error
	GetAllAccessControl() ([]*pb.FlameAccessControl, error)
	UpdateAccessControl(ac *pb.FlameAccessControl) error
	DeleteAccessControl(ac *pb.FlameAccessControl) error

	Lookup(input interface{}, checkValidity bool) (interface{}, error)
	ApplyProposal(pp *pb.FlameProposal, checkValidity bool) error
}
