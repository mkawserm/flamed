package globaloperation

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
)

var Iterate = &graphql.Field{
	Name:        "Iterate",
	Description: "Iterate data from the state store",
	Type:        graphql.NewList(kind.GQLStateEntryResponse),
	Args: graphql.FieldConfigArgument{
		"from": &graphql.ArgumentConfig{
			Description: "State address or address prefix in hex string format",
			Type:        graphql.String,
		},
		"prefix": &graphql.ArgumentConfig{
			Description: "State address prefix in hex string format",
			Type:        graphql.String,
		},
		"limit": &graphql.ArgumentConfig{
			Description: "Limit",
			Type:        graphql.NewNonNull(kind.GQLUInt64Type),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		from := ""
		if p.Args["from"] != nil {
			from = p.Args["from"].(string)
		}

		prefix := ""
		if p.Args["prefix"] != nil {
			prefix = p.Args["prefix"].(string)
		}

		limit := p.Args["limit"].(*kind.UInt64)
		ctx, ok := p.Source.(*context.GlobalOperationContext)
		if !ok {
			return nil, nil
		}

		if !utility.HasGlobalIteratePermission(ctx.AccessControl) {
			return nil, gqlerrors.NewFormattedError(x.ErrGlobalIteratePermissionRequired.Error())
		}

		o, err := ctx.GlobalOperation.Iterate(from, prefix, limit.Value())
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return o, nil
	},
}
