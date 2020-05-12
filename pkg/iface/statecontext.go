package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IStateContext interface {
	GetState(key []byte) ([]byte, error)
	SetState(key []byte, value []byte) error
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
