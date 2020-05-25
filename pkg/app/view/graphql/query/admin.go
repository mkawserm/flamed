package query

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var AdminType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Admin",
	Description: "`Admin` provides user,index meta,access control related information related to the cluster",
	Fields: graphql.Fields{

		// GetUser
		"getUser": &graphql.Field{
			Name:        "GetUser",
			Description: "Get user by username",
			Type:        types.UserType,
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
		},

		// IsUserAvailable
		"isUserAvailable": &graphql.Field{
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
		},

		"getAccessControl": &graphql.Field{
			Name:        "GetAccessControl",
			Description: "Get access control information of a user",
			Type:        types.AccessControlType,
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
		},

		// IsAccessControlAvailable
		"isAccessControlAvailable": &graphql.Field{
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
		},

		// IsIndexMetaAvailable
		"isIndexMetaAvailable": &graphql.Field{
			Name:        "IsIndexMetaAvailable",
			Description: "Is index meta available?",
			Type:        graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"namespace": &graphql.ArgumentConfig{
					Description: "Namespace",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				namespace := p.Args["namespace"].(string)
				namespaceBytes := []byte(namespace)

				//if !utility.IsNamespaceValid(namespaceBytes) {
				//	return nil, gqlerrors.NewFormattedError("invalid namespace")
				//}

				admin, ok := p.Source.(*flamed.Admin)
				if !ok {
					return nil, nil
				}

				return admin.IsIndexMetaAvailable(namespaceBytes), nil
			},
		},

		"getIndexMeta": &graphql.Field{
			Name:        "GetIndexMeta",
			Description: "",
			Type:        types.IndexMetaType,
			Args: graphql.FieldConfigArgument{
				"namespace": &graphql.ArgumentConfig{
					Description: "Namespace",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				namespace := p.Args["namespace"].(string)
				namespaceBytes := []byte(namespace)

				admin, ok := p.Source.(*flamed.Admin)
				if !ok {
					return nil, nil
				}

				indexMeta, err := admin.GetIndexMeta(namespaceBytes)

				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}

				return indexMeta, nil
			},
		},
	},
})
