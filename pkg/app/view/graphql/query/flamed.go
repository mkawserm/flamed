package query

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query/admin"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query/nodeadmin"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	fContext "github.com/mkawserm/flamed/pkg/context"
)

var FlamedType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Flamed",
	Description: "`Flamed` provides all information related to the cluster",
	Fields: graphql.Fields{
		"nodeHostInfo": &graphql.Field{
			Type:        types.NodeHostInfoType,
			Description: "Node host information",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fc, ok := p.Source.(*fContext.FlamedContext)
				if !ok {
					return nil, nil
				}
				return fc.Flamed.GetNodeHostInfo(), nil
			},
		},

		"nodeAdmin": &graphql.Field{
			Name:        "NodeAdmin",
			Type:        nodeadmin.GQLNodeAdminType,
			Description: "Query administrative information from NodeAdmin by clusterID",
			Args: graphql.FieldConfigArgument{
				"clusterID": &graphql.ArgumentConfig{
					Description: "Cluster ID",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*types.UInt64)
				fc, ok := p.Source.(*fContext.FlamedContext)
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
			Type:        admin.GQLAdminType,
			Description: "Query user,index meta,access control related information from Admin by clusterID",
			Args: graphql.FieldConfigArgument{
				"clusterID": &graphql.ArgumentConfig{
					Description: "Cluster ID",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*types.UInt64)
				fc, ok := p.Source.(*fContext.FlamedContext)
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

func Flamed(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "Flamed",
		Type:        FlamedType,
		Description: "Query flamed for all kinds administrative information",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if p.Context.Value("GraphQLContext") == nil {
				return nil, nil
			}

			gqlContext := p.Context.Value("GraphQLContext").(*fContext.GraphQLContext)
			if !gqlContext.AuthenticateSuperUser(flamedContext.Flamed.NewAdmin(
				1,
				flamedContext.GlobalRequestTimeout)) {
				return nil, gqlerrors.NewFormattedError("Access denied. Only super user can access")
			}

			return flamedContext, nil
		},
	}
}
