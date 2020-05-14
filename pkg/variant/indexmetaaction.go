package variant

import "github.com/mkawserm/flamed/pkg/pb"

type IndexMetaAction struct {
	Action    int8          `json:"action"` /*1 - UPSERT, 2 - DELETE, 3 - DEFAULT*/
	IndexMeta *pb.IndexMeta `json:"indexMeta"`
}
