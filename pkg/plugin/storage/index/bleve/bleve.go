package bleve

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type Bleve struct {
}

func (b *Bleve) Open(path string, secretKey []byte, configuration interface{}) error {
	return nil
}

func (b *Bleve) SetIndexMeta(meta *pb.FlameIndexMeta) error {
	return nil
}

func (b *Bleve) Index(data []*variant.IndexData) error {
	return nil
}

func (b *Bleve) Close() error {
	return nil
}

func (b *Bleve) CanIndex(namespace string) bool {
	return false
}
