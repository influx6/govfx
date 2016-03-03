// +build js

package main

import (
	"time"

	"github.com/influx6/govfx"
	_ "github.com/influx6/govfx/animators"
)

func main() {

	root := govfx.NewShadowRoot(govfx.QuerySelector(".root-shadow"))

	width := (govfx.Animation{
		Duration: 1 * time.Second,
		Delay:    2 * time.Second,
		Easing:   "ease-in",
		Loop:     0,
		Reverse:  true,
		Animates: []govfx.Value{
			{"animate": "width", "easing": "ease-in", "value": 500},
			{"animate": "height", "easing": "ease-in", "value": 10},
		},
	}).B(root.QuerySelectorAll(".zapps")...)

	govfx.Animate(width)
}
