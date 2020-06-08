package types

import "github.com/graphql-go/graphql"

var GQLStateMachineType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "StateMachineType",
	Description: "`StateMachineType` defines different types of state machine",
	Values: graphql.EnumValueConfigMap{
		"RegularStateMachine": &graphql.EnumValueConfig{
			Value:       1,
			Description: "Regular State Machine",
		},

		"ConcurrentStateMachine": &graphql.EnumValueConfig{
			Value:       2,
			Description: "Concurrent State Machine",
		},

		"OnDiskStateMachine": &graphql.EnumValueConfig{
			Value:       3,
			Description: "On Disk State Machine",
		},
	},
})
