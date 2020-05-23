package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/context"
	"sync/atomic"
)

func Increment(_ *context.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Int,
		Description: "Increment an in-memory integer by one",
		Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
			atomic.AddInt32(&i, 1)
			return i, nil
		},
	}
}
