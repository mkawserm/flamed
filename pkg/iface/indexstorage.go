package iface

import (
	"context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IIndexStorageSearchResult interface {
	RawResult() interface{}
	ToMap() map[string]interface{}
	ToBytes() []byte
	ToSearchResponse() ISearchResponse
}

type IIndexStorage interface {
	RunGC()
	Close() error
	Open(path string, secretKey []byte, configuration interface{}) error

	DefaultIndexMeta(namespace string) error
	UpsertIndexMeta(meta *pb.IndexMeta) error
	DeleteIndexMeta(meta *pb.IndexMeta) error

	CustomIndexRule(namespace string, indexRule interface{}) error

	CanIndex(namespace string) bool
	ApplyIndex(namespace string, data []*variant.IndexData) error

	GlobalSearch(ctx context.Context, input *pb.GlobalSearchInput) (IIndexStorageSearchResult, error)
}
