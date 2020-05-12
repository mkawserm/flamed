package variant

import "context"

type LookupRequest struct {
	Query   interface{}     `json:"-"`
	Context context.Context `json:"-"`

	Namespace []byte `json:"namespace"`
	TPFamily  string `json:"tpFamily"`
	TPVersion string `json:"tpVersion"`
}

type SearchRequest struct {
	Query     interface{}     `json:"-"`
	Context   context.Context `json:"-"`
	Namespace []byte          `json:"namespace"`
}
