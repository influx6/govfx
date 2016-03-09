package animators

import (
	"fmt"

	"github.com/influx6/faux/utils"
	"github.com/influx6/govfx"
)

//==============================================================================

// ScaleY defines a sequence for animating css Scale y-axes properties.
type ScaleY struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t ScaleY) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "scale")
		position, pr, _ := elem.Read("position", "")

		if utils.MatchAny(position, "none", "") {
			position = "relative"
		}

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsScale(transform) {
				mx, _ := govfx.ToScale(transform)
				x, y = mx.X, mx.Y
			} else {
				x, y = 0, 0
			}

			transform = fmt.Sprintf("scale(%.0fpx, %.0fpx)", x, y)

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
func (t ScaleY) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "scale")

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.ScaleX, mx.ScaleY
			} else if govfx.IsScale(transform) {
				mx, _ := govfx.ToScale(transform)
				x, y = mx.X, mx.Y
			}

			yd := float64(t.Value) - y

			yn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: y,
				DeltaValue:   yd,
			})

			transform = fmt.Sprintf("scale(%.0fpx, %.0fpx)", x, yn)
			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================

// ScaleX defines a sequence for animating css Scale x-axes properties.
type ScaleX struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t ScaleX) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "scale")
		position, pr, _ := elem.Read("position", "")

		if utils.MatchAny(position, "none", "") {
			position = "relative"
		}

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsScale(transform) {
				mx, _ := govfx.ToScale(transform)
				x, y = mx.X, mx.Y
			} else {
				x, y = 0, 0
			}

			transform = fmt.Sprintf("scale(%.0fpx, %.0fpx)", x, y)

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
func (t ScaleX) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "scale")

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.ScaleX, mx.ScaleY
			} else if govfx.IsScale(transform) {
				mx, _ := govfx.ToScale(transform)
				x, y = mx.X, mx.Y
			}

			xd := float64(t.Value) - x

			xn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: x,
				DeltaValue:   xd,
			})

			transform = fmt.Sprintf("scale(%.0fpx, %.0fpx)", xn, y)
			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================
