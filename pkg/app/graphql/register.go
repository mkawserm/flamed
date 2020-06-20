package graphql

import (
	"github.com/mkawserm/flamed/pkg/app/graphql/mutation"
	"github.com/mkawserm/flamed/pkg/app/graphql/mutation/globaloperationmutator"
	"github.com/mkawserm/flamed/pkg/app/graphql/mutation/intkeytpmutator"
	"github.com/mkawserm/flamed/pkg/app/graphql/mutation/jsontpmutator"
	"github.com/mkawserm/flamed/pkg/app/graphql/query"
	"github.com/mkawserm/flamed/pkg/app/graphql/query/globaloperation"
	"github.com/mkawserm/flamed/pkg/app/graphql/query/intkeytp"
	"github.com/mkawserm/flamed/pkg/app/graphql/query/jsontp"
)

func (g *GraphQL) register() {
	g.AddQueryField("serviceStatus", query.ServerStatus)
	g.AddMutationField("counterMutator", mutation.CounterMutator)

	g.AddQueryField("flamed", query.Flamed)
	g.AddMutationField("flamedMutator", mutation.FlamedMutator)

	g.AddQueryField("intKeyTP", intkeytp.IntKeyTP)
	g.AddMutationField("intKeyTPMutator", intkeytpmutator.IntKeyTPMutator)

	g.AddQueryField("jsonTP", jsontp.JSONTP)
	g.AddMutationField("jsonTPMutator", jsontpmutator.JSONTPMutator)

	g.AddQueryField("globalOperation", globaloperation.GlobalOperation)
	g.AddMutationField("globalOperationMutator", globaloperationmutator.GlobalOperationMutator)
}
