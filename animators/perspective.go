package animators

import (
	"fmt"

	"github.com/influx6/govfx"
)

//==============================================================================

// Perspective defines a sequence for animating css perspective x-axes properties.
type Perspective struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t *Perspective) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "perspective")

		func(e govfx.Elemental) {

			var r float64

			if govfx.IsPerspective(transform) {
				mx, _ := govfx.ToPerspective(transform)
				r = mx.Range
			} else {
				r = 0
			}

			transform = fmt.Sprintf("Perspective(%.0fpx)", r)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

// Next returns the writers for the next sequence.
func (t *Perspective) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "Perspective")

		func(e govfx.Elemental) {

			var r float64

			if govfx.IsPerspective(transform) {
				mx, _ := govfx.ToPerspective(transform)
				r = mx.Range
			} else {
				r = 0
			}

			rd := float64(t.Value) - r

			rn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: r,
				DeltaValue:   rd,
			})

			e.EraseMore("transform", "matrix", false)

			transform = fmt.Sprintf("Perspective(%.0fpx)", rn)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================
