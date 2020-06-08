package intkeytp

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/types"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLGetIntKeyList = &graphql.Field{
	Name:        "GetIntKeyList",
	Type:        graphql.NewList(types.GQLIntKeyStateType),
	Description: "",

	Args: graphql.FieldConfigArgument{
		"nameList": &graphql.ArgumentConfig{
			Description: "Name list",
			Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		nameList := p.Args["nameList"].([]interface{})
		ikc, ok := p.Source.(*intkey.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasReadPermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError("read permission required")
		}

		intKeyStateList, err := ikc.Client.GetIntKeyStateList(nameList)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return intKeyStateList, nil
	},
}
