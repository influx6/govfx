package animators

import (
	"fmt"

	"github.com/influx6/govfx"
)

//==============================================================================

type color struct {
	red   int
	green int
	blue  int
	alpah float64
}

func (c color) String() string {
	return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", c.red, c.green, c.blue, c.alpah)
}

// Color provides a animator for sequencing color animations.
type Color struct {
	Color  string `govfx:"color"`
	Easing string `govfx:"easing"`
	color  color
}

// Init returns the initial writers for the color sequencer.
func (t *Color) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	rw, gw, bw, aw := govfx.ParseRGB(t.Color)
	t.color = color{red: rw, green: gw, blue: bw, alpah: aw}

	var writers govfx.DeferWriters

	for _, elem := range elems {
		func(e govfx.Elemental) {
			color, pr, _ := e.Read("color", "")

			if color == "" {
				color = "rgba(0,0,0,1)"
			}

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("color", color, pr)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

// colorFormat defines the format for rendering color changes.
var colorFormat = `rgba(%.0f,%.0f,%.0f,%.2f)`

// Next returns the next writers for the color sequencer.
func (t *Color) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	easing := govfx.GetEasing(t.Easing)

	var writers govfx.DeferWriters
	var includeAlpha bool

	for _, elem := range elems {
		func(e govfx.Elemental) {
			color, pr, _ := e.Read("color", "")

			var r, g, b int
			var alpha float64

			if govfx.IsRGBFormat(color) {
				r, g, b, alpha = govfx.ParseRGB(color)

				if govfx.IsRGBA(color) {
					includeAlpha = true
				} else {
					alpha = t.color.alpah
				}

			} else {
				r, g, b = govfx.HexToRGB(color)
				alpha = 1
			}

			rx := t.color.red - r
			gx := t.color.green - g
			bx := t.color.blue - b

			var axl float64

			if includeAlpha {
				ax := t.color.alpah - alpha
				axl = easing.Ease(govfx.EaseConfig{
					Stat:         stats,
					CurrentValue: float64(alpha),
					DeltaValue:   float64(ax),
				})
			} else {
				axl = alpha
			}

			rxl := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: float64(r),
				DeltaValue:   float64(rx),
			})

			gxl := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: float64(g),
				DeltaValue:   float64(gx),
			})

			bxl := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: float64(b),
				DeltaValue:   float64(bx),
			})

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("color", fmt.Sprintf(backgroundColorFormat, rxl, gxl, bxl, axl), pr)
				e.Sync()
			}))

		}(elem)
	}

	return writers
}

//==============================================================================

// BackgroundColor provides a animator for sequencing background color animations.
type BackgroundColor struct {
	Color  string `govfx:"color"`
	Easing string `govfx:"easing"`
	color  color
}

// Init returns the initial writers for the color sequencer.
func (t *BackgroundColor) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	rw, gw, bw, aw := govfx.ParseRGB(t.Color)
	t.color = color{red: rw, green: gw, blue: bw, alpah: aw}

	var writers govfx.DeferWriters

	for _, elem := range elems {
		func(e govfx.Elemental) {
			color, pr, _ := e.Read("background-color", "")

			if color == "" {
				color = "none"
			}

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("background-color", color, pr)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}

// backgroundColorFormat defines the format for rendering background-color changes.
var backgroundColorFormat = `rgba(%.0f,%.0f,%.0f,%.2f)`

// Next returns the next writers for the color sequencer.
func (t *BackgroundColor) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	easing := govfx.GetEasing(t.Easing)

	var writers govfx.DeferWriters
	var includeAlpha bool

	for _, elem := range elems {
		func(e govfx.Elemental) {

			color, pr, _ := e.Read("background-color", "")

			var r, g, b int
			var alpha float64

			if govfx.IsRGBFormat(color) {
				r, g, b, alpha = govfx.ParseRGB(color)

				if govfx.IsRGBA(color) {
					includeAlpha = true
				} else {
					alpha = t.color.alpah
				}

			} else {
				r, g, b = govfx.HexToRGB(color)
				alpha = 1
			}

			rx := t.color.red - r
			gx := t.color.green - g
			bx := t.color.blue - b

			var axl float64

			if includeAlpha {
				ax := t.color.alpah - alpha
				axl = easing.Ease(govfx.EaseConfig{
					Stat:         stats,
					CurrentValue: float64(alpha),
					DeltaValue:   float64(ax),
				})
			} else {
				axl = alpha
			}

			rxl := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: float64(r),
				DeltaValue:   float64(rx),
			})

			gxl := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: float64(g),
				DeltaValue:   float64(gx),
			})

			bxl := easing.Ease(govfx.EaseConfig{
				Stat:         stats,
				CurrentValue: float64(b),
				DeltaValue:   float64(bx),
			})

			writers = append(writers, govfx.NewWriter(func() {
				e.Write("background-color", fmt.Sprintf(backgroundColorFormat, rxl, gxl, bxl, axl), pr)
				e.Sync()
			}))
		}(elem)
	}

	return writers
}
