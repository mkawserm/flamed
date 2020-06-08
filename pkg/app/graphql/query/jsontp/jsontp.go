package jsontp

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/types"
	"github.com/mkawserm/flamed/pkg/app/utility"
	fContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/tp/json"
)

var GQLJSONTPType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "JSONTP",
	Description: "JSON transaction processor",
	Fields: graphql.Fields{
		"get":              GQLGet,
		"getList":          GQLGetList,
		"getListByIndexID": GQLGetListByIndexID,
	},
})

func JSONTP(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "JSONTP",
		Type:        GQLJSONTPType,
		Description: "JSON transaction processor",

		Args: graphql.FieldConfigArgument{
			"clusterID": &graphql.ArgumentConfig{
				Description: "Cluster ID",
				Type:        graphql.NewNonNull(types.GQLUInt64Type),
			},
			"namespace": &graphql.ArgumentConfig{
				Description: "Namespace",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			clusterID, namespace, accessControl, err := utility.AuthCheck(p, flamedContext)
			if err != nil {
				return nil, err
			}

			client := &json.Client{}
			err = client.Setup(clusterID,
				namespace,
				flamedContext.Flamed(),
				flamedContext.GlobalRequestTimeout())

			if err != nil {
				return nil, gqlerrors.NewFormattedError(err.Error())
			}

			return &json.Context{
				AccessControl: accessControl,
				Client:        client,
			}, nil
		},
	}
}
