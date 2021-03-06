package query

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/app/graphql/query/admin"
	"github.com/mkawserm/flamed/pkg/app/graphql/query/nodeadmin"
	fContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/x"
)

var FlamedType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Flamed",
	Description: "`Flamed` provides all information related to the cluster",
	Fields: graphql.Fields{
		"nodeHostInfo": &graphql.Field{
			Type:        kind.NodeHostInfoType,
			Description: "Node host information",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fc, ok := p.Source.(*fContext.FlamedContext)
				if !ok {
					return nil, nil
				}
				return fc.Flamed().GetNodeHostInfo(), nil
			},
		},

		"nodeAdmin": &graphql.Field{
			Name:        "NodeAdmin",
			Type:        nodeadmin.GQLNodeAdminType,
			Description: "Global administrative information from NodeAdmin by clusterID",
			Args: graphql.FieldConfigArgument{
				"clusterID": &graphql.ArgumentConfig{
					Description: "Cluster ID",
					Type:        graphql.NewNonNull(kind.GQLUInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*kind.UInt64)
				fc, ok := p.Source.(*fContext.FlamedContext)
				if !ok {
					return nil, nil
				}

				if !fc.Flamed().IsClusterIDAvailable(clusterID.Value()) {
					return nil,
						gqlerrors.NewFormattedError(x.ErrClusterIsNotAvailable.Error())
				}
				return fc.Flamed().NewNodeAdmin(clusterID.Value(), fc.GlobalRequestTimeout()), nil
			},
		},

		"admin": &graphql.Field{
			Name:        "Admin",
			Type:        admin.GQLAdminType,
			Description: "Global user,index meta,access control related information from Admin by clusterID",
			Args: graphql.FieldConfigArgument{
				"clusterID": &graphql.ArgumentConfig{
					Description: "Cluster ID",
					Type:        graphql.NewNonNull(kind.GQLUInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*kind.UInt64)
				fc, ok := p.Source.(*fContext.FlamedContext)
				if !ok {
					return nil, nil
				}

				if !fc.Flamed().IsClusterIDAvailable(clusterID.Value()) {
					return nil,
						gqlerrors.NewFormattedError(x.ErrClusterIsNotAvailable.Error())
				}

				return fc.Flamed().NewAdmin(clusterID.Value(), fc.GlobalRequestTimeout()), nil
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

			gqlContext := p.Context.Value("GraphQLContext").(*fContext.AuthContext)
			if !gqlContext.AuthenticateSuperUser(flamedContext.Flamed().NewAdmin(
				1,
				flamedContext.GlobalRequestTimeout())) {
				return nil, gqlerrors.NewFormattedError(x.ErrAccessDenied.Error())
			}

			return flamedContext, nil
		},
	}
}
