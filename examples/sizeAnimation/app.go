package main

import (
	"time"

	"github.com/influx6/govfx"
	_ "github.com/influx6/govfx/animators"
)

func main() {

	elems := govfx.QuerySelectorAll(".zapps")
	width := govfx.Animate(govfx.Stat{
		Duration: 4 * time.Second,
		Loop:     2,
	}, govfx.Values{
		{"value": 500, "animate": "width", "easing": "ease-in"},
	}, elems)

	width.Start()

	// width.OnBegin(func(stats govfx.Frame) {
	// 	fmt.Println("Animation Has Begun.")
	// })
	//
	// width.OnEnd(func(stats govfx.Frame) {
	// 	fmt.Println("Animation Has Ended.")
	// })
	//
	// width.OnProgress(func(stats govfx.Frame) {
	// 	fmt.Println("Animation is progressing.")
	// })

	// govfx.Animate(width)
}
