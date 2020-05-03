package bleve

import (
	bleveSearch "github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/index/scorch"
	bleveMapping "github.com/blevesearch/bleve/mapping"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"os"
)

type Bleve struct {
	path          string
	secretKey     []byte
	configuration interface{}
}

func (b *Bleve) Open(path string, secretKey []byte, configuration interface{}) error {
	if len(path) == 0 {
		return x.ErrPathCanNotBeEmpty
	}

	b.path = path
	b.secretKey = secretKey
	b.configuration = configuration

	return nil
}

func (b *Bleve) CreateIndexMeta(meta *pb.FlameIndexMeta) error {
	p := b.path + "/" + string(meta.Namespace)
	if b.isPathExists(p) {
		b.removeAll(p)
	}

	_, err := bleveSearch.NewUsing(p,
		b.getMapping(meta),
		scorch.Name,
		scorch.Name,
		nil)

	if err != nil {
		internalLogger.Debug("error while adding index meta", zap.Error(err))
		return x.ErrFailedToAddIndexMeta
	}

	return nil
}

func (b *Bleve) UpdateIndexMeta(meta *pb.FlameIndexMeta) error {
	return b.CreateIndexMeta(meta)
}

func (b *Bleve) DeleteIndexMeta(meta *pb.FlameIndexMeta) error {
	p := b.path + "/" + string(meta.Namespace)
	if b.isPathExists(p) {
		b.removeAll(p)
	}

	return nil
}

func (b *Bleve) CreateIndex(namespace string, data []*variant.IndexData) error {
	p := b.path + "/" + namespace
	index, err := bleveSearch.OpenUsing(p, nil)

	if err != nil {
		internalLogger.Debug("index db opening error", zap.Error(err))
		return x.ErrFailedToCreateIndex
	}

	defer func() {
		_ = index.Close()
	}()

	batch := index.NewBatch()
	for idx := range data {
		err = batch.Index(data[idx].ID, data[idx].Data)
		if err != nil {
			internalLogger.Debug("indexing error", zap.Error(err))
			return x.ErrFailedToCreateIndex
		}
	}

	err = index.Batch(batch)
	if err != nil {
		internalLogger.Debug("batch processing error", zap.Error(err))
		return x.ErrFailedToCreateIndex
	}

	return nil
}

func (b *Bleve) UpdateIndex(namespace string, data []*variant.IndexData) error {
	p := b.path + "/" + namespace
	index, err := bleveSearch.OpenUsing(p, nil)

	if err != nil {
		internalLogger.Debug("index db opening error", zap.Error(err))
		return x.ErrFailedToUpdateIndex
	}

	defer func() {
		_ = index.Close()
	}()

	batch := index.NewBatch()
	for idx := range data {
		err = batch.Index(data[idx].ID, data[idx].Data)
		if err != nil {
			internalLogger.Debug("indexing error", zap.Error(err))
			return x.ErrFailedToUpdateIndex
		}
	}

	err = index.Batch(batch)
	if err != nil {
		internalLogger.Debug("batch processing error", zap.Error(err))
		return x.ErrFailedToUpdateIndex
	}

	return nil
}

func (b *Bleve) DeleteIndex(namespace string, data []*variant.IndexData) error {
	p := b.path + "/" + namespace
	index, err := bleveSearch.OpenUsing(p, nil)

	if err != nil {
		internalLogger.Debug("index db opening error", zap.Error(err))
		return x.ErrFailedToDeleteIndex
	}

	defer func() {
		_ = index.Close()
	}()

	batch := index.NewBatch()
	for idx := range data {
		batch.Delete(data[idx].ID)
	}

	err = index.Batch(batch)
	if err != nil {
		internalLogger.Debug("batch processing error", zap.Error(err))
		return x.ErrFailedToDeleteIndex
	}

	return nil
}

func (b *Bleve) Close() error {
	return nil
}

func (b *Bleve) CanIndex(namespace string) bool {
	return b.isPathExists(b.path + "/" + namespace)
}

func (b *Bleve) isPathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (b *Bleve) removeAll(path string) {
	_ = os.RemoveAll(path)
}

func (b *Bleve) getMapping(*pb.FlameIndexMeta) *bleveMapping.IndexMappingImpl {
	indexMapping := bleveMapping.NewIndexMapping()
	return indexMapping
}
