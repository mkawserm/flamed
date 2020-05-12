package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IIndexStorage interface {
	Open(path string, secretKey []byte, configuration interface{}) error
	Close() error

	UpsertIndexMeta(meta *pb.IndexMeta) error
	DeleteIndexMeta(meta *pb.IndexMeta) error
	DefaultIndexMeta(namespace string) error

	ApplyIndex(namespace string, data []*variant.IndexData) error
	CanIndex(namespace string) bool
}
