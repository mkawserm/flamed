package nodeadminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var RunStorageGC = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Run storage garbage collector",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
		if !ok {
			return false, nil
		}
		nodeAdmin.RunStorageGC()
		return true, nil
	},
}
