package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/context"
)

func Increment(_ *context.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Int,
		Description: "Increment an in-memory integer by one",
		Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
			i++
			return i, nil
		},
	}
}
