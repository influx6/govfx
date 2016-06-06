package govfx_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/influx6/govfx"
)

type mob struct{}

var renders int
var updates int

func (mob) Render(dt float64) {
	updates = 0
	fmt.Printf(".{Render %d -> Interpolate %.4f}.\n", renders, dt)
	renders++
	<-time.After(100 * time.Millisecond)
}

func (mob) Update(dt float64, totaltime float64) {
	updates++
}

// TestTimer validates the behaviour of the Timer API.
func TestTimer(t *testing.T) {
	var m mob

	mt := govfx.NewTimer(&m, govfx.ModeTimer{
		Delay:             1 * time.Second,
		MaxMSPerUpdate:    0.01,
		MaxDeltaPerUpdate: 1.5,
	})

	defer govfx.StopTimer(mt)

	for i := 50; i > 0; i-- {
		mt.Update()
		time.Sleep(100 * time.Millisecond)
		if i > 20 {
			mt.Pause()
		}
	}
}
