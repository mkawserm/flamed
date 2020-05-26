package admin

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var GetIndexMeta = &graphql.Field{
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
}