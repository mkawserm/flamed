package nodeadminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var AddWitness = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Add new node to the cluster",
	Args: graphql.FieldConfigArgument{
		"nodeID": &graphql.ArgumentConfig{
			Description: "Node ID",
			Type:        graphql.NewNonNull(types.UInt64Type),
		},

		"address": &graphql.ArgumentConfig{
			Description: "Raft address of the node",
			Type:        graphql.NewNonNull(graphql.String),
		},

		"configChangeIndex": &graphql.ArgumentConfig{
			Description: "Config change index",
			Type:        graphql.NewNonNull(types.UInt64Type),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		nodeID := p.Args["nodeID"].(*types.UInt64)
		address := p.Args["address"].(string)
		configChangeIndex := p.Args["configChangeIndex"].(*types.UInt64)

		nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
		if !ok {
			return nil, nil
		}
		err := nodeAdmin.AddWitness(nodeID.Value(), address, configChangeIndex.Value())
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return true, nil
	},
}
