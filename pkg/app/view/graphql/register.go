package graphql

import "github.com/mkawserm/flamed/pkg/app/view/graphql/query"

func (v *View) register() {
	v.AddQueryField("isLive", query.IsLive)
}
