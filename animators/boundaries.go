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
	Target int          `govfx:"value"`
	Easing string       `govfx:"easing"`
	Easer  govfx.Easing `govfx:"Easer"`

	current float64
	accum   float64

	elem govfx.Elemental

	ended bool
}

// Init initializes the width property with the provided element for animation.
func (w *Width) Init(elem govfx.Elemental) {
	w.elem = elem

	if w.Easer == nil {
		w.Easer = govfx.GetEasing(w.Easing)
	}

	if ws, _, ok := elem.ReadInt("width", ""); ok {
		w.current = float64(ws)
	}
}

// Update contains the update operations for the width property.
// All calculations are handled here, it recieves the delta value to
// allow
func (w *Width) Update(delta float64, timeline float64) {
	cu := int(w.current)

	easer := w.Easer.Ease(timeline)

	if cu < w.Target {
		w.current += (w.current * delta * easer) + 5
	} else {
		w.current = float64(w.Target)
		w.ended = true
	}
}

// Blend takes the last value to which allows us to correct the
// rendered position of our update.
func (w *Width) Blend(delta float64) {
	if !w.ended {
		w.current += delta
	}
}

// CSS writes the css output to the supplied writer
func (w *Width) CSS(wc io.Writer) {
	wc.Write([]byte(fmt.Sprintf("width: %d%s", int(w.current), "px")))
}

//==============================================================================

// Height provides animation sequencing for Height properties.
type Height struct {
	Target int          `govfx:"value"`
	Easing string       `govfx:"easing"`
	Easer  govfx.Easing `govfx:"easer"`

	current float64
	accum   float64

	elem govfx.Elemental

	ended bool
}

// Init initializes the width property with the provided element for animation.
func (h *Height) Init(elem govfx.Elemental) {
	h.elem = elem

	if h.Easer == nil {
		h.Easer = govfx.GetEasing(h.Easing)
	}

	if ws, _, ok := elem.ReadInt("width", ""); ok {
		h.current = float64(ws)
	}
}

// Update contains the update operations for the width property.
// All calculations are handled here, it recieves the delta value to
// allow
func (h *Height) Update(delta float64, timeline float64) {
	cu := int(h.current)

	easer := h.Easer.Ease(timeline)

	if cu < h.Target {
		h.current += (h.current * delta * easer) + 5
	} else {
		h.current = float64(h.Target)
		h.ended = true
	}
}

// Blend takes the last value to which allows us to correct the
// rendered position of our update.
func (h *Height) Blend(delta float64) {
	if !h.ended {
		h.current += delta
	}
}

// CSS writes the css output to the supplied writer
func (h *Height) CSS(wc io.Writer) {
	wc.Write([]byte(fmt.Sprintf("height: %d%s", int(h.current), "px")))
}

//==============================================================================
