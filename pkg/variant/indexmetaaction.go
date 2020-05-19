package variant

import "github.com/mkawserm/flamed/pkg/pb"

type IndexMetaAction struct {
	Action    pb.Action     `json:"action"`
	IndexMeta *pb.IndexMeta `json:"indexMeta"`
}
