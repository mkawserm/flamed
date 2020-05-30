package global

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
)

var SearchJ = &graphql.Field{
	Name:        "SearchJ",
	Description: "Search and get json response",
	Type:        types.GQLJSONType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description: "Search input",
			Type:        graphql.NewNonNull(types.GQLSearchInputType),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})
		ctx, ok := p.Source.(*Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasGlobalSearchPermission(ctx.AccessControl) {
			return nil, gqlerrors.NewFormattedError("global search permission required")
		}

		globalSearchInput := &pb.GlobalSearchInput{
			Namespace: []byte(ctx.Query.Namespace()),
		}

		globalSearchInput.Query = &pb.GlobalSearchInput_QueryString{
			QueryString: &pb.QueryString{Q: input["query"].(string)},
		}

		if v, ok := input["size"].(int); ok {
			globalSearchInput.Size = int32(v)
		} else {
			globalSearchInput.Size = 20
		}

		if v, ok := input["from"].(int); ok {
			globalSearchInput.From = int32(v)
		} else {
			globalSearchInput.From = 0
		}

		if v, ok := input["includeLocations"].(bool); ok {
			globalSearchInput.IncludeLocations = v
		}

		if v, ok := input["explain"].(bool); ok {
			globalSearchInput.Explain = v
		}

		if v, ok := input["fields"].([]interface{}); ok {
			for _, field := range v {
				globalSearchInput.Fields = append(globalSearchInput.Fields, field.(string))
			}
		}

		if v, ok := input["sort"].([]interface{}); ok {
			for _, field := range v {
				globalSearchInput.Sort = append(globalSearchInput.Sort, field.(string))
			}
		}

		if v, ok := input["highlight"].(bool); ok {
			globalSearchInput.Highlight = v
		}

		if v, ok := input["highlightFields"].([]interface{}); ok {
			for _, field := range v {
				globalSearchInput.HighlightFields = append(globalSearchInput.HighlightFields, field.(string))
			}
		}

		if v, ok := input["highlightStyle"].(string); ok {
			globalSearchInput.HighlightStyle = v
		}

		if facets, ok := input["facets"].([]interface{}); ok {
			for _, facet := range facets {
				if v, ok := facet.(map[string]interface{}); ok {
					pf := &pb.Facet{}
					pf.Name = v["name"].(string)
					pf.Field = v["field"].(string)
					pf.Size = int32(v["size"].(int))

					if v2, ok := v["dateTimeRangeFacets"].([]interface{}); ok {
						for _, v3 := range v2 {
							if v4, ok := v3.(map[string]interface{}); ok {
								d := &pb.DateTimeRangeFacet{}
								d.Name = v4["name"].(string)
								d.Start = v4["start"].(string)
								d.End = v4["end"].(string)
								pf.DateTimeRangeFacets = append(pf.DateTimeRangeFacets, d)
							}
						}
					}

					if v2, ok := v["numericRangeFacets"].([]interface{}); ok {
						for _, v3 := range v2 {
							if v4, ok := v3.(map[string]interface{}); ok {
								d := &pb.NumericRangeFacet{}
								d.Name = v4["name"].(string)
								d.Min = v4["min"].(float64)
								d.Max = v4["max"].(float64)
								pf.NumericRangeFacets = append(pf.NumericRangeFacets, d)
							}
						}
					}

					globalSearchInput.Facets = append(globalSearchInput.Facets, pf)
				}
			}
		}

		o, err := ctx.Query.Search(globalSearchInput)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		if o != nil {
			return o.ToMap(), nil
		}

		return nil, nil
	},
}
