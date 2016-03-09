package animators

import (
	"fmt"

	"github.com/influx6/faux/utils"
	"github.com/influx6/govfx"
)

//==============================================================================

// SkewY defines a sequence for animating css Skew y-axes properties.
type SkewY struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t SkewY) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "skew")
		position, pr, _ := elem.Read("position", "")

		if utils.MatchAny(position, "none", "") {
			position = "relative"
		}

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsSkew(transform) {
				mx, _ := govfx.ToSkew(transform)
				x, y = mx.X, mx.Y
			} else {
				x, y = 0, 0
			}

			transform = fmt.Sprintf("skew(%.0fpx, %.0fpx)", x, y)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Write("position", position, pr)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

// Next returns the writers for the next sequence.
func (t SkewY) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "skew")

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsSkew(transform) {
				mx, _ := govfx.ToSkew(transform)
				x, y = mx.X, mx.Y
			}

			yd := float64(t.Value) - y

			yn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: y,
				DeltaValue:   yd,
			})

			transform = fmt.Sprintf("skew(%.0fpx, %.0fpx)", x, yn)
			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================

// SkewX defines a sequence for animating css Skew x-axes properties.
type SkewX struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t SkewX) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "skew")
		position, pr, _ := elem.Read("position", "")

		if utils.MatchAny(position, "none", "") {
			position = "relative"
		}

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsSkew(transform) {
				mx, _ := govfx.ToSkew(transform)
				x, y = mx.X, mx.Y
			} else {
				x, y = 0, 0
			}

			transform = fmt.Sprintf("skew(%.0fpx, %.0fpx)", x, y)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Write("position", position, pr)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

// Next returns the writers for the next sequence.
func (t SkewX) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "skew")

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsSkew(transform) {
				mx, _ := govfx.ToSkew(transform)
				x, y = mx.X, mx.Y
			}

			xd := float64(t.Value) - x

			xn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: x,
				DeltaValue:   xd,
			})

			transform = fmt.Sprintf("skew(%.0fpx, %.0fpx)", xn, y)
			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================
