package adminmutator

import (
	"github.com/graphql-go/graphql"
)

var GQLAdminMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AdminMutator",
	Description: "`AdminMutator`",
	Fields: graphql.Fields{
		"upsertAccessControl": UpsertAccessControl,
		"deleteAccessControl": DeleteIndexMeta,
		"deleteUser":          DeleteUser,
		"upsertUser":          UpsertUser,
		"updateUser":          UpdateUser,
		"changeUserPassword":  ChangeUserPassword,

		"upsertIndexMeta": UpsertIndexMeta,
		"deleteIndexMeta": DeleteIndexMeta,
	},
})
