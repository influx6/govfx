package animators

import (
	"fmt"

	"github.com/influx6/faux/utils"
	"github.com/influx6/govfx"
)

//==============================================================================

// RotateY defines a sequence for animating css Rotate y-axes properties.
type RotateY struct {
	Value  float64 `govfx:"value"`
	Easing string  `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t *RotateY) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "rotate")
		position, pr, _ := elem.Read("position", "")

		if utils.MatchAny(position, "none", "") {
			position = "relative"
		}

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsRotation(transform) {
				mx, _ := govfx.ToRotation(transform)
				x, y = mx.X, mx.Y
			} else {
				x, y = 0, 0
			}

			transform = fmt.Sprintf("rotate(%.0fdeg, %.0fdeg)", x, y)

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
func (t *RotateY) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "rotate")

		func(e govfx.Elemental) {

			var x, y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x, y = mx.PositionX, mx.PositionY
			} else if govfx.IsRotate(transform) {
				mx, _ := govfx.ToRotate(transform)
				x, y = mx.X, mx.Y
			}

			yd := float64(t.Value) - y

			yn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: y,
				DeltaValue:   yd,
			})

			transform = fmt.Sprintf("rotate(%.0fdeg, %.0fdeg)", x, yn)
			e.EraseMore("transform", "matrix", false)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================

// RotateX defines a sequence for animating css Rotate x-axes properties.
type RotateX struct {
	Value  float64 `govfx:"value"`
	Easing string  `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t *RotateX) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "rotate")
		position, pr, _ := elem.Read("position", "")

		if utils.MatchAny(position, "none", "") {
			position = "relative"
		}

		func(e govfx.Elemental) {

			var x float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x = mx.PositionX
			} else if govfx.IsRotate(transform) {
				mx, _ := govfx.ToRotate(transform)
				x = mx.X
			} else {
				x = 0
			}

			transform = fmt.Sprintf("rotateX(%.0fdeg)", x)

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
func (t *RotateX) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "rotate")

		func(e govfx.Elemental) {

			var x float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x = mx.PositionX
			} else if govfx.IsRotate(transform) {
				mx, _ := govfx.ToRotate(transform)
				x = mx.X
			}

			xd := float64(t.Value) - x

			xn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: x,
				DeltaValue:   xd,
			})

			transform = fmt.Sprintf("rotateX(%.0fdeg)", xn)
			e.EraseMore("transform", "matrix", false)
			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================
