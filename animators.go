package govfx

import (
	"strings"
	"sync"

	"github.com/influx6/faux/reflection"
)

//==============================================================================

// Values represents a map of properties to be merge into animators.
type Values map[string]interface{}

// Animator defines a type of function that recieves a constructor to build
// a sequence from.
type Animator func(Values) Sequence

//==============================================================================

// Animators provides a sequence constructor that provides the ability to
// generate a new sequence using map info.
type Animators interface {
	Get(string) Animator
	Add(string, Animator)
}

// animatorsRegister defines a animators registery that stores different animators
// types keyed by name.
type animatorsRegister struct {
	rl sync.RWMutex
	c  map[string]Animator
}

// NewAnimatorsRegister returns a new instance of animatorsRegister.
func NewAnimatorsRegister() Animators {
	esr := animatorsRegister{c: make(map[string]Animator)}
	return &esr
}

// Get returns the animators provider using the giving name.
func (s *animatorsRegister) Get(name string) Animator {
	name = strings.ToLower(name)

	s.rl.RLock()
	defer s.rl.RUnlock()

	return s.c[name]
}

// Add adds the specific animators provide keyed by the name.
func (s *animatorsRegister) Add(name string, es Animator) {
	name = strings.ToLower(name)

	s.rl.Lock()
	defer s.rl.Unlock()

	s.c[name] = es
}

//==============================================================================

// Merge merges the values within the map with the giving fields of the
// struct passed in using the govfx tag: "govfx".
func Merge(instance interface{}, m Values) Sequence {
	reflection.MergeMap(VFXTag, instance, m)
	return instance.(Sequence)
}
