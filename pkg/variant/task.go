package variant

type Task struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Command string      `json:"command"`
	Payload interface{} `json:"payload"`
}

type TaskQueue chan Task
