package query

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var NodeAdminType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NodeAdmin",
	Description: "`NodeAdmin` provides all administrative information related to the cluster",
	Fields: graphql.Fields{
		"leaderID": &graphql.Field{
			Type:        types.UInt64Type,
			Description: "Current leader id of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _, _ := nodeAdmin.GetLeaderID()
					return types.NewUInt64FromUInt64(id), nil
				}

				return nil, nil
			},
		},

		"appliedIndex": &graphql.Field{
			Type:        types.UInt64Type,
			Description: "Current applied raft index of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _ := nodeAdmin.GetAppliedIndex()
					return types.NewUInt64FromUInt64(id), nil
				}
				return nil, nil
			},
		},

		"hasNodeInfo": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Has node information with the provided nodeID",
			Args: graphql.FieldConfigArgument{
				"nodeID": &graphql.ArgumentConfig{
					Description: "Node id",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nodeID := p.Args["nodeID"].(*types.UInt64)
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					return nodeAdmin.HasNodeInfo(nodeID.Value()), nil
				}

				return nil, nil
			},
		},

		"clusterMembership": &graphql.Field{
			Type:        types.ClusterMembershipType,
			Description: "Cluster membership information",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					m, err := nodeAdmin.GetClusterMembership()
					if err != nil {
						return nil, nil
					}
					return m, nil
				}

				return nil, nil
			},
		},
	},
})

var AdminType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Admin",
	Description: "`Admin` provides user,index meta,access control related information related to the cluster",
	Fields: graphql.Fields{
		"isUserAvailable": &graphql.Field{
			Name:        "IsUserAvailable",
			Description: "Is user available?",
			Type:        graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Description: "Username",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				username := p.Args["username"].(string)
				admin, ok := p.Source.(*flamed.Admin)
				if !ok {
					return nil, nil
				}

				return admin.IsUserAvailable(username), nil
			},
		},
	},
})

var FlamedType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Flamed",
	Description: "`Flamed` provides all information related to the cluster",
	Fields: graphql.Fields{
		"nodeHostInfo": &graphql.Field{
			Type:        types.NodeHostInfoType,
			Description: "Node host information",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fc, ok := p.Source.(*flamedContext.FlamedContext)
				if !ok {
					return nil, nil
				}
				return fc.Flamed.GetNodeHostInfo(), nil
			},
		},

		"nodeAdmin": &graphql.Field{
			Name:        "NodeAdmin",
			Type:        NodeAdminType,
			Description: "Query administrative information from NodeAdmin by clusterID",
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

		"admin": &graphql.Field{
			Name:        "Admin",
			Type:        AdminType,
			Description: "Query user,index meta,access control related information from Admin by clusterID",
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

				return fc.Flamed.NewAdmin(clusterID.Value(), fc.GlobalRequestTimeout), nil
			},
		},
	},
})

func Flamed(flamedContext *flamedContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "Flamed",
		Type:        FlamedType,
		Description: "Query flamed for all kinds administrative information",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//TODO: Request must be authenticated

			return flamedContext, nil
		},
	}
}
