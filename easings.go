package govfx

import (
	"strings"
	"sync"
)

//==============================================================================

// Easing defines a interface that returns a new value for the provided values.
type Easing interface {
	Ease(esc EaseConfig) float64
}

// EasingProviders provides a interface type to expose easing function providers.
type EasingProviders interface {
	Get(string) Easing
	Add(string, Easing)
}

// EaseConfig provides a easing configuration to help simplify easing
// calculations. This is provided to unify the
type EaseConfig struct {
	Stats        Stats
	DeltaValue   float64
	CurrentValue float64
}

//==============================================================================

// easingRegister defines a easing registery that stores different easing
// types keyed by name.
type easingRegister struct {
	rl sync.RWMutex
	c  map[string]Easing
}

// NewEasingRegister returns a new instance of easingRegister.
func NewEasingRegister() EasingProviders {
	esr := easingRegister{c: make(map[string]Easing)}
	return &esr
}

// Get returns the easing provider using the giving name.
func (s *easingRegister) Get(name string) Easing {
	name = strings.ToLower(name)

	s.rl.RLock()
	defer s.rl.RUnlock()

	return s.c[name]
}

// Add adds the specific easing provide keyed by the name.
func (s *easingRegister) Add(name string, es Easing) {
	name = strings.ToLower(name)

	s.rl.Lock()
	defer s.rl.Unlock()

	s.c[name] = es
}

//==============================================================================
