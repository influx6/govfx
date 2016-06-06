package govfx

import (
	"sync"

	"github.com/influx6/faux/loop"
)

//==============================================================================

// stopCache contains all loop.Loopers that pertain to any frame, to allow
// stopping any frame immediately
var stopCache = newLoopCache()

// StopTimer stops the frame within the animation step, removing its registered
// loopere.
func StopTimer(t Timer) {
	if looper := stopCache.Get(t); looper != nil {
		looper.End()
		// time.Sleep(1 * time.Millisecond)
		stopCache.Delete(t)
	}
}

//==============================================================================

// loopCache defines a struct for storing loop.Loopers keyed by frames.
type loopCache struct {
	rl sync.RWMutex
	c  map[Timer]loop.Looper
}

// newLoopCache returns a new instance of a loopCache.
func newLoopCache() *loopCache {
	lc := loopCache{c: make(map[Timer]loop.Looper)}
	return &lc
}

// Get returns the looper connected with the frame.
func (s *loopCache) Get(f Timer) loop.Looper {
	s.rl.RLock()
	defer s.rl.RUnlock()
	return s.c[f]
}

// Add sets a looper for a specific frame.
func (s *loopCache) Add(f Timer, l loop.Looper) {
	s.rl.Lock()
	defer s.rl.Unlock()
	s.c[f] = l
}

// Delete removes a looper keyed by its frame.
func (s *loopCache) Delete(f Timer) {
	s.rl.Lock()
	defer s.rl.Unlock()
	delete(s.c, f)
}

//==============================================================================
