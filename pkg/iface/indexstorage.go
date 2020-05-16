package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IIndexStorage interface {
	RunGC()
	Close() error
	Open(path string, secretKey []byte, configuration interface{}) error

	DefaultIndexMeta(namespace string) error
	UpsertIndexMeta(meta *pb.IndexMeta) error
	DeleteIndexMeta(meta *pb.IndexMeta) error

	CanIndex(namespace string) bool
	ApplyIndex(namespace string, data []*variant.IndexData) error
}
