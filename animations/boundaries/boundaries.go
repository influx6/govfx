package boundaries

import (
	"fmt"

	"github.com/influx6/faux/vfx"
)

//==============================================================================

// WidthCSSWriter defines a DeferWriter for writing width properties.
type WidthCSSWriter struct {
	width    int
	unit     string
	priority bool
	elem     vfx.Elemental
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
func (w *Width) Init(stats vfx.Stats, elems vfx.Elementals) vfx.DeferWriters {
	var writers vfx.DeferWriters

	for _, elem := range elems {
		width, priority, _ := elem.Read("width")
		writers = append(writers, &WidthCSSWriter{
			width:    vfx.ParseInt(width),
			unit:     "px",
			priority: priority,
			elem:     elem,
		})
	}

	return writers
}

// Next returns the writers for the current sequence iteration.
func (w *Width) Next(stats vfx.Stats, elems vfx.Elementals) vfx.DeferWriters {
	var writers vfx.DeferWriters

	easing := vfx.GetEasing(stats.Easing())

	for _, elem := range elems {
		width, priority, _ := elem.Read("width")

		realWidth := vfx.ParseInt(width)
		change := w.Width - realWidth

		newWidth := int(easing.Ease(vfx.EaseConfig{
			CurrentStep:  stats.CurrentIteration(),
			TotalSteps:   stats.TotalIterations(),
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
