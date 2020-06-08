package admin

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var IsAccessControlAvailable = &graphql.Field{
	Name:        "IsAccessControlAvailable",
	Description: "Is access control available?",
	Type:        graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Description: "Username",
			Type:        graphql.NewNonNull(graphql.String),
		},
		"namespace": &graphql.ArgumentConfig{
			Description: "Namespace",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username := p.Args["username"].(string)
		namespace := p.Args["namespace"].(string)
		namespaceBytes := []byte(namespace)

		//if !utility.IsUsernameValid(username) {
		//	return nil, gqlerrors.NewFormattedError("invalid username")
		//}
		//
		//if !utility.IsNamespaceValid(namespaceBytes) {
		//	return nil, gqlerrors.NewFormattedError("invalid namespace")
		//}

		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, nil
		}

		return admin.IsAccessControlAvailable(username, namespaceBytes), nil
	},
}
