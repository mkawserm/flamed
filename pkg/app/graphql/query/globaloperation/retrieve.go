package globaloperation

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
)

var Retrieve = &graphql.Field{
	Name:        "Retrieve",
	Description: "Retrieve data from the state store",
	Type:        graphql.NewList(kind.GQLStateEntryResponse),
	Args: graphql.FieldConfigArgument{
		"addresses": &graphql.ArgumentConfig{
			Description: "Address in hex string format",
			Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		addresses := p.Args["addresses"].([]interface{})
		ctx, ok := p.Source.(*context.GlobalOperationContext)
		if !ok {
			return nil, nil
		}

		if !utility.HasGlobalRetrievePermission(ctx.AccessControl) {
			return nil, gqlerrors.NewFormattedError(x.ErrGlobalRetrievePermissionRequired.Error())
		}

		o, err := ctx.GlobalOperation.Retrieve(addresses)
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return o, nil
	},
}
