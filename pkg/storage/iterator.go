package storage

import "github.com/mkawserm/flamed/pkg/pb"

type Iterator struct {
	Seek     []byte
	Prefix   []byte
	Limit    int
	Receiver func(entry *pb.FlameEntry) bool
}
