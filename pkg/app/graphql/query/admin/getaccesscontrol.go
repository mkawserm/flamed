package admin

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/types"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var GetAccessControl = &graphql.Field{
	Name:        "GetAccessControl",
	Description: "Get access control information of a user",
	Type:        types.GQLAccessControlType,
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

		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, nil
		}

		accessControl, err := admin.GetAccessControl(username, namespaceBytes)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return accessControl, nil
	},
}
