package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IIndexStorage interface {
	Open(path string, secretKey []byte, configuration interface{}) error
	Close() error

	SetIndexMeta(meta *pb.FlameIndexMeta) error
	DeleteIndexMeta(meta *pb.FlameIndexMeta) error

	ApplyIndex(namespace string, data []*variant.IndexData) error
	CanIndex(namespace string) bool
}
