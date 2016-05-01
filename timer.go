package govfx

import (
	"fmt"
	"sync/atomic"
	"time"
)

// TimeBehaviour defines an interface for timeable structures which want to
// both render and update, it allows timer to effectively call the appropriate
// method for each step.
type TimeBehaviour interface {
	Render(dt float64, current int)
	Update(dt float64, current, total int)
}

//==============================================================================

// TimelineBehaviour defines an interface for creating a timeline management
// system.
type TimelineBehaviour interface {
	TimeBehaviour
	Reset()
}

// Timeline defines a struct to manage the behaviour of a animation frame.
type Timeline struct {
	stat Stat
	tb   TimelineBehaviour
}

// NewTimeline returns a new timeline to manage the lifetime of a animation.
func NewTimeline(t TimelineBehaviour, stat Stat) *Timeline {
	tm := Timeline{stat: stat, tb: t}
	return &tm
}

// Render implements the TimeBehaviour interface Render() function.
func (t *Timeline) Render(delta float64, current int) {}

// Update implements the TimeBehaviour interface Update() function.
func (t *Timeline) Update(delta float64, current, total int) {}

// Sync implements the core operation function for the timeline manager.
func (t *Timeline) Sync() {}

//==============================================================================

// oneSixth defines the minimum delta value for a frame.
var oneSixth = 1 / 60

// defaultDelta defines a fixed timestep for each frame.
var defaultDelta = 0.01

// Timer defines a interface for definining a timer.
type Timer interface {
	Update()
}

// NewTimer returns a new timer struct which calculates the delta and elapse time
// each calls of run.
func NewTimer(b TimeBehaviour, duration time.Duration, delay time.Duration) Timer {
	tm := timer{behaviour: b, duration: duration, delay: delay}
	tm.totalDuration = duration + delay
	return &tm
}

// timer defines a internal clock which calculates appropriate
// elapsed time for animations.
type timer struct {
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

// Update updates the timers internal clocks, calculating the necessary durations
// and delta values
func (t *timer) Update() {
	if !t.hasBegun() {
		t.init()
	}

	now := time.Now()

	t.delta = now.Sub(t.elapsed)
	t.progressTime = t.progressTime.Add(t.delta)

	// If we are still awaiting delay the just add the progress time and
	// return as we are still awaiting delay.
	if now.Sub(t.begin) < 0 {
		return
	}

	totalRun := now.Add(t.delay).Sub(t.start)
	currentRun := totalRun / t.totalDuration

	fmt.Printf("Run: total %.4f current %s\n", now.Sub(t.begin).Seconds(), currentRun)

	t.elapsed = now
	t.accumulator += t.delta.Seconds()

	for ; t.accumulator > defaultDelta; t.accumulator -= defaultDelta {
		t.behaviour.Update(defaultDelta, 0, 0)
		t.totalDelta += defaultDelta
	}

	t.behaviour.Render(t.delta.Seconds(), 0)
}

func (t *timer) init() {
	t.start = time.Now()
	t.begin = t.start.Add(t.delay)
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
func (t *timer) hasBegun() bool {
	return atomic.LoadInt64(&t.winding) > 0
}

//==============================================================================
