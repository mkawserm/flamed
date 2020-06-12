package jsontpmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLBatchProcess = &graphql.Field{
	Name:        "BatchProcess",
	Type:        kind.GQLProposalResponseType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description: "Input JSON as list. " +
				"supported action: INSERT, UPSERT, UPDATE, MERGE, DELETE " +
				"example `[{action:\"INSERT\",data:{id:1,value:100}}," +
				"{action:\"UPSERT\",data:{id:2,value:5}}]`",
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(kind.GQLJSONType))),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].([]interface{})
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

		if !utility.HasUpdatePermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError("update permission required")
		}

		if !utility.HasDeletePermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError("delete permission required")
		}

		batch := jsonContext.Client.NewBatch()

		for _, si := range input {
			sim := si.(map[string]interface{})
			action := sim["action"].(string)
			data := sim["data"].(map[string]interface{})
			switch action {
			case "INSERT":
				err := batch.InsertJSONMap(data)
				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}
			case "MERGE":
				err := batch.MergeJSONMap(data)
				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}
			case "UPDATE":
				err := batch.UpdateJSONMap(data)
				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}
			case "UPSERT":
				err := batch.UpsertJSONMap(data)
				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}
			case "DELETE":
				err := batch.DeleteJSONMap(data)
				if err != nil {
					return nil, gqlerrors.NewFormattedError(err.Error())
				}
			}
		}

		pr, err := jsonContext.Client.ApplyBatch(batch)
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return pr, nil
	},
}
