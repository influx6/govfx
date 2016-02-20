package govfx

import (
	"strings"
	"sync"
)

//==============================================================================

// EasingValues provides a map of precalculated easing transition values.
var EasingValues = map[string][]float64{
	"ease":           {0.25, 0.1, 0.25, 1.0},
	"linear":         {0, 0, 1.0, 1.0},
	"easeIn":         {0.42, 0.0, 1.0, 1.0},
	"easeOut":        {0, 0, 0.58, 1.0},
	"easeInOut":      {0.42, 0, 0.58, 1.0},
	"easeInSine":     {0.47, 0, 0.745, 0.715},
	"easeOutSine":    {0.39, 0.575, 0.565, 1},
	"easeInOutSine":  {0.445, 0.05, 0.55, 0.95},
	"easeInQuad":     {0.55, 0.085, 0.68, 0.53},
	"easeOutQuad":    {0.25, 0.46, 0.45, 0.94},
	"easeInOutQuad":  {0.455, 0.03, 0.515, 0.955},
	"easeInCubic":    {0.55, 0.055, 0.675, 0.19},
	"easeOutCubic":   {0.215, 0.61, 0.355, 1},
	"easeInOutCubic": {0.645, 0.045, 0.355, 1},
	"easeInQuart":    {0.895, 0.03, 0.685, 0.22},
	"easeOutQuart":   {0.165, 0.84, 0.44, 1},
	"easeInOutQuart": {0.77, 0, 0.175, 1},
	"easeInQuint":    {0.755, 0.05, 0.855, 0.06},
	"easeOutQuint":   {0.23, 1, 0.32, 1},
	"easeInOutQuint": {0.86, 0, 0.07, 1},
	"easeInExpo":     {0.95, 0.05, 0.795, 0.035},
	"easeOutExpo":    {0.19, 1, 0.22, 1},
	"easeInOutExpo":  {1, 0, 0, 1},
	"easeInCirc":     {0.6, 0.04, 0.98, 0.335},
	"easeOutCirc":    {0.075, 0.82, 0.165, 1},
	"easeInOutCirc":  {0.785, 0.135, 0.15, 0.86},
}

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
	Stat         Stats
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
