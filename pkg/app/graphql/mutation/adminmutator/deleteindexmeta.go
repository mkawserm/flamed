package adminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/x"
)

var DeleteIndexMeta = &graphql.Field{
	Name:        "DeleteIndexMeta",
	Description: "",
	Type:        kind.GQLProposalResponseType,
	Args: graphql.FieldConfigArgument{
		"namespace": &graphql.ArgumentConfig{
			Description: "Namespace",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		namespace := []byte(p.Args["namespace"].(string))
		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, gqlerrors.NewFormattedError(x.ErrInvalidSourceType.Error())
		}

		pr, err := admin.DeleteIndexMeta(namespace)
		if err != nil {
			return nil, err
		}

		return pr, nil
	},
}
