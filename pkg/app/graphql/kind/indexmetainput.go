package kind

import (
	"github.com/graphql-go/graphql"
)

var GQLIndexDocumentInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "IndexDocumentInput",
	Description: "`IndexDocumentInput`",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":            &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"enabled":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
		"defaultAnalyzer": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"dynamic":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
		"indexFieldList": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(GQLIndexFieldInputType)),
		},
	},
})

var GQLIndexMetaInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "IndexMetaInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"namespace": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"version":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(GQLUInt64Type)},
		"enabled":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
		"default":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},

		"indexDocumentList": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(GQLIndexDocumentInputType)),
		},

		"indexDynamic":          &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"storeDynamic":          &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"docValuesDynamic":      &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"defaultType":           &graphql.InputObjectFieldConfig{Type: graphql.String},
		"defaultAnalyzer":       &graphql.InputObjectFieldConfig{Type: graphql.String},
		"defaultDateTimeParser": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"defaultField":          &graphql.InputObjectFieldConfig{Type: graphql.String},
		"typeField":             &graphql.InputObjectFieldConfig{Type: graphql.String},
		"customAnalysis":        &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})
