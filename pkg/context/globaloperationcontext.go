package context

import (
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/pb"
)

type GlobalOperationContext struct {
	GlobalOperation *flamed.GlobalOperation
	AccessControl   *pb.AccessControl
}
