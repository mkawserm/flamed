package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IStateIterator interface {
	Next()
	Close()
	Valid() bool
	Rewind()
	Seek(key []byte)
	StateSnapshot() *pb.StateSnapshot
	ValidForPrefix(prefix []byte) bool
}

type IStateContext interface {
	GetForwardIterator() IStateIterator
	GetReverseIterator() IStateIterator
	GetKeyOnlyForwardIterator() IStateIterator
	GetKeyOnlyReverseIterator() IStateIterator

	GetState(key []byte) (*pb.StateEntry, error)
	SetState(key []byte, entry *pb.StateEntry) error
	DeleteState(key []byte) error

	SetIndex(id string, data interface{}) error
	DeleteIndex(id string) error

	AutoIndexMeta() bool
	CanIndex(namespace string) bool
	SetIndexMeta(meta *pb.IndexMeta) error
	DeleteIndexMeta(meta *pb.IndexMeta) error
	DefaultIndexMeta(namespace string) error
	ApplyIndex(namespace string, data []*variant.IndexData) error
}
