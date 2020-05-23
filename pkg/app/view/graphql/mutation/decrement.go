package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/context"
	"sync/atomic"
)

func Decrement(_ *context.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "Decrement",
		Type:        graphql.Int,
		Description: "Decrement an in-memory integer by one",
		Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
			atomic.AddInt32(&i, -1)
			return i, nil
		},
	}
}
