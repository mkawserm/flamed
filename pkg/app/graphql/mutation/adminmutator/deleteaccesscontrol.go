package adminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/types"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var DeleteAccessControl = &graphql.Field{
	Name:        "DeleteAccessControl",
	Description: "",
	Type:        types.GQLProposalResponseType,
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
		namespace := []byte(p.Args["namespace"].(string))
		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, gqlerrors.NewFormattedError("Unknown source type." +
				" FlamedContext required")
		}

		pr, err := admin.DeleteAccessControl(namespace, username)
		if err != nil {
			return nil, err
		}

		return pr, nil
	},
}
