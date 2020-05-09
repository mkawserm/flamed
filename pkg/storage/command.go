package storage

type CommandID uint8

const (
	SyncFullIndex CommandID = iota
	SyncUpdateIndex
	SyncRunGC
)

type Command struct {
	CommandID CommandID   `json:"commandID"`
	Data      interface{} `json:"data"`
}
