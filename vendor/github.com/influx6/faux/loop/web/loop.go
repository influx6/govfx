package web

import (
	"sync/atomic"

	"github.com/influx6/faux/loop"
	"github.com/influx6/faux/loop/web/raf"
)

//==============================================================================

// Loop registers the giving callback
func Loop(mx loop.Mux, queue int) loop.Looper {
	sub := Sub{mux: mx, queue: queue}
	sub.Connect()
	return &sub
}

//==============================================================================

// Sub defines a loop subscriber, implements the loop.Looper interface.
type Sub struct {
	mux   loop.Mux
	id    int64
	queue int
}

// End cancels the giving subscriber from the gameloop.
func (s *Sub) End(f ...func()) {
	id := atomic.LoadInt64(&s.id)
	raf.CancelAnimationFrame(int(id), f...)
}

// animate calls the muxer for this subscriber
func (s *Sub) animate(dl float64) {
	s.mux(dl)
	s.Connect()
}

// Connect connects the subscriber into the animation loop.
func (s *Sub) Connect() {
	atomic.StoreInt64(&s.id, int64(raf.RequestAnimationFrame(s.animate)))
}

//==============================================================================
