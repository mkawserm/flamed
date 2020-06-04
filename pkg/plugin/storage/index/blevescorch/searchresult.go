package blevescorch

import (
	"encoding/json"
	bleveSearch "github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/utility"
	"time"
)

type BleveDateRangeFacet struct {
	mDRF *search.DateRangeFacet
}

func (b *BleveDateRangeFacet) Name() string {
	return b.mDRF.Name
}

func (b *BleveDateRangeFacet) Start() string {
	var s = *b.mDRF.Start
	return s
}

func (b *BleveDateRangeFacet) End() string {
	var s = *b.mDRF.End
	return s
}

func (b *BleveDateRangeFacet) Count() int {
	return b.mDRF.Count
}

type BleveNumericRangeFacet struct {
	mNRF *search.NumericRangeFacet
}

func (b *BleveNumericRangeFacet) Name() string {
	return b.mNRF.Name
}

func (b *BleveNumericRangeFacet) Min() float64 {
	return *b.mNRF.Min
}

func (b *BleveNumericRangeFacet) Max() float64 {
	return *b.mNRF.Max
}

func (b *BleveNumericRangeFacet) Count() int {
	return b.mNRF.Count
}

type BleveTerm struct {
	mTerm *search.TermFacet
}

func (b *BleveTerm) Term() string {
	return b.mTerm.Term
}

func (b *BleveTerm) Count() int {
	return b.mTerm.Count
}

type BleveFacet struct {
	mName        string
	mFacetResult *search.FacetResult
}

func (b *BleveFacet) Name() string {
	return b.mName
}

func (b *BleveFacet) Field() string {
	return b.mFacetResult.Field
}

func (b *BleveFacet) Total() int {
	return b.mFacetResult.Total
}

func (b *BleveFacet) Missing() int {
	return b.mFacetResult.Missing
}

func (b *BleveFacet) Other() int {
	return b.mFacetResult.Other
}

func (b *BleveFacet) Terms() []iface.ITerm {
	if len(b.mFacetResult.Terms) == 0 {
		return nil
	}

	terms := make([]iface.ITerm, 0, len(b.mFacetResult.Terms))
	for _, doc := range b.mFacetResult.Terms {
		terms = append(terms, &BleveTerm{mTerm: doc})
	}

	return terms
}

func (b *BleveFacet) NumericRanges() []iface.INumericRangeFacet {
	if len(b.mFacetResult.NumericRanges) == 0 {
		return nil
	}

	numericRangeFacets := make([]iface.INumericRangeFacet, 0, len(b.mFacetResult.NumericRanges))
	for _, doc := range b.mFacetResult.NumericRanges {
		numericRangeFacets = append(numericRangeFacets, &BleveNumericRangeFacet{mNRF: doc})
	}

	return numericRangeFacets
}

func (b *BleveFacet) DateRanges() []iface.IDateRangeFacet {
	if len(b.mFacetResult.DateRanges) == 0 {
		return nil
	}

	dateRangeFacets := make([]iface.IDateRangeFacet, 0, len(b.mFacetResult.DateRanges))
	for _, doc := range b.mFacetResult.DateRanges {
		dateRangeFacets = append(dateRangeFacets, &BleveDateRangeFacet{mDRF: doc})
	}

	return dateRangeFacets
}

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
	Result *bleveSearch.SearchResult
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
	documents := make([]iface.IDocument, 0, len(b.Result.Hits))
	for _, doc := range b.Result.Hits {
		documents = append(documents, &BleveSearchDocument{DocMatch: doc})
	}
	return documents
}

func (b *BleveSearchResponse) Facets() []iface.IFacet {
	if len(b.Result.Facets) == 0 {
		return nil
	}

	facets := make([]iface.IFacet, 0, len(b.Result.Facets))
	for name, doc := range b.Result.Facets {
		facets = append(facets, &BleveFacet{mName: name, mFacetResult: doc})
	}
	return facets
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
	return &BleveSearchResponse{
		Result: b.Result,
	}
}
