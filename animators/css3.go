package animators

import (
	"fmt"
	"regexp"

	"github.com/influx6/govfx"
)

//==============================================================================

var tranlateMatch = regexp.MustCompile("translate\\(([\\d,\\s]+)\\)")
var matrixMatch = regexp.MustCompile("matrix\\(([,\\d\\s]+)\\)")

// TranslateY defines a sequence for animating css translate y-axes properties.
type TranslateY struct {
	Value int `govfx:"value"`
}

// Init returns the initial writers for the sequence.
func (t TranslateY) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform")

		func(e govfx.Elemental) {
			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

// Next returns the writers for the next sequence.
func (t TranslateY) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, _, _ := elem.Read("transform")

		// var x, y int

		if matrixMatch.MatchString(transform) {
			mx := matrixMatch.FindStringSubmatch(transform)[1]
			fmt.Printf("Current matrix: %s\n", mx)
		}

		fmt.Printf("Value matrix: %s\n", transform)
		// func(e govfx.Elemental) {
		// 	writers = append(writers, govfx.NewWriter(func() {
		// 		mx := fmt.Sprintf("transform")
		// 		e.Write("transform", transform, priority)
		// 		e.Sync()
		// 	}))
		// }(elem)
	}

	return writers
}

//==============================================================================

// TranslateX defines a sequence for animating css translate y-axes properties.
type TranslateX struct {
	Value int
}

//==============================================================================

// Bounce provides a sequence to animating a bounce effect.
type Bounce struct {
}

//==============================================================================
