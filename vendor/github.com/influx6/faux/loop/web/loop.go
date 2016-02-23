package web

import (
	"sync/atomic"

	"github.com/influx6/faux/loop"
	"github.com/influx6/faux/loop/web/raf"
)

//==============================================================================

// Loop registers the giving callback
func Loop(mx loop.Mux, queue int) loop.Looper {
	sub := sub{mux: mx, queue: queue}
	sub.connect()
	return &sub
}

//==============================================================================

// sub defines a loop subscriber, implements the loop.Looper interface.
type sub struct {
	mux   loop.Mux
	id    int64
	queue int
}

// End cancels the giving subscriber from the gameloop.
func (s *sub) End(f ...func()) {
	id := atomic.LoadInt64(&s.id)
	raf.CancelAnimationFrame(int(id), f...)
}

// animate calls the muxer for this subscriber
func (s *sub) animate(dl float64) {
	s.mux(dl)
	s.connect()
}

// connect connects the subscriber into the animation loop.
func (s *sub) connect() {
	atomic.StoreInt64(&s.id, int64(raf.RequestAnimationFrame(s.animate)))
}

//==============================================================================
