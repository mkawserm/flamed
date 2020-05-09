package storage

import "github.com/mkawserm/flamed/pkg/pb"

type Iterator struct {
	Seek     []byte                          `json:"seek"`
	Prefix   []byte                          `json:"prefix"`
	Limit    int                             `json:"limit"`
	Receiver func(entry *pb.FlameEntry) bool `json:"-"`
}
