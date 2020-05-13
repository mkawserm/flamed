package iface

import (
	sm "github.com/lni/dragonboat/v3/statemachine"
	"github.com/mkawserm/flamed/pkg/pb"
	"time"
)

type IReader interface {
	Read(clusterID uint64,
		query interface{},
		timeout time.Duration) (interface{}, error)
}

type IWriter interface {
	Write(clusterID uint64,
		pp *pb.Proposal,
		timeout time.Duration) (sm.Result, error)
}

type IRW interface {
	IReader
	IWriter
}
