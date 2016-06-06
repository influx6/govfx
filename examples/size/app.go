// +build js

package main

import (
	"time"

	"github.com/influx6/govfx"
	_ "github.com/influx6/govfx/animators"
)

func main() {

	stat := govfx.Stat{
		Duration: 1 * time.Second,
		Delay:    2 * time.Second,
		Loop:     4,
		Reverse:  true,
	}

	props := []govfx.Value{
		{"animate": "width", "easing": "ease-in", "value": 500},
		{"animate": "height", "easing": "ease-in", "value": 200},
	}

	elems := govfx.QuerySelectorAll(".zapps")

	timeline := govfx.Animate(stat, props, elems)
	timeline.Start()

}
