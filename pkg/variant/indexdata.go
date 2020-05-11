package variant

const SET = 1
const DELETE = 2

type IndexData struct {
	ID     string      `json:"id"`
	Data   interface{} `json:"data"`
	Action int8        `json:"action"` /*1 - SET, 2 - DELETE*/
}
