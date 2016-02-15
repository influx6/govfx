package vfx

import (
	"sync/atomic"
	"time"
)

//==============================================================================

// AnimationStepsPerSec defines the total steps taking per second of each clock
// tick.
var AnimationStepsPerSec int64 = 60

//==============================================================================

// Stats defines a interface which holds stat information
// regarding the current frame and configuration for a sequence.
type Stats interface {
	Delay() time.Duration
	Loop() bool
	TotalLoops() int
	Next(float64)
	Delta() float64
	Clone() Stats
	Easing() string
	IsDone() bool
	IsFirstDone() bool
	Reversed() bool
	Optimized() bool
	Reversible() bool
	CurrentIteration() int
	TotalIterations() int
	DeltaIteration() float64
}

//==============================================================================

// GetIterations returns the total iterations for a specific time.Duration that
// is supplied.
func GetIterations(ms time.Duration) int64 {
	return AnimationStepsPerSec * int64(ms.Seconds())
}

//==============================================================================

// StatConfig provides a configuration for building a Stats object for animators.
type StatConfig struct {
	Duration time.Duration
	Delay    time.Duration
	Easing   string
	Loop     int
	Reverse  bool
	Optimize bool
}

// Stat defines a the stats report strucuture for animation.
type Stat struct {
	config           StatConfig
	delay            time.Duration
	totalIteration   int64
	currentIteration int64
	reversed         int64
	completed        int64
	completedReverse int64
	delta            float64
	done             bool
}

// TimeStat returns a new Stats instance which provide information concering
// the current animation frame, it uses the provided duration to calculate the
// total iteration for the animation.
func TimeStat(config StatConfig) Stats {
	st := Stat{
		config:         config,
		totalIteration: GetIterations(config.Duration),
	}

	return &st
}

// Clone returns a clone for the stats.
func (s *Stat) Clone() Stats {
	return TimeStat(s.config)
}

// Delay returns the time duration defined as the delay before the start of
// a animation sequence on every complete cycle(after a forward+reverse run).
func (s *Stat) Delay() time.Duration {
	return s.config.Delay
}

// Easing returns the easing value for this specifc stat.
func (s *Stat) Easing() string {
	return s.config.Easing
}

// Delta returns the current time delta from the last update.
func (s *Stat) Delta() float64 {
	return s.delta
}

// Next calls the appropriate iteration step for the stat.
func (s *Stat) Next(m float64) {
	if !s.CompletedFirstTransition() {
		s.NextIteration(m)
		return
	}

	if s.Reversible() {
		atomic.StoreInt64(&s.reversed, 1)
		s.PreviousIteration(m)
		return
	}
}

// NextIteration increments the iteration count.
func (s *Stat) NextIteration(m float64) {
	atomic.AddInt64(&s.currentIteration, 1)

	it := atomic.LoadInt64(&s.totalIteration)
	ct := atomic.LoadInt64(&s.currentIteration)

	if ct >= it {
		atomic.StoreInt64(&s.completed, 1)
	}

	s.delta = m
}

// PreviousIteration increments the iteration count.
func (s *Stat) PreviousIteration(m float64) {
	atomic.AddInt64(&s.currentIteration, -1)

	ct := atomic.LoadInt64(&s.currentIteration)

	if ct <= 0 {
		atomic.StoreInt64(&s.completedReverse, 1)
	}

	s.delta = m
}

// Optimized returns true/false if the stat is said to use optimization
// strategies.
func (s *Stat) Optimized() bool {
	return s.config.Optimize
}

// IsFirstDone returns true/false if the stat has completed a full first
// sequence without a reversal.
func (s *Stat) IsFirstDone() bool {
	ct := atomic.LoadInt64(&s.completed)

	if ct <= 0 {
		return false
	}

	return true
}

// IsDone returns true/false if the stat is done.
func (s *Stat) IsDone() bool {
	ct := atomic.LoadInt64(&s.completed)

	if ct <= 0 {
		return false
	}

	if !s.Reversible() {
		return true
	}

	rs := atomic.LoadInt64(&s.completedReverse)

	if rs <= 0 {
		return false
	}

	return true
}

// Reversed returns true/false if the stats has entered a reversed state.
func (s *Stat) Reversed() bool {
	return atomic.LoadInt64(&s.reversed) > 0
}

// Reversible returns true/false if the stat animation is set to loop.
func (s *Stat) Reversible() bool {
	return s.config.Reverse
}

// CompletedFirstTransition returns true/false if the stat has completed a full
// iteration of its total iteration, this is useful to know when loop or
// reversal is turned on, to check if this stat has entered its looping or
// reversal state. It only reports for the first completion of total iterations.
func (s *Stat) CompletedFirstTransition() bool {
	return atomic.LoadInt64(&s.completed) > 0
}

// TotalLoops returns the total count of loops for this stat. If value is
// less than 0, it is considered an infinite loop.
func (s *Stat) TotalLoops() int {
	return s.config.Loop
}

// Loop returns true/false if the stat animation is set to loop.
func (s *Stat) Loop() bool {
	return s.config.Loop != 0
}

// TotalIterations returns the total iteration for this specific stat.
func (s *Stat) TotalIterations() int {
	return int(atomic.LoadInt64(&s.totalIteration))
}

// CurrentIteration returns the current iteration for this specific stat.
func (s *Stat) CurrentIteration() int {
	return int(atomic.LoadInt64(&s.currentIteration))
}

// DeltaIteration reduces the delta value of currentIteration/TotalIterations,
// useful for easing calculations.
func (s *Stat) DeltaIteration() float64 {
	return float64(s.CurrentIteration()) / float64(s.TotalIterations())
}

//==============================================================================
