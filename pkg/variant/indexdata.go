package variant

const CREATE = 1
const UPDATE = 2
const DELETE = 3

type IndexData struct {
	ID     string      `json:"id"`
	Data   interface{} `json:"data"`
	Action int8        `json:"action"` /* 1 - CREATE, 2 - UPDATE, 3 - DELETE */
}
