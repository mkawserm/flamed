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
	ReadUsingPrefix(prefix []byte) ([]*pb.FlameEntry, error)

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
}
