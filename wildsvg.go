package govfx

import (
	"time"

	"honnef.co/go/js/dom"
)

// CurvePoint defines a area along a svg path for a specific progress and x value.
type CurvePoint struct {
	X        float64
	Y        float64
	Xd       float64
	Yd       float64
	Progress float64
	Length   int
}

// SVGPathEaser defines a easing provider using SVG Paths.
type SVGPathEaser struct {
	conf       SVGConfig
	samples    []CurvePoint
	svgElement dom.Element
	pathlength int
	start      time.Time
	end        time.Time
	delta      time.Duration
}

// SVGConfig defines a configuration object for the SVGPathEaser.
// WARNING: Requires to be loaded within the dom.
type SVGConfig struct {
	SVGPath      string
	Width        int
	Height       int
	SamplingSize int
}

// NewSVGPathEaser returns a new instance of the SVGPathEaser.
func NewSVGPathEaser(conf SVGConfig) *SVGPathEaser {
	var svg SVGPathEaser

	svg.conf = conf
	svg.svgElement = dom.WrapElement(Document().Underlying().Call("createElementNS", "http://www.w3.org/2000/svg", "path"))
	svg.svgElement.SetAttribute("d", conf.SVGPath)

	svg.pathlength = svg.svgElement.Underlying().Call("getTotalLength", nil).Int()

	svg.generateSampling()
	return &svg
}

// Sample for the giving t value between 0..1, we return a CurvePoint that is
// either at that range or within that specified range. Also returns a bool
// whether it was a success or failure.
func (svg *SVGPathEaser) Sample(t float64) (CurvePoint, bool) {
	var item CurvePoint
	var found bool

	for ind, elem := range svg.samples {
		nextInd := ind + 1

		// Clamp the nextInd to ensure it never overflows.
		if nextInd >= svg.conf.SamplingSize {
			nextInd = svg.conf.SamplingSize - 1
		}

		next := svg.samples[nextInd]
		if elem.Xd == t || (elem.Xd < t && next.Xd > t) {
			item = elem
			found = true
			break
		}
	}

	return item, found
}

// Ease returns the giving value of t(time) within the giving svg sample.
func (svg *SVGPathEaser) Ease(t float64) float64 {
	if item, ok := svg.Sample(t); ok {
		return item.Yd
	}
	return 0
}

// generateSampling the samplings based on the sampling size from the
// svg data.
func (svg *SVGPathEaser) generateSampling() {
	samplePct := float64(1 / svg.conf.SamplingSize)

	svg.start = time.Now()

	for i := 0; i < svg.conf.SamplingSize; i++ {
		step := float64(i) * samplePct
		fsm := step * float64(svg.pathlength)
		fsmObj := svg.svgElement.Underlying().Call("getPointAtLength", fsm)

		var point CurvePoint

		point.Progress = step
		point.Length = int(fsm)

		point.X = fsmObj.Get("x").Float()
		point.Xd = point.X / float64(svg.conf.Width)

		point.Y = fsmObj.Get("y").Float()
		point.Yd = 1 - (point.Y / float64(svg.conf.Height))

		svg.samples = append(svg.samples, point)
	}

	svg.end = time.Now()
	svg.delta = svg.end.Sub(svg.start)
}
