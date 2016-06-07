package animators

import (
	"fmt"
	"io"

	"github.com/influx6/govfx"
)

//==============================================================================

// Colors defines a global handle for color transisition.
var Colors ColorTransistion

// ColorTransistion defines a struct for defining color functions.
type ColorTransistion struct{}

// Interpolate interpolates the current color towards the base color using the
// delta and timeline values and the easing provider.
func (ColorTransistion) Interpolate(easing govfx.Easing, base, current ColorValue, delta float64, timeline float64) ColorValue {
	var newcolor ColorValue

	return newcolor
}

// Blend returns a new ColorValue with the giving blend function.
func (ColorTransistion) Blend(current ColorValue, blend float64) ColorValue {
	var newcolor ColorValue

	return newcolor
}

//==============================================================================

// ColorValue defines a struct which represents a color value set.
type ColorValue struct {
	red   int
	green int
	blue  int
	alpah float64
}

// RGBA writes out the color values in RGBA format.
func (c ColorValue) RGBA() string {
	return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", c.red, c.green, c.blue, c.alpah)
}

// RGB writes out the color values in RGB format.
func (c ColorValue) RGB() string {
	return fmt.Sprintf("rgba(%d,%d,%d,1)", c.red, c.green, c.blue)
}

// Color provides a animator for sequencing color animations.
type Color struct {
	Alpha  bool   `govfx:"alpha"`
	Color  string `govfx:"color"`
	Easing string `govfx:"easing"`

	color ColorValue
	base  ColorValue

	elem  govfx.Elemental
	easer govfx.Easing
}

// Init initializes the property for execution.
func (t *Color) Init(elem govfx.Elemental) {
	t.elem = elem

	rw, gw, bw, aw := govfx.ParseRGB(t.Color)
	t.color = ColorValue{red: rw, green: gw, blue: bw, alpah: aw}

	if color, _, ok := elem.Read("color", ""); ok {
		r, g, b, a := govfx.ParseRGB(color)
		t.color = ColorValue{red: r, green: g, blue: b, alpah: a}
	}
}

// Update updates the property details.
func (t *Color) Update(delta float64, timeline float64) {
	easing := govfx.GetEasing(t.Easing)
	t.color = Colors.Interpolate(easing, t.base, t.color, delta, timeline)
}

// Blend adjust the property details to match appropriate state.
func (t *Color) Blend(interpolation float64) {
	t.color = Colors.Blend(t.color, interpolation)
}

// CSS writes out the current state of the property in css format to the provided
// writer.
func (t *Color) CSS(owner io.Writer) {
	if t.Alpha {
		owner.Write([]byte("color:" + t.color.RGBA()))
		return
	}

	owner.Write([]byte("color:" + t.color.RGB()))
}

//==============================================================================

// BackgroundColor provides a animator for sequencing background color animations.
type BackgroundColor struct {
	Alpha  bool   `govfx:"alpha"`
	Color  string `govfx:"color"`
	Easing string `govfx:"easing"`

	color ColorValue
	base  ColorValue

	elem  govfx.Elemental
	easer govfx.Easing
}

// Update updates the property details.
func (t *BackgroundColor) Update(delta float64, timeline float64) {
	easing := govfx.GetEasing(t.Easing)
	t.color = Colors.Interpolate(easing, t.base, t.color, delta, timeline)
}

// Blend adjust the property details to match appropriate state.
func (t *BackgroundColor) Blend(interpolation float64) {
	t.color = Colors.Blend(t.color, interpolation)
}

// CSS writes out the current state of the property in css format to the provided
// writer.
func (t *BackgroundColor) CSS(owner io.Writer) {
	if t.Alpha {
		owner.Write([]byte("background-color: " + t.color.RGBA()))
		return
	}

	owner.Write([]byte("background-color: " + t.color.RGB()))
}
