package main

import (
	"fmt"
	"time"

	"github.com/influx6/govfx"
	"github.com/influx6/govfx/animations/boundaries"
)

func main() {

	width := govfx.NewAnimationSequence(".zapps",
		govfx.NewStat(govfx.StatConfig{
			Duration: 1 * time.Second,
			Delay:    2 * time.Second,
			Easing:   "ease-in",
			Loop:     4,
			Reverse:  true,
			Optimize: true,
		}),
		&boundaries.Width{Width: 500})

	width.OnBegin(func(stats govfx.Frame) {
		fmt.Println("Animation Has Begun.")
	})

	width.OnEnd(func(stats govfx.Frame) {
		fmt.Println("Animation Has Ended.")
	})

	width.OnProgress(func(stats govfx.Frame) {
		fmt.Println("Animation is progressing.")
	})

	govfx.Animate(width)
}
