package adminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var DeleteUser = &graphql.Field{
	Name:        "DeleteUser",
	Description: "",
	Type:        types.GQLProposalResponseType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Description: "Username",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username := p.Args["username"].(string)
		admin, ok := p.Source.(*flamed.Admin)

		if username == "admin" {
			return nil, gqlerrors.NewFormattedError("delete operation is not" +
				" allowed on admin user")
		}

		if !ok {
			return nil, gqlerrors.NewFormattedError("Unknown source type." +
				" FlamedContext required")
		}

		pr, err := admin.DeleteUser(username)
		if err != nil {
			return nil, err
		}

		return pr, nil
	},
}
