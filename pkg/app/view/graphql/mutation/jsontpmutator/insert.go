package jsontpmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLInsert = &graphql.Field{
	Name:        "Insert",
	Type:        types.GQLProposalResponseType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description: "Input object",
			Type:        graphql.NewNonNull(types.GQLJSONType),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})
		jsonContext, ok := p.Source.(*json.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasReadPermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError("read permission required")
		}

		if !utility.HasWritePermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError("write permission required")
		}

		pr, err := jsonContext.Client.InsertJSONMap(input)
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return pr, nil
	},
}
