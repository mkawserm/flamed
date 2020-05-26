package intkeymutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLDelete = &graphql.Field{
	Name:        "GQLDelete",
	Type:        types.GQLProposalResponseType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Description: "Name",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name := p.Args["name"].(string)
		ikc, ok := p.Source.(*intkey.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasWritePermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError("write permission required")
		}

		pr, err := ikc.Client.Delete(name)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return pr, nil
	},
}
