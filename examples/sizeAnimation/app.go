// +build js

package main

import (
	"time"

	"github.com/influx6/govfx"
	_ "github.com/influx6/govfx/animators"
)

func main() {

	width := (govfx.Animation{
		Duration: 1 * time.Second,
		Delay:    2 * time.Second,
		Easing:   "ease-in",
		Loop:     4,
		Reverse:  true,
		Animates: []govfx.Value{
			{"animate": "width", "value": 500},
			{"animate": "translate-y", "value": 100},
		},
	}).B(govfx.QuerySelectorAll(".zapps")...)

	govfx.Animate(width)
}
