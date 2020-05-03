package variant

type IndexData struct {
	ID        string      `json:"id"`
	Namespace string      `json:"namespace"`
	Data      interface{} `json:"data"`
}
