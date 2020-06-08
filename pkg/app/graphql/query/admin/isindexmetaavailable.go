package admin

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var IsIndexMetaAvailable = &graphql.Field{
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
}
