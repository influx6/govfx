package animators

import (
	"fmt"

	"github.com/influx6/govfx"
)

//==============================================================================

// Width provides animation sequencing for width properties, it uses flat integers
// values and pixels.
type Width struct {
	Value int
}

// Init returns the initial writers for the sequence.
func (w *Width) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		width, priority, _ := elem.ReadInt("width")

		(func(e govfx.Elemental) {
			writers = append(writers, govfx.NewWriter(func() {
				val := fmt.Sprintf("%d%s", width, "px")
				e.Write("width", val, priority)
				e.Sync()
			}))
		})(elem)
	}

	return writers
}

// Next returns the writers for the current sequence iteration.
func (w *Width) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(stats.Easing())

	for _, elem := range elems {
		(func(e govfx.Elemental) {
			width, priority, _ := elem.ReadInt("width")

			change := w.Value - width

			newWidth := int(easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: float64(width),
				DeltaValue:   float64(change),
			}))

			writers = append(writers, govfx.NewWriter(func() {
				val := fmt.Sprintf("%d%s", newWidth, "px")
				e.Write("width", val, priority)
				e.Sync()
			}))
		}(elem))

	}

	return writers
}

//==============================================================================

// Height provides animation sequencing for Height properties.
type Height struct {
	Value int
}

// Init returns the initial writers for the sequence.
func (h *Height) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		(func(e govfx.Elemental) {
			height, priority, _ := elem.ReadInt("height")
			writers = append(writers, govfx.NewWriter(func() {
				val := fmt.Sprintf("%d%s", height, "px")
				e.Write("height", val, priority)
				e.Sync()
			}))
		}(elem))
	}

	return writers
}

// Next returns the writers for the current sequence iteration.
func (h *Height) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(stats.Easing())

	for _, elem := range elems {
		(func(e govfx.Elemental) {
			height, priority, _ := e.ReadInt("height")

			change := h.Value - height

			newHeight := int(easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: float64(height),
				DeltaValue:   float64(change),
			}))

			writers = append(writers, govfx.NewWriter(func() {
				val := fmt.Sprintf("%d%s", newHeight, "px")
				e.Write("height", val, priority)
				e.Sync()
			}))
		}(elem))
	}

	return writers
}

//==============================================================================
