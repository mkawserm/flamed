package storage

type CommandID uint8

const (
	SyncCompleteIndexUpdate CommandID = iota
	SyncPartialIndexUpdate
	SyncRunGC
)

type Command struct {
	CommandID CommandID   `json:"commandID"`
	Data      interface{} `json:"data"`
}
