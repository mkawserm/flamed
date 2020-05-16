package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
)

type IStateIterator interface {
	Next()
	Close()
	Valid() bool
	Rewind()
	Seek(address []byte)
	StateSnapshot() *pb.StateSnapshot
	ValidForPrefix(prefix []byte) bool
}

type IStateContext interface {
	GetForwardIterator() IStateIterator
	GetReverseIterator() IStateIterator
	GetKeyOnlyForwardIterator() IStateIterator
	GetKeyOnlyReverseIterator() IStateIterator

	GetState(address []byte) (*pb.StateEntry, error)
	UpsertState(address []byte, entry *pb.StateEntry) error
	DeleteState(address []byte) error

	UpsertIndex(id string, data interface{}) error
	DeleteIndex(id string) error

	//AutoIndexMeta() bool
	//CanIndex(namespace string) bool
	UpsertIndexMeta(meta *pb.IndexMeta) error
	DeleteIndexMeta(meta *pb.IndexMeta) error
	DefaultIndexMeta(namespace string) error
	//ApplyIndex(namespace string, data []*variant.IndexData) error
}
