package animators

import (
	"fmt"

	"github.com/influx6/govfx"
)

//==============================================================================

// RotateY defines a sequence for animating css Rotate y-axes properties.
type RotateY struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t *RotateY) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "rotate")

		func(e govfx.Elemental) {

			var y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				y = mx.RotationY
			} else if govfx.IsRotation(transform) {
				mx, _ := govfx.ToRotation(transform)
				y = mx.Angle
			} else {
				y = 0
			}

			transform = fmt.Sprintf("rotateY(%.0fdeg)", y)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
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

			var y float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				y = mx.RotationY
			} else if govfx.IsRotation(transform) {
				mx, _ := govfx.ToRotation(transform)
				y = mx.Angle
			}

			yd := float64(t.Value) - y

			yn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: y,
				DeltaValue:   yd,
			})

			e.EraseMore("transform", "matrix", false)

			transform = fmt.Sprintf("rotateY(%.0fdeg)", yn)

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
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t *RotateX) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "rotate")

		func(e govfx.Elemental) {

			var x float64

			if govfx.IsMatrix(transform) {
				mx, _ := govfx.ToMatrix2D(transform)
				x = mx.RotationX
			} else if govfx.IsRotation(transform) {
				mx, _ := govfx.ToRotation(transform)
				x = mx.Angle
			} else {
				x = 0
			}

			transform = fmt.Sprintf("rotateX(%.0fdeg)", x)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
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
				x = mx.RotationX
			} else if govfx.IsRotation(transform) {
				mx, _ := govfx.ToRotation(transform)
				x = mx.Angle
			}

			xd := float64(t.Value) - x

			xn := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: x,
				DeltaValue:   xd,
			})

			e.EraseMore("transform", "matrix", false)
			transform = fmt.Sprintf("rotateX(%.0fdeg)", xn)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================

// Rotate defines a sequence for animating css Rotate x-axes properties.
type Rotate struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (t *Rotate) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "rotate")

		func(e govfx.Elemental) {

			var r float64

			if govfx.IsSimpleRotation(transform) {
				mx, _ := govfx.ToRotation(transform)
				r = mx.Angle
			} else {
				r = 0
			}

			transform = fmt.Sprintf("rotate(%.0fdeg)", r)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

// Next returns the writers for the next sequence.
func (t *Rotate) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(t.Easing)

	for _, elem := range elems {
		transform, priority, _ := elem.Read("transform", "rotate")

		func(e govfx.Elemental) {

			var r float64

			if govfx.IsSimpleRotation(transform) {
				mx, _ := govfx.ToRotation(transform)
				r = mx.Angle
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

			transform = fmt.Sprintf("rotate(%.0fdeg)", rn)

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("transform", transform, priority)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

//==============================================================================
