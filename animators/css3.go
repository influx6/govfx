package animators

import (
	"regexp"

	"github.com/influx6/govfx"
)

//==============================================================================

var tranlateMatch = regexp.MustCompile("translate([\\d,]+)")

// TranslateY defines a sequence for animating css translate y-axes properties.
type TranslateY struct {
	Value int `govfx:"value"`
}

// Init returns the initial writers for the sequence.
func (t TranslateY) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	// for _, elem := range elems {
	// 	transform, priority, _ := elem.Read("transform")
	//
	// 	(func(e govfx.Elemental) {
	// 		writers = append(writers, govfx.NewWriter(func() {
	// 			// val := fmt.Sprintf("%d%s", width, govfx.Unit(w.Unit))
	// 			// e.Write("width", val, priority)
	// 			// e.Sync()
	// 		}))
	// 	})(elem)
	// }

	return writers
}

// TranslateX defines a sequence for animating css translate y-axes properties.
type TranslateX struct {
	Value int
}

//==============================================================================

// Bounce provides a sequence to animating a bounce effect.
type Bounce struct {
}

//==============================================================================
