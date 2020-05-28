package jsontp

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLGetList = &graphql.Field{
	Name:        "GetList",
	Type:        types.GQLJSONType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"idList": &graphql.ArgumentConfig{
			Description: "ID",
			Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.ID))),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idList := p.Args["idList"].([]interface{})

		idStringList := make([]string, 0, len(idList))
		for _, n := range idList {
			idStringList = append(idStringList, n.(string))
		}

		ikc, ok := p.Source.(*json.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasReadPermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError("read permission required")
		}

		obj, err := ikc.Client.GetList(idStringList)
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return obj, nil
	},
}
