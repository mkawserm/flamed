package graphql

import (
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/query"
)

func (v *View) register() {
	v.AddQueryField("isLive", query.IsLive)
	v.AddQueryField("nodeAdmin", query.NodeAdmin)

	v.AddMutationField("increment", mutation.Increment)
	v.AddMutationField("decrement", mutation.Decrement)
}
