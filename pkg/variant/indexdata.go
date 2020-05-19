package variant

import "github.com/mkawserm/flamed/pkg/pb"

type IndexData struct {
	ID     string      `json:"id"`
	Data   interface{} `json:"data"`
	Action pb.Action   `json:"action"`
}
