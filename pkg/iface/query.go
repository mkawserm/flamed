package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"time"
)

type IQuery interface {
	Query(clusterID uint64, query *pb.FlameQuery, timeout time.Duration) (*pb.FlameResponse, error)
}
