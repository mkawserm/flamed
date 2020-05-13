package variant

type IndexData struct {
	ID     string      `json:"id"`
	Data   interface{} `json:"data"`
	Action int8        `json:"action"` /*1 - UPSERT, 2 - DELETE*/
}
