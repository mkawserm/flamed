package iface

import "github.com/mkawserm/flamed/pkg/pb"

type IKVStorage interface {
	Open(path string, secretKey []byte, readOnly bool, configuration interface{}) (bool, error)
	Close() error

	RunGC()

	ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) (bool, error)

	IsExists(namespace []byte, key []byte) bool

	Read(namespace []byte, key []byte) ([]byte, error)
	ReadUsingPrefix(prefix []byte) ([]*pb.FlameEntry, error)

	Delete(namespace []byte, key []byte) (bool, error)
	Create(namespace []byte, key []byte, value []byte) (bool, error)
	Update(namespace []byte, key []byte, value []byte) (bool, error)

	ApplyBatch(batch *pb.FlameBatch) (bool, error)
	ApplyAction(action *pb.FlameAction) (bool, error)

	AsyncSnapshot(snapshot chan<- *pb.FlameSnapshot) error
	ApplyAsyncSnapshot(snapshot <-chan *pb.FlameSnapshot) (bool, error)

	SyncSnapshot() (*pb.FlameSnapshot, error)
	ApplySyncSnapshot(snapshot *pb.FlameSnapshot) (bool, error)
}
