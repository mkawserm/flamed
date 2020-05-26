package admin

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var GetUser = &graphql.Field{
	Name:        "GetUser",
	Description: "Get user by username",
	Type:        types.GQLUserType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Description: "Username",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username := p.Args["username"].(string)

		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, nil
		}

		user, err := admin.GetUser(username)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return user, nil
	},
}
