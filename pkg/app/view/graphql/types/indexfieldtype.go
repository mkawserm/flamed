package types

import "github.com/graphql-go/graphql"

var GQLIndexFieldTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "IndexFieldType",
	Description: "`IndexFieldType` defines enum for `IndexField`",
	Values: graphql.EnumValueConfigMap{
		"TEXT": &graphql.EnumValueConfig{
			Value:       0,
			Description: "TEXT",
		},

		"NUMERIC": &graphql.EnumValueConfig{
			Value:       1,
			Description: "NUMERIC",
		},

		"BOOLEAN": &graphql.EnumValueConfig{
			Value:       2,
			Description: "BOOLEAN",
		},

		"GEO_POINT": &graphql.EnumValueConfig{
			Value:       3,
			Description: "GEO_POINT",
		},

		"DATE_TIME": &graphql.EnumValueConfig{
			Value:       4,
			Description: "DATE_TIME",
		},
	},
})
