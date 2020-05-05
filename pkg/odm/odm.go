package odm

import "github.com/mkawserm/flamed/pkg/iface"

type ODM struct {
	QueryObject       iface.IQuery
	MutationObject    iface.IMutation
	TransformerObject iface.ITransformer
}
