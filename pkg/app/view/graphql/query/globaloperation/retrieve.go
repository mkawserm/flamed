package globaloperation

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/utility"
)

var Retrieve = &graphql.Field{
	Name:        "Retrieve",
	Description: "Retrieve data from the state store",
	Type:        graphql.NewList(types.GQLStateEntryResponse),
	Args: graphql.FieldConfigArgument{
		"addresses": &graphql.ArgumentConfig{
			Description: "Address in hex string format",
			Type:        graphql.NewList(graphql.String),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		addresses := p.Args["addresses"].([]interface{})
		ctx, ok := p.Source.(*Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasGlobalRetrievePermission(ctx.AccessControl) {
			return nil, gqlerrors.NewFormattedError("globaloperation retrieve permission required")
		}

		o, err := ctx.Query.Retrieve(addresses)
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return o, nil
	},
}
