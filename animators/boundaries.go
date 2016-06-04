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
	current float64
	accum   float64
	elem    *govfx.Element
	Target  int    `govfx:"value"`
	Easing  string `govfx:"easing"`
}

// Init initializes the width property with the provided element for animation.
func (w *Width) Init(elem *govfx.Element) {
	w.elem = elem
	if ws, ok, _ := elem.ReadInt("width", ""); ok {
		w.current = float64(ws)
	}
}

// Update contains the update operations for the width property.
// All calculations are handled here, it recieves the delta value to
// allow
func (w *Width) Update(delta float64) {
	w.current += (w.current * delta) + 5
}

// Blend takes the last value to which allows us to correct the
// rendered position of our update.
func (w *Width) Blend(delta float64) {
	w.current += delta
}

// CSS writes the css output to the supplied writer
func (w *Width) CSS(wc io.Writer) {
	wc.Write([]byte(fmt.Sprintf("width: %d%s", int(w.current), "px")))
}

//==============================================================================

// Height provides animation sequencing for Height properties.
type Height struct {
	current float64
	elem    *govfx.Element
	Target  int    `govfx:"value"`
	Easing  string `govfx:"easing"`
}

// Init initializes the width property with the provided element for animation.
func (h *Height) Init(elem *govfx.Element) {
	h.elem = elem
	if ws, ok, _ := elem.ReadInt("height", ""); ok {
		h.current = float64(ws)
	}
}

// Update contains the update operations for the width property.
// All calculations are handled here, it recieves the delta value to
// allow
func (h *Height) Update(delta float64) {
	h.current += (delta * 5)
}

// CSS writes the css output to the supplied writer
func (h *Height) CSS(wc io.Writer) {
	wc.Write([]byte(fmt.Sprintf("height: %d%s", int(h.current), "px")))
}

//==============================================================================
