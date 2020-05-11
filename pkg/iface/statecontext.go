package iface

import "github.com/mkawserm/flamed/pkg/pb"

type IStateContext interface {
	GetState(key []byte) ([]byte, error)
	SetState(key []byte, value []byte) error
	DeleteState(key []byte) error

	SetIndex(id string, data interface{}) error
	DeleteIndex(id string) error

	SetIndexMeta(meta *pb.FlameIndexMeta) error
	DeleteIndexMeta(meta *pb.FlameIndexMeta) error
}
