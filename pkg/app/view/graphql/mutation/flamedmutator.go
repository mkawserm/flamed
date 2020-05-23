package mutation

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var NodeAdminMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NodeAdminMutator",
	Description: "`NodeAdminMutator` gives the ability to perform administrative tasks of the cluster",
	Fields: graphql.Fields{

		// AddNode
		"addNode": &graphql.Field{
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
				err := nodeAdmin.AddNode(nodeID.Value(), address, configChangeIndex.Value())
				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}

				return true, nil
			},
		},

		// AddObserver
		"addObserver": &graphql.Field{
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
				err := nodeAdmin.AddObserver(nodeID.Value(), address, configChangeIndex.Value())
				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}

				return true, nil
			},
		},

		// AddWitness
		"addWitness": &graphql.Field{
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
		},

		// DeleteNode
		"deleteNode": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add new node to the cluster",
			Args: graphql.FieldConfigArgument{
				"nodeID": &graphql.ArgumentConfig{
					Description: "Node ID",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
				"configChangeIndex": &graphql.ArgumentConfig{
					Description: "Config change index",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nodeID := p.Args["nodeID"].(*types.UInt64)
				configChangeIndex := p.Args["configChangeIndex"].(*types.UInt64)

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
		},
	},
})

var FlamedMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "FlamedMutator",
	Description: "`FlamedMutator` gives the ability to perform any tasks related to the cluster",
	Fields: graphql.Fields{

		// NodeAdminMutator
		"nodeAdminMutator": &graphql.Field{
			Name:        "NodeAdminMutator",
			Type:        NodeAdminMutatorType,
			Description: "Perform administrative tasks using NodeAdmin by clusterID",
			Args: graphql.FieldConfigArgument{
				"clusterID": &graphql.ArgumentConfig{
					Description: "Cluster ID",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*types.UInt64)
				fc, ok := p.Source.(*flamedContext.FlamedContext)
				if !ok {
					return nil, nil
				}

				if !fc.Flamed.IsClusterIDAvailable(clusterID.Value()) {
					return nil,
						gqlerrors.NewFormattedError(
							fmt.Sprintf("clusterID [%d] is not available", clusterID.Value()))
				}
				return fc.Flamed.NewNodeAdmin(clusterID.Value(), fc.GlobalRequestTimeout), nil
			},
		},

		// AdminMutator

	},
})

func FlamedMutator(flamedContext *flamedContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "FlamedMutator",
		Type:        FlamedMutatorType,
		Description: "Flamed mutator helps to modify cluster",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//TODO: Request must be authenticated

			return flamedContext, nil
		},
	}
}
