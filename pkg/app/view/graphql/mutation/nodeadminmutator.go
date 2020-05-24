package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/spf13/viper"
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

		//RunStorageGC
		"runStorageGC": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Run storage garbage collector",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
				if !ok {
					return false, nil
				}
				nodeAdmin.RunStorageGC()
				return true, nil
			},
		},

		//BuildIndex
		"buildIndex": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Build or rebuild full state database index",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
				if !ok {
					return false, nil
				}
				nodeAdmin.BuildIndex()
				return true, nil
			},
		},

		// BuildIndexByNamespace
		"buildIndexByNamespace": &graphql.Field{
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
		},

		// RequestSnapshot
		"requestSnapshot": &graphql.Field{
			Type:        types.UInt64Type,
			Description: "Add new node to the cluster",
			Args: graphql.FieldConfigArgument{
				"compactionOverhead": &graphql.ArgumentConfig{
					Description: "Compaction overhead",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
				"exported": &graphql.ArgumentConfig{
					Description: "Will the snapshot be exported?",
					Type:        graphql.NewNonNull(graphql.Boolean),
				},
				"overrideCompactionOverhead": &graphql.ArgumentConfig{
					Description: "Compaction overhead override flag",
					Type:        graphql.NewNonNull(graphql.Boolean),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				compactionOverhead := p.Args["compactionOverhead"].(*types.UInt64)
				exported := p.Args["exported"].(bool)
				overrideCompactionOverhead := p.Args["overrideCompactionOverhead"].(bool)
				exportPath := viper.GetString(constant.StoragePath) + "/snapshot"

				nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
				if !ok {
					return nil, nil
				}
				n, err := nodeAdmin.RequestSnapshot(compactionOverhead.Value(),
					exportPath,
					exported,
					overrideCompactionOverhead)

				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}

				return types.NewUInt64FromUInt64(n), nil
			},
		},
	},
})
