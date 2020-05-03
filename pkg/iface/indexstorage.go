package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IIndexStorage interface {
	Open(path string, secretKey []byte, configuration interface{}) error
	SetIndexMeta(meta *pb.FlameIndexMeta) error
	Index(data []*variant.IndexData) error
	Close() error

	CanIndex(namespace string) bool
}
