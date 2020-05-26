package graphql

import (
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation/intkeymutator"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query/intkey"
)

func (v *View) register() {
	v.AddQueryField("flamed", query.Flamed)
	v.AddQueryField("serviceStatus", query.ServiceStatus)

	v.AddQueryField("intKey", intkey.IntKey)

	v.AddMutationField("flamedMutator", mutation.FlamedMutator)
	v.AddMutationField("counterMutator", mutation.CounterMutator)
	v.AddMutationField("intKeyMutator", intkeymutator.IntKeyMutator)
}
