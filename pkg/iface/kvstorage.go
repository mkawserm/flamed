package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"io"
)

type IKVStorage interface {
	Open(path string, secretKey []byte, readOnly bool, configuration interface{}) error
	Close() error

	RunGC()

	ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) error

	IsExists(namespace []byte, key []byte) bool

	Read(namespace []byte, key []byte) ([]byte, error)
	Iterate(seek, prefix []byte, limit int, receiver func(entry *pb.FlameEntry) bool) error
	IterateKeyOnly(seek, prefix []byte, limit int, receiver func(entry *pb.FlameEntry) bool) error

	Delete(namespace []byte, key []byte) error
	Create(namespace []byte, key []byte, value []byte) error
	Update(namespace []byte, key []byte, value []byte) error
	Append(namespace []byte, key []byte, value []byte) error

	ReadBatch(batch *pb.FlameBatchRead) error
	ApplyBatchAction(batch *pb.FlameBatchAction) error
	ApplyAction(action *pb.FlameAction) error

	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error
	RecoverFromSnapshot(r io.Reader) error

	CreateIndexMeta(meta *pb.FlameIndexMeta) error
	IsIndexMetaExists(meta *pb.FlameIndexMeta) bool
	GetIndexMeta(meta *pb.FlameIndexMeta) error
	GetAllIndexMeta() ([]*pb.FlameIndexMeta, error)
	UpdateIndexMeta(meta *pb.FlameIndexMeta) error
	DeleteIndexMeta(meta *pb.FlameIndexMeta) error

	CreateUser(user *pb.FlameUser) error
	IsUserExists(user *pb.FlameUser) bool
	GetUser(user *pb.FlameUser) error
	GetAllUser() ([]*pb.FlameUser, error)
	UpdateUser(user *pb.FlameUser) error
	DeleteUser(user *pb.FlameUser) error

	CreateAccessControl(ac *pb.FlameAccessControl) error
	IsAccessControlExists(ac *pb.FlameAccessControl) bool
	GetAccessControl(ac *pb.FlameAccessControl) error
	GetAllAccessControl() ([]*pb.FlameAccessControl, error)
	UpdateAccessControl(ac *pb.FlameAccessControl) error
	DeleteAccessControl(ac *pb.FlameAccessControl) error
}
