package global

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/utility"
)

var Iterate = &graphql.Field{
	Name:        "Iterate",
	Description: "Iterate data from the state store",
	Type:        graphql.NewList(types.GQLStateEntryResponse),
	Args: graphql.FieldConfigArgument{
		"from": &graphql.ArgumentConfig{
			Description: "State address in hex string format",
			Type:        graphql.String,
		},
		"limit": &graphql.ArgumentConfig{
			Description: "Limit",
			Type:        graphql.NewNonNull(types.GQLUInt64Type),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		from := ""
		if p.Args["from"] != nil {
			from = p.Args["from"].(string)
		}

		limit := p.Args["limit"].(*types.UInt64)
		ctx, ok := p.Source.(*Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasGlobalIteratePermission(ctx.AccessControl) {
			return nil, gqlerrors.NewFormattedError("global iterate permission required")
		}

		o, err := ctx.Query.Iterate(from, limit.Value())
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return o, nil
	},
}
