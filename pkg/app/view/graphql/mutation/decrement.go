package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/context"
)

func Decrement(_ *context.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Int,
		Description: "Decrement an in-memory integer by one",
		Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
			i--
			return i, nil
		},
	}
}
