package intkeytpmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLDecrement = &graphql.Field{
	Name:        "Decrement",
	Type:        kind.GQLProposalResponseType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Description: "Name",
			Type:        graphql.NewNonNull(graphql.String),
		},

		"value": &graphql.ArgumentConfig{
			Description: "Value",
			Type:        graphql.NewNonNull(kind.GQLUInt64Type),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name := p.Args["name"].(string)
		value := p.Args["value"].(*kind.UInt64)

		ikc, ok := p.Source.(*intkey.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasUpdatePermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError("update permission required")
		}

		pr, err := ikc.Client.Decrement(name, value.Value())

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return pr, nil
	},
}
