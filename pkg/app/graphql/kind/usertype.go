package kind

import "github.com/graphql-go/graphql"

var GQLUserTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserType",
	Description: "`UserType` can be either SUPER_USER or NORMAL_USER",
	Values: graphql.EnumValueConfigMap{
		"SUPER_USER": &graphql.EnumValueConfig{
			Value:       0,
			Description: "SUPER_USER",
		},
		"NORMAL_USER": &graphql.EnumValueConfig{
			Value:       1,
			Description: "NORMAL_USER",
		},
	},
})
