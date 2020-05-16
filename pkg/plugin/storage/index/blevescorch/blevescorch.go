package blevescorch

import (
	bleveSearch "github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/index/scorch"
	bleveMapping "github.com/blevesearch/bleve/mapping"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"os"
)

type BleveScorch struct {
	path          string
	secretKey     []byte
	configuration interface{}
}

func (b *BleveScorch) Open(path string, secretKey []byte, configuration interface{}) error {
	if len(path) == 0 {
		return x.ErrPathCanNotBeEmpty
	}

	b.path = path
	b.secretKey = secretKey
	b.configuration = configuration

	return nil
}

func (b *BleveScorch) UpsertIndexMeta(meta *pb.IndexMeta) error {
	if !meta.Enabled {
		return nil
	}

	p := b.path + "/" + string(meta.Namespace)
	if b.isPathExists(p) {
		b.removeAll(p)
	}

	index, err := bleveSearch.NewUsing(p,
		b.getMapping(meta),
		scorch.Name,
		scorch.Name,
		nil)

	if err != nil {
		internalLogger.Debug("error while adding indexmeta meta", zap.Error(err))
		return x.ErrFailedToCreateIndexMeta
	}

	_ = index.Close()

	return nil
}

func (b *BleveScorch) DeleteIndexMeta(meta *pb.IndexMeta) error {
	p := b.path + "/" + string(meta.Namespace)
	if b.isPathExists(p) {
		b.removeAll(p)
	}

	return nil
}

func (b *BleveScorch) DefaultIndexMeta(namespace string) error {
	p := b.path + "/" + namespace
	if b.isPathExists(p) {
		b.removeAll(p)
	}

	indexMapping := bleveMapping.NewIndexMapping()
	index, err := bleveSearch.NewUsing(p,
		indexMapping,
		scorch.Name,
		scorch.Name,
		nil)

	if err != nil {
		internalLogger.Debug("error while adding default indexmeta meta", zap.Error(err))
		return x.ErrFailedToCreateIndexMeta
	}

	_ = index.Close()

	return nil
}

func (b *BleveScorch) ApplyIndex(namespace string, data []*variant.IndexData) error {
	p := b.path + "/" + namespace
	if !b.isPathExists(p) {
		return nil
	}

	index, err := bleveSearch.OpenUsing(p, nil)

	if err != nil {
		internalLogger.Debug("indexmeta db opening error", zap.Error(err))
		return x.ErrFailedToApplyIndex
	}

	defer func() {
		_ = index.Close()
	}()

	batch := index.NewBatch()
	for idx := range data {
		if data[idx].Action == constant.UPSERT {
			err = batch.Index(data[idx].ID, data[idx].Data)
			if err != nil {
				internalLogger.Debug("indexing error", zap.Error(err))
				return x.ErrFailedToCreateIndex
			}
		}

		if data[idx].Action == constant.DELETE {
			batch.Delete(data[idx].ID)
		}
	}

	err = index.Batch(batch)
	if err != nil {
		internalLogger.Debug("batch processing error", zap.Error(err))
		return x.ErrFailedToApplyIndex
	}

	return nil
}

func (b *BleveScorch) Close() error {
	return nil
}

func (b *BleveScorch) RunGC() {

}

func (b *BleveScorch) CanIndex(namespace string) bool {
	return b.isPathExists(b.path + "/" + namespace)
}

func (b *BleveScorch) isPathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (b *BleveScorch) removeAll(path string) {
	_ = os.RemoveAll(path)
}

func (b *BleveScorch) getMapping(meta *pb.IndexMeta) *bleveMapping.IndexMappingImpl {
	indexMapping := bleveMapping.NewIndexMapping()
	if meta.Default {
		return indexMapping
	}

	return indexMapping
}
