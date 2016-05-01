package govfx

import (
	"sync"

	"github.com/influx6/faux/loop"
)

//==============================================================================

// loopCache defines a struct for storing loop.Loopers keyed by frames.
type loopCache struct {
	rl sync.RWMutex
	c  map[*Frame]loop.Looper
}

// newLoopCache returns a new instance of a loopCache.
func newLoopCache() *loopCache {
	lc := loopCache{c: make(map[*Frame]loop.Looper)}
	return &lc
}

// Get returns the looper connected with the frame.
func (s *loopCache) Get(f *Frame) loop.Looper {
	s.rl.RLock()
	defer s.rl.RUnlock()
	return s.c[f]
}

// Add sets a looper for a specific frame.
func (s *loopCache) Add(f *Frame, l loop.Looper) {
	s.rl.Lock()
	defer s.rl.Unlock()
	s.c[f] = l
}

// Delete removes a looper keyed by its frame.
func (s *loopCache) Delete(f *Frame) {
	s.rl.Lock()
	defer s.rl.Unlock()
	delete(s.c, f)
}

//==============================================================================
