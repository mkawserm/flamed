package nodeadminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var BuildIndex = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Build or rebuild full state database index",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
		if !ok {
			return false, nil
		}
		nodeAdmin.BuildIndex()
		return true, nil
	},
}
