package graphql

import (
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query"
)

func (v *View) register() {
	v.AddQueryField("flamed", query.Flamed)
	v.AddQueryField("serviceStatus", query.ServiceStatus)

	v.AddMutationField("flamedMutator", mutation.FlamedMutator)

	v.AddMutationField("increment", mutation.Increment)
	v.AddMutationField("decrement", mutation.Decrement)
}
