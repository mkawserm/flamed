package blevescorch

import (
	"context"
	bleveSearch "github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/standard"
	"github.com/blevesearch/bleve/index/scorch"
	bleveMapping "github.com/blevesearch/bleve/mapping"
	bleveSearchQuery "github.com/blevesearch/bleve/search/query"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"os"
)

import _ "github.com/blevesearch/bleve/analysis/analyzer/keyword"
import _ "github.com/blevesearch/bleve/analysis/analyzer/simple"
import _ "github.com/blevesearch/bleve/analysis/analyzer/web"

type BleveScorch struct {
	path          string
	secretKey     []byte
	configuration interface{}
}

func (b *BleveScorch) IndexStorageName() string {
	return Name
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
		logger.L(Name).Debug("error while adding index meta", zap.Error(err))
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
		logger.L(Name).Debug("error while adding default index meta", zap.Error(err))
		return x.ErrFailedToCreateIndexMeta
	}

	_ = index.Close()

	return nil
}

func (b *BleveScorch) CustomIndexRule(namespace string, indexRule interface{}) error {
	indexMapping, ok := indexRule.(*bleveMapping.IndexMappingImpl)
	if !ok {
		return x.ErrUnknownValue
	}

	p := b.path + "/" + namespace
	if b.isPathExists(p) {
		b.removeAll(p)
	}

	index, err := bleveSearch.NewUsing(p,
		indexMapping,
		scorch.Name,
		scorch.Name,
		nil)

	if err != nil {
		logger.L(Name).Debug("error while adding custom index rule", zap.Error(err))
		return x.ErrFailedToAddCustomIndexRule
	}

	_ = index.Close()

	return nil
}

func (b *BleveScorch) ApplyIndex(namespace string, data []*variant.IndexData) error {
	p := b.path + "/" + namespace
	if !b.isPathExists(p) {
		return nil
	}

	index, err := bleveSearch.Open(p)

	if err != nil {
		logger.L(Name).Error("db opening error", zap.Error(err))
		return x.ErrFailedToApplyIndex
	}

	defer func() {
		_ = index.Close()
	}()

	batch := index.NewBatch()
	for idx := range data {
		if data[idx].Action == pb.Action_UPSERT {
			err = batch.Index(data[idx].ID, data[idx].Data)
			if err != nil {
				logger.L(Name).Debug("indexing error", zap.Error(err))
				return x.ErrFailedToCreateIndex
			}
		}

		if data[idx].Action == pb.Action_DELETE {
			batch.Delete(data[idx].ID)
		}
	}

	err = index.Batch(batch)
	if err != nil {
		logger.L(Name).Debug("batch processing error", zap.Error(err))
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

func (b *BleveScorch) GlobalSearch(_ context.Context, input *pb.GlobalSearchInput) (iface.IIndexStorageSearchResult, error) {
	indexPath := b.path + "/" + string(input.Namespace)
	index, err := bleveSearch.Open(indexPath)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = index.Close()
	}()

	var bleveQuery bleveSearchQuery.Query = bleveSearchQuery.NewMatchAllQuery()
	// TODO: implement all type of query
	if input.GetQueryString() != nil {
		bleveQuery = bleveSearchQuery.NewQueryStringQuery(input.GetQueryString().Q)
	}

	searchRequest := bleveSearch.NewSearchRequest(bleveQuery)
	searchRequest.Size = int(input.GetSize())
	searchRequest.From = int(input.GetFrom())
	searchRequest.Fields = input.GetFields()
	searchRequest.Explain = input.GetExplain()
	searchRequest.SortBy(input.GetSort())
	searchRequest.IncludeLocations = input.GetIncludeLocations()
	searchRequest.Score = input.GetScore()
	searchRequest.SearchAfter = input.GetSearchAfter()
	searchRequest.SearchBefore = input.GetSearchBefore()

	// HIGHLIGHT
	if input.GetHighlight() {
		searchRequest.Highlight = bleveSearch.NewHighlight()
		if input.GetHighlightStyle() != "" {
			style := input.GetHighlightStyle()
			searchRequest.Highlight.Style = &style
		}

		if input.GetHighlightFields() != nil {
			searchRequest.Highlight.Fields = input.GetHighlightFields()
		}
	}

	// FACET
	for _, facet := range input.GetFacets() {
		if facet.GetName() != "" && facet.GetField() != "" {
			facetRequest := bleveSearch.NewFacetRequest(facet.GetField(), int(facet.GetSize()))

			for _, dtr := range facet.GetDateTimeRangeFacets() {
				if dtr.GetName() != "" {
					start := dtr.GetStart()
					end := dtr.GetEnd()
					facetRequest.AddDateTimeRangeString(dtr.GetName(), &start, &end)
				}
			}

			for _, nr := range facet.GetNumericRangeFacets() {
				if nr.GetName() != "" {
					min := nr.GetMin()
					max := nr.GetMax()
					facetRequest.AddNumericRange(nr.GetName(), &min, &max)
				}
			}

			searchRequest.AddFacet(facet.GetName(), facetRequest)
		}
	}

	searchResult, err := index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	return &BleveSearchResult{Result: searchResult}, nil
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

	indexMapping.IndexDynamic = meta.IndexDynamic
	indexMapping.StoreDynamic = meta.StoreDynamic
	indexMapping.DocValuesDynamic = meta.DocValuesDynamic

	if meta.DefaultType != "" {
		indexMapping.DefaultType = meta.DefaultType
	}

	if meta.DefaultAnalyzer != "" {
		indexMapping.DefaultAnalyzer = meta.DefaultAnalyzer
	}

	if meta.DefaultDateTimeParser != "" {
		indexMapping.DefaultDateTimeParser = meta.DefaultDateTimeParser
	}

	if meta.DefaultField != "" {
		indexMapping.DefaultField = meta.DefaultField
	}

	if meta.TypeField != "" {
		indexMapping.TypeField = meta.TypeField
	}

	if meta.CustomAnalysis != "" {
		// Get custom analyzer from registry
		// not defined yet
	}

	for _, indexDocument := range meta.IndexDocumentList {
		documentMapping := bleveSearch.NewDocumentMapping()

		if indexDocument.DefaultAnalyzer != "" {
			documentMapping.DefaultAnalyzer = indexDocument.DefaultAnalyzer
		} else {
			documentMapping.DefaultAnalyzer = standard.Name
		}

		documentMapping.Enabled = indexDocument.Enabled
		documentMapping.Dynamic = indexDocument.Dynamic

		for _, documentField := range indexDocument.IndexFieldList {
			if !documentField.Enabled {
				disable := bleveSearch.NewDocumentMapping()
				disable.Enabled = false
				documentMapping.AddSubDocumentMapping(indexDocument.Name, disable)
				continue
			}

			var fieldMap *bleveMapping.FieldMapping
			if documentField.IndexFieldType == pb.IndexFieldType_TEXT {
				fieldMap = bleveSearch.NewTextFieldMapping()
			} else if documentField.IndexFieldType == pb.IndexFieldType_BOOLEAN {
				fieldMap = bleveSearch.NewBooleanFieldMapping()
			} else if documentField.IndexFieldType == pb.IndexFieldType_DATE_TIME {
				fieldMap = bleveSearch.NewDateTimeFieldMapping()
			} else if documentField.IndexFieldType == pb.IndexFieldType_NUMERIC {
				fieldMap = bleveSearch.NewNumericFieldMapping()
			} else if documentField.IndexFieldType == pb.IndexFieldType_GEO_POINT {
				fieldMap = bleveSearch.NewGeoPointFieldMapping()
			} else {
				fieldMap = bleveSearch.NewTextFieldMapping()
			}

			fieldMap.Name = documentField.Name
			fieldMap.Analyzer = documentField.Analyzer
			fieldMap.DocValues = documentField.DocValues
			fieldMap.Store = documentField.Store
			fieldMap.Index = documentField.Index
			fieldMap.IncludeTermVectors = documentField.IncludeTermVectors
			fieldMap.IncludeInAll = documentField.IncludeInAll

			if documentField.DateFormat != "" {
				fieldMap.DateFormat = documentField.DateFormat
			}

			documentMapping.AddFieldMappingsAt(documentField.Name, fieldMap)
		}

		indexMapping.AddDocumentMapping(indexDocument.Name, documentMapping)
	}

	return indexMapping
}
