package intkeytp

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/app/utility"
	fContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
)

var GQLIntKeyTPType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "IntKeyTP",
	Description: "Integer key transaction processor",
	Fields: graphql.Fields{
		"getIntKey":     GQLGetIntKey,
		"getIntKeyList": GQLGetIntKeyList,
	},
})

func IntKeyTP(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "IntKeyTP",
		Type:        GQLIntKeyTPType,
		Description: "Integer key transaction processor",

		Args: graphql.FieldConfigArgument{
			"clusterID": &graphql.ArgumentConfig{
				Description: "Cluster ID",
				Type:        graphql.NewNonNull(kind.GQLUInt64Type),
			},
			"namespace": &graphql.ArgumentConfig{
				Description: "Namespace",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			clusterID, namespace, accessControl, err := utility.GraphQLAuthCheck(p, flamedContext)
			if err != nil {
				return nil, err
			}

			intKeyClient := &intkey.Client{}
			err = intKeyClient.Setup(clusterID,
				namespace,
				flamedContext.Flamed(),
				flamedContext.GlobalRequestTimeout())

			if err != nil {
				return nil, gqlerrors.NewFormattedError(err.Error())
			}

			return &intkey.Context{
				AccessControl: accessControl,
				Client:        intKeyClient,
			}, nil
		},
	}
}
