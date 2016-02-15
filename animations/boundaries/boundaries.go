package boundaries

import (
	"fmt"

	"github.com/influx6/govfx"
)

//==============================================================================

// WidthCSSWriter defines a DeferWriter for writing width properties.
type WidthCSSWriter struct {
	width    int
	unit     string
	priority bool
	elem     govfx.Elemental
}

// Write writes out the necessary output for a css width property.
func (w *WidthCSSWriter) Write() {
	val := fmt.Sprintf("%d%s", w.width, w.unit)
	w.elem.Write("width", val, w.priority)
	w.elem.Sync()
}

//==============================================================================

// Width provides animation sequencing for width properties.
type Width struct {
	Width int
}

// Init returns the initial writers for the sequence.
func (w *Width) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	for _, elem := range elems {
		width, priority, _ := elem.Read("width")
		writers = append(writers, &WidthCSSWriter{
			width:    govfx.ParseInt(width),
			unit:     "px",
			priority: priority,
			elem:     elem,
		})
	}

	return writers
}

// Next returns the writers for the current sequence iteration.
func (w *Width) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	var writers govfx.DeferWriters

	easing := govfx.GetEasing(stats.Easing())

	for _, elem := range elems {
		width, priority, _ := elem.Read("width")

		realWidth := govfx.ParseInt(width)
		change := w.Width - realWidth

		newWidth := int(easing.Ease(govfx.EaseConfig{
			Stat:         stats,
			CurrentValue: float64(realWidth),
			DeltaValue:   float64(change),
		}))

		writers = append(writers, &WidthCSSWriter{
			width:    newWidth,
			unit:     "px",
			priority: priority,
			elem:     elem,
		})
	}

	return writers
}

//==============================================================================
