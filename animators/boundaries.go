package animators

import (
	"fmt"
	"io"

	"github.com/influx6/govfx"
)

//==============================================================================

// Width provides animation sequencing for width properties, it uses flat integers
// values and pixels.
type Width struct {
	Target  int            `govfx:"value"`
	Easing  string         `govfx:"easing"`
	Elem    *govfx.Element `govfx:"elem"`
	current float64
	accum   float64
}

// Update contains the update operations for the width property.
// All calculations are handled here, it recieves the delta value to
// allow
func (w *Width) Update(delta float64) {
	w.accum += delta
}

// Blend takes the last value to which allows us to correct the
// rendered position of our update.
func (w *Width) Blend(delta float64) {
	w.current += (w.current * delta * w.accum) + 5
	fmt.Println("Blending")
}

// CSS writes the css output to the supplied writer
func (w *Width) CSS(wc io.Writer) {
	wc.Write([]byte(fmt.Sprintf("width: %d%s", int(w.current), "px")))
}

//==============================================================================

// Height provides animation sequencing for Height properties.
type Height struct {
	Target  int            `govfx:"value"`
	Easing  string         `govfx:"easing"`
	Elem    *govfx.Element `govfx:"elem"`
	current int
}

// Update contains the update operations for the width property.
// All calculations are handled here, it recieves the delta value to
// allow
func (w *Height) Update(delta float64) {
	w.current += int(delta * 5)
}

// CSS writes the css output to the supplied writer
func (w *Height) CSS(wc io.Writer) {
	wc.Write([]byte(fmt.Sprintf("width: %d%s", w.current, "px")))
}

//==============================================================================
