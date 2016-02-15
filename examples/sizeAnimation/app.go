package main

import (
	"fmt"
	"time"

	"github.com/influx6/faux/vfx"
	"github.com/influx6/faux/vfx/animations/boundaries"
)

func main() {

	width := vfx.NewAnimationSequence(".zapps",
		vfx.TimeStat(vfx.StatConfig{
			Duration: 1 * time.Second,
			Delay:    2 * time.Second,
			Easing:   "ease-in",
			Loop:     4,
			Reverse:  true,
			Optimize: true,
		}),
		&boundaries.Width{Width: 500})

	width.OnBegin(func(stats vfx.Stats) {
		fmt.Println("Animation Has Begun.")
	})

	width.OnEnd(func(stats vfx.Stats) {
		fmt.Println("Animation Has Ended.")
	})

	width.OnProgress(func(stats vfx.Stats) {
		fmt.Println("Animation is progressing.")
	})

	vfx.Animate(width)
}
