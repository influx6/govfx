package govfx_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/influx6/govfx"
)

type mob struct{}

func (mob) Render(dt float64, cu int) {
	fmt.Printf("...")
}

func (mob) Update(dt float64, cu, total int) {
	fmt.Printf("||")
}

// TestTimer validates the behaviour of the Timer API.
func TestTimer(t *testing.T) {
	var m mob

	mt := govfx.NewTimer(&m, 4*time.Second, 1*time.Second)

	for i := 50; i > 0; i-- {
		mt.Update()
		time.Sleep(300 * time.Millisecond)
	}
}
