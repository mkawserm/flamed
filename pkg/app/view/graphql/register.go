package graphql

import (
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation/intkeytpmutator"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query/intkey"
)

func (v *View) register() {
	v.AddQueryField("serviceStatus", query.ServiceStatus)

	v.AddMutationField("counterMutator", mutation.CounterMutator)

	v.AddQueryField("flamed", query.Flamed)
	v.AddMutationField("flamedMutator", mutation.FlamedMutator)

	v.AddQueryField("intKeyTP", intkey.IntKeyTP)
	v.AddMutationField("intKeyTPMutator", intkeytpmutator.IntKeyTPMutator)
}
