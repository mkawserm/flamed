package query

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/context"
)

func IsLive(_ *context.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return true, nil
		},
		Description: "Service availability status",
	}
}
