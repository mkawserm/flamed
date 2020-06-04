package iface

import "time"

type ITerm interface {
	Term() string
	Count() int
}

type IDateRangeFacet interface {
	Name() string
	Start() string
	End() string
	Count() int
}

type INumericRangeFacet interface {
	Name() string
	Min() float64
	Max() float64
	Count() int
}

type IFacet interface {
	Name() string
	Field() string
	Total() int
	Missing() int
	Other() int
	Terms() []ITerm
	NumericRanges() []INumericRangeFacet
	DateRanges() []IDateRangeFacet
}

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
	Facets() []IFacet
}
