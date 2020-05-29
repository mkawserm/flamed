package blevescorch

import (
	"encoding/json"
	bleveSearch "github.com/blevesearch/bleve"
	"github.com/mkawserm/flamed/pkg/utility"
)

type BleveSearchResult struct {
	Result *bleveSearch.SearchResult
}

func (b *BleveSearchResult) RawResult() interface{} {
	return b.Result
}

func (b *BleveSearchResult) ToMap() map[string]interface{} {
	output := map[string]interface{}{}
	data, _ := utility.LowerCamelCaseMarshaller{Value: b.Result}.MarshalJSON()
	_ = json.Unmarshal(data, &output)
	return output
}

func (b *BleveSearchResult) ToBytes() []byte {
	output, err := utility.LowerCamelCaseMarshaller{Value: b.Result}.MarshalJSON()
	if err != nil {
		return nil
	}
	return output
}
