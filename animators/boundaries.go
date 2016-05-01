package animators

import (
	"fmt"

	"github.com/influx6/govfx"
)

//==============================================================================

// Width provides animation sequencing for width properties, it uses flat integers
// values and pixels.
type Width struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (w *Width) Init(delta float64, elem govfx.Elemental) govfx.DeferWriters {
	var writers govfx.DeferWriters

	width, priority, _ := elem.ReadInt("width", "")

	return govfx.NewWriter(func() {
		val := fmt.Sprintf("%d%s", width, "px")
		elem.Write("width", val, priority)
		elem.Sync()
	})
}

// Next returns the writers for the current sequence iteration.
func (w *Width) Next(delta float64, e govfx.Elemental) govfx.DeferWriters {

	easing := govfx.GetEasing(w.Easing)

	width, priority, _ := elem.ReadInt("width", "")

	change := w.Value - width

	newWidth := int(easing.Ease(govfx.EaseConfig{
		Stat:         stats,
		CurrentValue: float64(width),
		DeltaValue:   float64(change),
	}))

	return govfx.NewWriter(func() {
		val := fmt.Sprintf("%d%s", newWidth, "px")
		e.Write("width", val, priority)
		e.Sync()
	})
}

//==============================================================================

// Height provides animation sequencing for Height properties.
type Height struct {
	Value  int    `govfx:"value"`
	Easing string `govfx:"easing"`
}

// Init returns the initial writers for the sequence.
func (h *Height) Init(stats govfx.Stats, elem govfx.Elemental) govfx.DeferWriters {

	height, priority, _ := elem.ReadInt("height", "")
	return govfx.NewWriter(func() {
		val := fmt.Sprintf("%d%s", height, "px")
		e.Write("height", val, priority)
		e.Sync()
	})

}

// Next returns the writers for the current sequence iteration.
func (h *Height) Next(stats govfx.Stats, e govfx.Elemental) govfx.DeferWriters {

	easing := govfx.GetEasing(h.Easing)

	height, priority, _ := e.ReadInt("height", "")

	change := h.Value - height

	newHeight := int(easing.Ease(govfx.EaseConfig{
		Stat:         stats,
		CurrentValue: float64(height),
		DeltaValue:   float64(change),
	}))

	return govfx.NewWriter(func() {
		val := fmt.Sprintf("%d%s", newHeight, "px")
		e.Write("height", val, priority)
		e.Sync()
	})

}

//==============================================================================
