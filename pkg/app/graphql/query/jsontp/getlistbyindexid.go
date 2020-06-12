package jsontp

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLGetListByIndexID = &graphql.Field{
	Name:        "GetListByIndexID",
	Type:        kind.GQLJSONType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"idList": &graphql.ArgumentConfig{
			Description: "List of index id",
			Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ids := p.Args["idList"].([]interface{})

		ikc, ok := p.Source.(*json.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasReadPermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError("read permission required")
		}

		obj, err := ikc.Client.GetListByIndexID(ids)
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return obj, nil
	},
}
