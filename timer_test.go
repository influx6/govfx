package govfx_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/influx6/govfx"
)

type mob struct{}

func (mob) Update(dt, totalRun float64) {
	fmt.Printf("Update: Delta:%.2f Run:%.2f\n", dt, totalRun)
}

func (mob) Render(dt float64) {
	fmt.Printf("Render: %.2f \n", dt)
}

// TestTimer validates the behaviour of the Timer API.
func TestTimer(t *testing.T) {
	var m mob

	mt := govfx.NewTimer(&m, 3*time.Second, 0)

	for i := 50; i > 0; i-- {
		mt.Update()
		time.Sleep(300 * time.Millisecond)
	}
}
