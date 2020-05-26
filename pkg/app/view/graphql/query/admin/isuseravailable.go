package admin

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var IsUserAvailable = &graphql.Field{
	Name:        "IsUserAvailable",
	Description: "Is user available?",
	Type:        graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Description: "Username",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username := p.Args["username"].(string)
		//if !utility.IsUsernameValid(username) {
		//	return nil, gqlerrors.NewFormattedError("invalid username")
		//}
		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, nil
		}

		return admin.IsUserAvailable(username), nil
	},
}
