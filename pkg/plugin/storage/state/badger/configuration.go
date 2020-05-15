package badger

import (
	badgerDb "github.com/dgraph-io/badger/v2"
)

type Configuration struct {
	SliceCap        int              `json:"sliceCap"`
	LogPrefix       string           `json:"logPrefix"`
	GoroutineNumber int              `json:"goroutineNumber"`
	BadgerOptions   badgerDb.Options `json:"badgerOptions"`
	GCDiscardRatio  float64          `json:"gcDiscardRatio"`
}
