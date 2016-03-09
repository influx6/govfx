// +build js

package main

import (
	"time"

	"github.com/influx6/govfx"
	_ "github.com/influx6/govfx/animators"
)

func main() {

	colorScale := govfx.X(govfx.Animation{
		Duration: 3 * time.Second,
		Delay:    2 * time.Second,
		Loop:     1,
		Reverse:  true,
		Animates: []govfx.Value{
			{"animate": "background-color", "easing": "ease-in-out", "color": "rgb(201, 30, 93)"},
		},
	}, govfx.Animation{
		Duration: 3 * time.Second,
		Delay:    2 * time.Second,
		Loop:     1,
		Reverse:  true,
		Animates: []govfx.Value{
			{"animate": "scale-x", "easing": "ease", "value": 1.0},
		},
	})

	colorScale.Use(govfx.QuerySelectorAll(".boxy"))
	govfx.Animate(colorScale)

	colorScale2 := govfx.X(govfx.Animation{
		Duration: 3 * time.Second,
		Delay:    2 * time.Second,
		Loop:     2,
		Reverse:  true,
		Animates: []govfx.Value{
			{"animate": "scale-x", "easing": "ease", "value": 1.0},
			{"animate": "background-color", "easing": "ease-in-out", "color": "rgb(222, 233, 241)"},
		},
	})

	colorScale2.Use(govfx.QuerySelectorAll(".box"))
	govfx.Animate(colorScale2)

}
