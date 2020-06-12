package jsontpmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLUpsert = &graphql.Field{
	Name:        "Upsert",
	Type:        kind.GQLProposalResponseType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description: "Input object",
			Type:        graphql.NewNonNull(kind.GQLJSONType),
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

		if !utility.HasUpdatePermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError("update permission required")
		}

		if !utility.HasWritePermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError("write permission required")
		}

		pr, err := jsonContext.Client.UpsertJSONMap(input)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return pr, nil
	},
}
