package graphql

import (
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation/intkeytpmutator"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation/jsontpmutator"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query/global"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query/intkeytp"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query/jsontp"
)

func (v *View) register() {
	v.AddQueryField("global", global.Global)
	v.AddQueryField("serviceStatus", query.ServiceStatus)

	v.AddMutationField("counterMutator", mutation.CounterMutator)

	v.AddQueryField("flamed", query.Flamed)
	v.AddMutationField("flamedMutator", mutation.FlamedMutator)

	v.AddQueryField("intKeyTP", intkeytp.IntKeyTP)
	v.AddMutationField("intKeyTPMutator", intkeytpmutator.IntKeyTPMutator)

	v.AddQueryField("jsonTP", jsontp.JSONTP)
	v.AddMutationField("jsonTPMutator", jsontpmutator.JSONTPMutator)
}
