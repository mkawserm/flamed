package kind

import "github.com/graphql-go/graphql"

var GQLNumericRangeFacetInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "NumericRangeFacetInput",
	Description: "Numeric range facet input",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"min": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"max": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
	},
})

var GQLDateTimeRangeFacetInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "DateTimeRangeFacetInput",
	Description: "Date time range facet input",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"start": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"end": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

var GQLFacetInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "FacetInput",
	Description: "Facet input",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"field": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"size": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},

		"dateTimeRangeFacets": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(GQLDateTimeRangeFacetInputType),
		},
		"numericRangeFacets": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(GQLNumericRangeFacetInputType),
		},
	},
})

var GQLSearchInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "SearchInput",
		Description: "Search input",
		Fields: graphql.InputObjectConfigFieldMap{
			"query": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"from": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"size": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"includeLocations": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"explain": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"fields": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"sort": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"highlight": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"highlightStyle": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"highlightFields": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"searchAfter": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"searchBefore": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},

			"facets": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(GQLFacetInputType),
			},
		},
	},
)
