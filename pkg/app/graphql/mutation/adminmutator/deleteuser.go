package adminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/x"
)

var DeleteUser = &graphql.Field{
	Name:        "DeleteUser",
	Description: "",
	Type:        kind.GQLProposalResponseType,
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
			return nil, gqlerrors.NewFormattedError(x.ErrInvalidOperation.Error())
		}

		if !ok {
			return nil, gqlerrors.NewFormattedError(x.ErrInvalidSourceType.Error())
		}

		pr, err := admin.DeleteUser(username)
		if err != nil {
			return nil, err
		}

		return pr, nil
	},
}
