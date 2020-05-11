package iface

import "github.com/mkawserm/flamed/pkg/pb"

type IStateContext interface {
	GetState(key []byte) ([]byte, error)
	SetState(key []byte, value []byte) error
	DeleteState(key []byte) error

	SetIndex(id string, data interface{}) error
	DeleteIndex(id string) error

	CanIndex(namespace string) bool
	SetIndexMeta(meta *pb.IndexMeta) error
	DeleteIndexMeta(meta *pb.IndexMeta) error
	AutoIndexMeta() bool
}
