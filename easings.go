package govfx

import (
	"strings"
	"sync"
)

// DefaultEasing defines the default easing key when a invalid easing name is
// given or no easing name is giving.
const DefaultEasing = "ease-in"

var easingProviders = NewEasingRegister()

//==============================================================================

// CSS3Easings defines the different easing functions within the css3 specs
// for use with css transformation rules.
var CSS3Easings = map[string]string{
	"ease-in":           "ease-in",
	"ease-out":          "ease-out",
	"ease-in-out":       "ease-in-out",
	"snap":              "cubic-bezier(0,1,.5,1)",
	"linear":            "cubic-bezier(0.250, 0.250, 0.750, 0.750)",
	"ease-in-quad":      "cubic-bezier(0.550, 0.085, 0.680, 0.530)",
	"ease-in-cubic":     "cubic-bezier(0.550, 0.055, 0.675, 0.190)",
	"ease-in-quart":     "cubic-bezier(0.895, 0.030, 0.685, 0.220)",
	"ease-in-quint":     "cubic-bezier(0.755, 0.050, 0.855, 0.060)",
	"ease-in-sine":      "cubic-bezier(0.470, 0.000, 0.745, 0.715)",
	"ease-in-expo":      "cubic-bezier(0.950, 0.050, 0.795, 0.035)",
	"ease-in-circ":      "cubic-bezier(0.600, 0.040, 0.980, 0.335)",
	"ease-in-back":      "cubic-bezier(0.600, -0.280, 0.735, 0.045)",
	"ease-out-quad":     "cubic-bezier(0.250, 0.460, 0.450, 0.940)",
	"ease-out-cubic":    "cubic-bezier(0.215, 0.610, 0.355, 1.000)",
	"ease-out-quart":    "cubic-bezier(0.165, 0.840, 0.440, 1.000)",
	"ease-out-quint":    "cubic-bezier(0.230, 1.000, 0.320, 1.000)",
	"ease-out-sine":     "cubic-bezier(0.390, 0.575, 0.565, 1.000)",
	"ease-out-expo":     "cubic-bezier(0.190, 1.000, 0.220, 1.000)",
	"ease-out-circ":     "cubic-bezier(0.075, 0.820, 0.165, 1.000)",
	"ease-out-back":     "cubic-bezier(0.175, 0.885, 0.320, 1.275)",
	"ease-out-quad2":    "cubic-bezier(0.455, 0.030, 0.515, 0.955)",
	"ease-out-cubic2":   "cubic-bezier(0.645, 0.045, 0.355, 1.000)",
	"ease-in-out-quart": "cubic-bezier(0.770, 0.000, 0.175, 1.000)",
	"ease-in-out-quint": "cubic-bezier(0.860, 0.000, 0.070, 1.000)",
	"ease-in-out-sine":  "cubic-bezier(0.445, 0.050, 0.550, 0.950)",
	"ease-in-out-expo":  "cubic-bezier(1.000, 0.000, 0.000, 1.000)",
	"ease-in-out-circ":  "cubic-bezier(0.785, 0.135, 0.150, 0.860)",
	"ease-in-out-back":  "cubic-bezier(0.680, -0.550, 0.265, 1.550)",
}

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

// RegisterEasing adds a easing provider into the registery with the specified
// name, we allow replacing a easing provider for a keyed name, if you so wish.
func RegisterEasing(name string, easing Easing) {
	easingProviders.Add(name, easing)
}

// GetEasingProvider returns the central easing provider for vfx.
func GetEasingProvider() EasingProviders {
	return easingProviders
}

// GetEasing returns the easing function matching the specific easing function
// name if it exists else it returns the default easing provider set by
// DefaultEasing constant.
func GetEasing(easing string) Easing {
	es := easingProviders.Get(easing)
	if es == nil {
		es = easingProviders.Get(DefaultEasing)
	}

	return es
}

//==============================================================================

// Easing defines a interface that returns a new value for the provided values.
type Easing interface {
	Ease(float64) float64
}

// EasingProviders provides a interface type to expose easing function providers.
type EasingProviders interface {
	Get(string) Easing
	Add(string, Easing)
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
