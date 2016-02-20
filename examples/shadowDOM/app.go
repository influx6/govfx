package main

import (
	"fmt"
	"time"

	"github.com/influx6/govfx"
	"github.com/influx6/govfx/animators"
)

func main() {

	root := govfx.NewShadowRoot(govfx.QuerySelector(".root-shadow"))
	elems := root.QuerySelectorAll(".zapps")

	width := govfx.DOMSequence(elems,
		govfx.NewStat(govfx.StatConfig{
			Duration: 1 * time.Second,
			Delay:    2 * time.Second,
			Easing:   "ease-in",
			Loop:     4,
			Reverse:  true,
			Optimize: true,
		}),
		&animators.Width{Value: 500})

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