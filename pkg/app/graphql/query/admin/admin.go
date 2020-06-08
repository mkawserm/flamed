package admin

import (
	"github.com/graphql-go/graphql"
)

var GQLAdminType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Admin",
	Description: "`Admin` provides user,index meta,access control related information related to the cluster",
	Fields: graphql.Fields{

		// IsUserAvailable
		"isUserAvailable": IsUserAvailable,
		// GetUser
		"getUser": GetUser,
		// IsAccessControlAvailable
		"isAccessControlAvailable": IsAccessControlAvailable,

		"getAccessControl": GetAccessControl,

		// IsIndexMetaAvailable
		"isIndexMetaAvailable": IsIndexMetaAvailable,

		"getIndexMeta": GetIndexMeta,
	},
})
