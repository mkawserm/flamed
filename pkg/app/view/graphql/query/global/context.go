package global

import (
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/pb"
)

type Context struct {
	Query         *flamed.Query
	AccessControl *pb.AccessControl
}
