package context

import (
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/pb"
)

type GlobalOperationContext struct {
	AccessControl   *pb.AccessControl
	GlobalOperation *flamed.GlobalOperation
}
