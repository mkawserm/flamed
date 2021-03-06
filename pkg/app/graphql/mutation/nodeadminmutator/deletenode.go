package nodeadminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var DeleteNode = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Add new node to the cluster",
	Args: graphql.FieldConfigArgument{
		"nodeID": &graphql.ArgumentConfig{
			Description: "Node ID",
			Type:        graphql.NewNonNull(kind.GQLUInt64Type),
		},
		"configChangeIndex": &graphql.ArgumentConfig{
			Description: "Config change index",
			Type:        graphql.NewNonNull(kind.GQLUInt64Type),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		nodeID := p.Args["nodeID"].(*kind.UInt64)
		configChangeIndex := p.Args["configChangeIndex"].(*kind.UInt64)

		nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
		if !ok {
			return nil, nil
		}
		err := nodeAdmin.DeleteNode(nodeID.Value(), configChangeIndex.Value())
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return true, nil
	},
}
