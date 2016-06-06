package utils

import (
	"fmt"
	"sync/atomic"
)

//==============================================================================

// Incr provides a struct which houses a ever-increasing ID.
type Incr struct {
	uuid  string
	steps int64
	add   int64
}

//==============================================================================

// NewIncr returns a new instance of the Incr struct.
// Its one arguments allows you to set the starting point for the
// counter.
func NewIncr(tag string, point int) *Incr {
	in := Incr{
		uuid:  tag,
		steps: int64(point),
		add:   1,
	}

	return &in
}

// NewIncrBy returns a new instance of the Incr struct.
// It allos you to provide a starting point and incrementer value
// for how the values increases over time.
func NewIncrBy(tag string, point int, by int) *Incr {
	in := Incr{
		uuid:  tag,
		steps: int64(point),
		add:   int64(by),
	}

	return &in
}

//==============================================================================

// New returns a new ID from the incrementer.
func (in *Incr) New() string {
	atomic.AddInt64(&in.steps, in.add)
	return fmt.Sprintf("%d:%s", atomic.LoadInt64(&in.steps), in.uuid)
}
