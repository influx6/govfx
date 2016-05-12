package animators

import (
	"fmt"

	"github.com/influx6/govfx"
)

//==============================================================================

// Width provides animation sequencing for width properties, it uses flat integers
// values and pixels.
type Width struct {
	Value        int    `govfx:"value"`
	Easing       string `govfx:"easing"`
	initialValue int
	newValue     int
}

// Update contains the update operations for the width property.
// All calculations are handled here, it recieves the delta value to
// allow
func (w *Width) Update(delta float64) {
	w.newValue += int(delta * 5)
}

// Init returns the initial writers for the sequence.
func (w *Width) Init(elem govfx.Elemental) govfx.DeferWriter {
	width, priority, _ := elem.ReadInt("width", "")
	w.newValue = width

	return govfx.NewWriter(func() {
		val := fmt.Sprintf("%d%s", width, "px")
		elem.Write("width", val, priority)
	})
}

// Write returns the writers for the current sequence iteration.
func (w *Width) Write(e govfx.Elemental) govfx.DeferWriter {
	m := w.newValue
	return govfx.NewWriter(func() {
		val := fmt.Sprintf("%d%s", m, "px")
		e.Write("width", val, true)
	})
}

//==============================================================================

// Height provides animation sequencing for Height properties.
type Height struct {
	Value        int    `govfx:"value"`
	Easing       string `govfx:"easing"`
	initialValue int
	newValue     int
}

// Init returns the initial writers for the sequence.
func (h *Height) Init(elem govfx.Elemental) govfx.DeferWriter {
	height, priority, _ := elem.ReadInt("height", "")
	h.initialValue = height

	return govfx.NewWriter(func() {
		val := fmt.Sprintf("%d%s", height, "px")
		elem.Write("height", val, priority)
	})
}

// Update contains the update operations for the width property.
// All calculations are handled here, it recieves the delta value to
// allow
func (h *Height) Update(delta float64) {
	h.newValue += int(delta * 5)
}

// Write returns the writers for the current sequence iteration.
func (h *Height) Write(e govfx.Elemental) govfx.DeferWriter {
	newHeight := h.newValue
	return govfx.NewWriter(func() {
		val := fmt.Sprintf("%d%s", newHeight, "px")
		e.Write("height", val, true)
	})
}

//==============================================================================
