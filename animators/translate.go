package animators

import (
	"fmt"

	"github.com/influx6/faux/utils"
	"github.com/influx6/govfx"
)

//==============================================================================

// TranslateY defines a sequence for animating css translate y-axes properties.
type TranslateY struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t TranslateY) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "translate")
		position, pr, _ := elem.Read("position", "")

		if utils.MatchAny(position, "none", "") {
			position = "relative"
		}

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsTranslation(transform) {
				mx, _ := govfx.ToTranslation(transform)
				x, y = mx.X, mx.Y
			} else {
				// x, y = govfx.Position(e)
				x, y = 0, 0
			}

			transform = fmt.Sprintf("translate(%.0fpx, %.0fpx)", x, y)

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
func (t TranslateY) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "translate")

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsTranslation(transform) {
				mx, _ := govfx.ToTranslation(transform)
				x, y = mx.X, mx.Y
			}

			yd := float64(t.Value) - y

			yn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: y,
				DeltaValue:   yd,
			})

			transform = fmt.Sprintf("translate(%.0fpx, %.0fpx)", x, yn)
			e.EraseMore("transform", "matrix", false)
			writers = append(writers, govfx.NewWriter(func() {
				e.WriteMore("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================

// TranslateX defines a sequence for animating css translate x-axes properties.
type TranslateX struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t TranslateX) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "translate")
		position, pr, _ := elem.Read("position", "")

		if utils.MatchAny(position, "none", "") {
			position = "relative"
		}

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionX
			} else if govfx.IsTranslation(transform) {
				mx, _ := govfx.ToTranslation(transform)
				x, y = mx.X, mx.X
			} else {
				// x, y = govfx.Position(e)
				x, y = 0, 0
			}

			transform = fmt.Sprintf("translate(%.0fpx, %.0fpx)", x, y)

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
func (t TranslateX) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "translate")

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionX
			} else if govfx.IsTranslation(transform) {
				mx, _ := govfx.ToTranslation(transform)
				x, y = mx.X, mx.X
			}

			xd := float64(t.Value) - x

			xn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: x,
				DeltaValue:   xd,
			})

			transform = fmt.Sprintf("translate(%.0fpx, %.0fpx)", xn, y)
			e.EraseMore("transform", "matrix", false)
			writers = append(writers, govfx.NewWriter(func() {
				e.WriteMore("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================
