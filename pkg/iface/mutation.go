package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"time"
)

type IMutation interface {
	Mutation(clusterID uint64, pp *pb.FlameProposal, timeout time.Duration) error
}
