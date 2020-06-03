package blevescorch

import (
	"encoding/json"
	bleveSearch "github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/utility"
	"time"
)

type BleveSearchDocument struct {
	DocMatch *search.DocumentMatch
}

func (b *BleveSearchDocument) ID() string {
	return b.DocMatch.ID
}

func (b *BleveSearchDocument) Score() float64 {
	return b.DocMatch.Score
}

func (b *BleveSearchDocument) Index() string {
	return b.DocMatch.Index
}

type BleveSearchResponse struct {
	Result    *bleveSearch.SearchResult
	Documents []iface.IDocument
}

func (b *BleveSearchResponse) Total() uint64 {
	return b.Result.Total
}

func (b *BleveSearchResponse) MaxScore() float64 {
	return b.Result.MaxScore
}

func (b *BleveSearchResponse) Took() time.Duration {
	return b.Result.Took
}

func (b *BleveSearchResponse) Hits() []iface.IDocument {
	return b.Documents
}

type BleveSearchResult struct {
	Result *bleveSearch.SearchResult
}

func (b *BleveSearchResult) RawResult() interface{} {
	return b.Result
}

func (b *BleveSearchResult) ToMap() map[string]interface{} {
	if b.Result == nil {
		return nil
	}

	output := map[string]interface{}{}
	data, _ := utility.LowerCamelCaseMarshaller{Value: b.Result}.MarshalJSON()
	//fmt.Println(string(data))
	_ = json.Unmarshal(data, &output)
	return output
}

func (b *BleveSearchResult) ToBytes() []byte {
	if b.Result == nil {
		return nil
	}

	output, err := utility.LowerCamelCaseMarshaller{Value: b.Result}.MarshalJSON()
	if err != nil {
		return nil
	}
	return output
}

func (b *BleveSearchResult) ToSearchResponse() iface.ISearchResponse {
	documents := make([]iface.IDocument, 0, len(b.Result.Hits))
	for _, doc := range b.Result.Hits {
		documents = append(documents, &BleveSearchDocument{DocMatch: doc})
	}

	return &BleveSearchResponse{
		Result:    b.Result,
		Documents: documents,
	}
}
