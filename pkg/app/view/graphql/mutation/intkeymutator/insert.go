package intkeymutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLInsert = &graphql.Field{
	Name:        "GQLInsert",
	Type:        types.GQLProposalResponseType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Description: "Name",
			Type:        graphql.NewNonNull(graphql.String),
		},

		"value": &graphql.ArgumentConfig{
			Description: "Value",
			Type:        graphql.NewNonNull(types.GQLUInt64Type),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name := p.Args["name"].(string)
		value := p.Args["value"].(*types.UInt64)

		ikc, ok := p.Source.(*intkey.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasWritePermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError("write permission required")
		}

		pr, err := ikc.Client.Insert(name, value.Value())

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return pr, nil
	},
}
