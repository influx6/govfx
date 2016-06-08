package govfx

import (
	"time"

	"honnef.co/go/js/dom"
)

// CurvePoint defines a area along a svg path for a specific progress and x value.
type CurvePoint struct {
	X        float64
	Y        float64
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
}

// SVGConfig defines a configuration object for the SVGPathEaser.
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
	svg.svgElement = Document().CreateElementNS("http://w3.org/2000/svg", "path")
	svg.svgElement.SetAttribute("d", conf.SVGPath)

	svg.pathlength = svg.svgElement.Underlying().Call("getTotalLength", nil).Int()

	svg.generateSampling()
	return &svg
}

// Ease returns the giving value of t(time) within the giving svg sample.
func (svg *SVGPathEaser) Ease(t float64) float64 {
	return 0
}

// generateSampling the samplings based on the sampling size from the
// svg data.
func (svg *SVGPathEaser) generateSampling() {
	samplePct := 1 / svg.conf.SamplingSize

	svg.start = time.Now()

	for i := 0; i < svg.conf.SamplingSize; i++ {
		fsm := i * samplePct * svg.pathlength
		fsmObj := svg.svgElement.Underlying().Call("getPointAtLength", fsm)

		var point CurvePoint

		point.Length = svg.pathlength
		point.X = fsmObj.Get("x").Float() / float64(svg.conf.Width)
		point.Y = 1 - fsmObj.Get("y").Float()/float64(svg.conf.Height)
		svg.samples = append(svg.samples, point)
	}

	svg.end = time.Now()
}
