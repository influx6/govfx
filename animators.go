package govfx

import (
	"strings"
	"sync"

	"github.com/influx6/faux/reflection"
)

//==============================================================================

// animationProviders defines a global registery for animator functions.
var animationProviders Animators

// RegisterAnimator adds a sequence into the lists with a giving name, this can
// be retrieved later to build a animations lists from.
func RegisterAnimator(name string, ani Animator, defaults Value) {
	animationProviders.Add(name, ani, defaults)
}

//==============================================================================

// Value represents a map of properties to be merge into animators.
type Value map[string]interface{}

// Values defines a lists of value maps.
type Values []Value

// Valueset defines a lists of value slices.
type Valueset []Values

// Animator defines a type of function that recieves a constructor to build
// a sequence from.
type Animator func(defaults, new Value) Sequence

//==============================================================================

// Animators provides a sequence constructor that provides the ability to
// generate a new sequence using map info.
type Animators interface {
	Get(string) (Animator, Value)
	Add(string, Animator, Value)
}

// animatorsRegister defines a animators registery that stores different animators
// types keyed by name.
type animatorsRegister struct {
	rl sync.RWMutex
	c  map[string]Animator
	v  map[string]Value
}

// NewAnimatorsRegister returns a new instance of animatorsRegister.
func NewAnimatorsRegister() Animators {
	esr := animatorsRegister{
		c: make(map[string]Animator),
		v: make(map[string]Value),
	}
	return &esr
}

// Get returns the animators provider using the giving name.
// It also returns the default values used with this struct.
func (s *animatorsRegister) Get(name string) (Animator, Value) {
	name = strings.ToLower(name)

	s.rl.RLock()
	defer s.rl.RUnlock()

	return s.c[name], s.v[name]
}

// Add adds the specific animators provide keyed by the name.
func (s *animatorsRegister) Add(name string, es Animator, defaultVals Value) {
	name = strings.ToLower(name)

	s.rl.Lock()
	defer s.rl.Unlock()

	s.c[name] = es
	s.v[name] = defaultVals
}

//==============================================================================

// VFXTag defines the tag to be associted with a giving struct field definition
// to adequately allow the use of Animators map merge functions.
const VFXTag = "govfx"

// Merge merges the values within the map with the giving fields of the
// struct passed in using the govfx tag: "govfx".
func Merge(instance interface{}, defaults, newVals Value) Sequence {
	if defaults != nil {
		reflection.MergeMap(VFXTag, instance, defaults, false)
	}

	reflection.MergeMap(VFXTag, instance, newVals, false)
	return instance.(Sequence)
}
