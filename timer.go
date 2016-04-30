package govfx

import (
	"fmt"
	"sync/atomic"
	"time"
)

//==============================================================================

// TimeBehaviour defines an interface for timeable structures which want to
// both render and update, it allows Timer to effectively call the appropriate
// method for each step.
type TimeBehaviour interface {
	Render(dt float64)
	Update(dt float64, totalRun float64)
}

//==============================================================================

// oneSixth defines the minimum delta value for a frame.
var oneSixth = 1 / 60

// defaultDelta defines a fixed timestep for each frame.
var defaultDelta = 0.01

// Timer defines a internal clock which calculates appropriate
// elapsed time for animations.
type Timer struct {
	behaviour     TimeBehaviour
	delay         time.Duration
	duration      time.Duration
	totalDuration time.Duration
	delta         time.Duration
	elapsed       time.Time
	start         time.Time
	begin         time.Time
	progressTime  time.Time
	totalDelta    float64
	accumulator   float64
	winding       int64
	progress      int64
	reverse       int64
}

// NewTimer returns a new Timer struct which calculates the delta and elapse time
// each calls of run.
func NewTimer(b TimeBehaviour, duration time.Duration, delay time.Duration) *Timer {
	tm := Timer{behaviour: b, duration: duration, delay: delay}
	tm.totalDuration = duration + delay
	return &tm
}

// InActiveArea returns true/false if the giving timer has entered within the
// active time range i.e after the giving delay.
func (t *Timer) InActiveArea() bool {
	return true
}

// Elapse returns the current elapse time before the
// next update.
func (t *Timer) Elapse() time.Time {
	return t.elapsed
}

// Delta returns the current delta value before the next update.
func (t *Timer) Delta() time.Duration {
	return t.delta
}

// Update updates the timers internal clocks, calculating the necessary durations
// and delta values
func (t *Timer) Update() {
	if !t.hasBegun() {
		t.init()
	}

	now := time.Now()

	t.delta = now.Sub(t.elapsed)
	t.elapsed = now
	t.progressTime = t.progressTime.Add(t.delta)
	t.accumulator += t.delta.Seconds()

	fmt.Printf("Delta: %s Accumulator: %.4f\n", t.delta, t.accumulator)

	for ; t.accumulator > defaultDelta; t.accumulator -= defaultDelta {
		t.behaviour.Update(defaultDelta, t.totalDelta)
		t.totalDelta += defaultDelta
	}

	t.behaviour.Render(t.delta.Seconds())
}

func (t *Timer) init() {
	t.start = time.Now()
	t.elapsed = t.start
	t.progressTime = t.start

	atomic.StoreInt64(&t.winding, 1)

	// Set the direction of the timer.
	if atomic.LoadInt64(&t.reverse) > 0 {
		atomic.StoreInt64(&t.progress, -1)
	} else {
		atomic.StoreInt64(&t.progress, 1)
	}
}

// hasBegun returns true/false if the clock has begun running.
func (t *Timer) hasBegun() bool {
	return atomic.LoadInt64(&t.winding) > 0
}

//==============================================================================
