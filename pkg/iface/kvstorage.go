package iface

import "github.com/mkawserm/flamed/pkg/pb"

type IKVStorage interface {
	Open(path string, secretKey []byte, configuration interface{}) (bool, error)
	Close() error
	RunGC()

	Read(namespace []byte, key []byte) ([]byte, error)
	Delete(namespace []byte, key []byte) (bool, error)
	Create(namespace []byte, key []byte, value []byte) (bool, error)
	Update(namespace []byte, key []byte, value []byte) (bool, error)

	ApplyBatch(batch *pb.FlameBatch) (bool, error)
	ApplyAction(batch *pb.FlameAction) (bool, error)

	SetSnapshotConfiguration(configuration interface{})
	AsyncSnapshot(snapshot chan *pb.FlameSnapshot, maxItem int) error
	SyncSnapshot() (*pb.FlameSnapshot, error)
	ApplySyncSnapshot(snapshot *pb.FlameSnapshot) (bool, error)
}