package globaloperation

import (
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/pb"
)

type Context struct {
	GlobalOperation *flamed.GlobalOperation
	AccessControl   *pb.AccessControl
}
