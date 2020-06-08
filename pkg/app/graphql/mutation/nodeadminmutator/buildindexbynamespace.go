package nodeadminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/utility"
)

var BuildIndexByNamespace = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Build or rebuild state database index by namespace",
	Args: graphql.FieldConfigArgument{
		"namespace": &graphql.ArgumentConfig{
			Description: "Namespace",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		namespace := p.Args["namespace"].(string)
		namespaceBytes := []byte(namespace)
		if !utility.IsNamespaceValid(namespaceBytes) {
			return nil, gqlerrors.NewFormattedError("invalid namespace")
		}

		nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
		if !ok {
			return false, nil
		}

		nodeAdmin.BuildIndexByNamespace(namespaceBytes)
		return true, nil
	},
}
