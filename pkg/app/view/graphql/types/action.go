package types

import "github.com/graphql-go/graphql"

var ActionEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Action",
	Description: "`Action` defines different types of action",
	Values: graphql.EnumValueConfigMap{

		"GET": &graphql.EnumValueConfig{
			Value:       0,
			Description: "GET",
		},

		"SEARCH": &graphql.EnumValueConfig{
			Value:       1,
			Description: "SEARCH",
		},

		"ITERATE": &graphql.EnumValueConfig{
			Value:       2,
			Description: "ITERATE",
		},

		"MERGE": &graphql.EnumValueConfig{
			Value:       3,
			Description: "MERGE",
		},

		"INSERT": &graphql.EnumValueConfig{
			Value:       4,
			Description: "INSERT",
		},

		"UPDATE": &graphql.EnumValueConfig{
			Value:       5,
			Description: "UPDATE",
		},

		"UPSERT": &graphql.EnumValueConfig{
			Value:       6,
			Description: "UPSERT",
		},

		"DELETE": &graphql.EnumValueConfig{
			Value:       7,
			Description: "DELETE",
		},

		"DEFAULT": &graphql.EnumValueConfig{
			Value:       8,
			Description: "DEFAULT",
		},
	},
})
