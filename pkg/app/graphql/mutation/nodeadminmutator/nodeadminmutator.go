package nodeadminmutator

import (
	"github.com/graphql-go/graphql"
)

var GQLNodeAdminMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NodeAdminMutator",
	Description: "`NodeAdminMutator` gives the ability to perform administrative tasks of the cluster",
	Fields: graphql.Fields{
		// AddNode
		"addNode": AddNode,
		// AddObserver
		"addObserver": AddObserver,
		// AddWitness
		"addWitness": AddWitness,
		// DeleteNode
		"deleteNode": DeleteNode,
		//RunStorageGC
		"runStorageGC": RunStorageGC,
		//BuildIndex
		"buildIndex": BuildIndex,
		// BuildIndexByNamespace
		"buildIndexByNamespace": BuildIndexByNamespace,
		// RequestSnapshot
		"requestSnapshot": RequestSnapshot,
	},
})
