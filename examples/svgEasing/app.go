package main

import (
	"fmt"
	"time"

	"github.com/influx6/govfx"
	_ "github.com/influx6/govfx/animators"
)

func main() {

	begin := govfx.NewListener(func(dl float64) {
		fmt.Printf("Animation Has Begun at %.4f .\n", dl)
	})

	end := govfx.NewListener(func(dl float64) {
		fmt.Printf("Animation Has Ended at %.4f .\n", dl)
	})

	progress := govfx.NewListener(func(dl float64) {
		fmt.Printf("Animation Is	 Progressing at %.4f .\n", dl)
	})

	easer := govfx.NewSVGPathEaser(govfx.SVGConfig{
		SVGPath:      "M0,100 C40,100 50,90 50,50 C50,10 60,3.46944696e-18 100,0",
		Width:        100,
		Height:       100,
		SamplingSize: 300,
	})

	elems := govfx.QuerySelectorAll(".zapps")
	width := govfx.Animate(govfx.Stat{
		Duration: 1 * time.Second,
		Loop:     2,
		Reverse:  true,
		Begin:    begin,
		End:      end,
		Progress: progress,
	}, govfx.Values{
		{"value": 500, "animate": "width", "easer": easer},
	}, elems)

	<-width.Simulate()

	width.Start()

}
