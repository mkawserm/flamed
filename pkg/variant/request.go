package variant

import "context"

type LookupRequest struct {
	Query   interface{}     `json:"-"`
	Context context.Context `json:"-"`

	FamilyName    string `json:"familyName"`
	FamilyVersion string `json:"familyVersion"`
}

type SearchRequest struct {
	Query     interface{}     `json:"-"`
	Context   context.Context `json:"-"`
	Namespace []byte          `json:"namespace"`
}
