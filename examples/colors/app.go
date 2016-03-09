// +build js

package main

import (
	"time"

	"github.com/influx6/govfx"
	_ "github.com/influx6/govfx/animators"
)

func main() {

	govfx.Animate(govfx.Animation{
		Duration: 3 * time.Second,
		Delay:    2 * time.Second,
		Loop:     -1,
		Reverse:  true,
		Animates: []govfx.Value{
			{"animate": "scale-x", "easing": "ease", "value": 0.5},
			{"animate": "background-color", "easing": "ease-in-out", "color": "rgb(201, 30, 93)"},
		},
	}.B(govfx.QuerySelectorAll(".boxy")...))

	govfx.Animate(govfx.Animation{
		Duration: 3 * time.Second,
		Delay:    2 * time.Second,
		Loop:     -1,
		Reverse:  true,
		Animates: []govfx.Value{
			{"animate": "scale-x", "easing": "ease", "value": 0.5},
			{"animate": "background-color", "easing": "ease-in-out", "color": "rgb(222, 233, 241)"},
		},
	}.B(govfx.QuerySelectorAll(".box")...))

}
