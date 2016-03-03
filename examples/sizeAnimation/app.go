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
		Loop:     4,
		Reverse:  true,
		Animates: []govfx.Value{
			{"animate": "width", "easing": "ease-in", "value": 500},
			{"animate": "translate-y", "easing": "ease", "value": 100},
		},
	}).B(govfx.QuerySelectorAll(".zapps")...)

	govfx.Animate(width)
}
