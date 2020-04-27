package badger

type SnapshotConfiguration struct {
	GoroutineNumber int    `json:"goroutineNumber"`
	LogPrefix       string `json:"logPrefix"`
}
