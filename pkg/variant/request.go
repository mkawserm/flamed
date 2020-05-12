package variant

import "context"

type LookupRequest struct {
	Query   interface{}     `json:"-"`
	Context context.Context `json:"-"`

	Namespace []byte `json:"namespace"`
	Family    string `json:"family"`
	Version   string `json:"version"`
}

type SearchRequest struct {
	Query     interface{}     `json:"-"`
	Context   context.Context `json:"-"`
	Namespace []byte          `json:"namespace"`
}
