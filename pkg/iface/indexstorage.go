package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IIndexStorage interface {
	Open(path string, secretKey []byte, configuration interface{}) error
	Close() error

	CreateIndexMeta(meta *pb.FlameIndexMeta) error
	UpdateIndexMeta(meta *pb.FlameIndexMeta) error
	DeleteIndexMeta(meta *pb.FlameIndexMeta) error

	CreateIndex(namespace string, data []*variant.IndexData) error
	UpdateIndex(namespace string, data []*variant.IndexData) error
	DeleteIndex(namespace string, data []*variant.IndexData) error

	CanIndex(namespace string) bool
}
