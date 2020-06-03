package iface

import "time"

type IDocument interface {
	ID() string
	Score() float64
	Index() string
}

type ISearchResponse interface {
	Total() uint64
	MaxScore() float64
	Took() time.Duration
	Hits() []IDocument
}
