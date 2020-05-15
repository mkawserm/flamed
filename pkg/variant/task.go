package variant

type Task struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Command string      `json:"command"`
	Options interface{} `json:"options"`
}

type TaskQueue chan Task
