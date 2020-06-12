package kind

import (
	"github.com/graphql-go/graphql"
)

var GQLIndexFieldInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "IndexFieldInput",
	Description: "`IndexFieldInput`",
	Fields: graphql.InputObjectConfigFieldMap{
		"indexFieldType": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(GQLIndexFieldTypeEnum)},
		"name":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"analyzer":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"enabled":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},

		"index":              &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
		"store":              &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
		"includeTermVectors": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
		"includeInAll":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
		"docValues":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
	},
})
