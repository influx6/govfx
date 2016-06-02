package govfx

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// TimeBehaviour defines an interface for timeable structures which want to
// both render and update, it allows timer to effectively call the appropriate
// method for each step.
type TimeBehaviour interface {
	Render(interpolate float64)
	Update(delta, progress float64)
}

// StartableBehaviour defines an interface for calling a start function
// by a timer.
type StartableBehaviour interface {
	Begin(time.Time)
}

//==============================================================================

// TimelineBehaviour defines an interface for creating a timeline management
// system.
type TimelineBehaviour interface {
	TimeBehaviour
	// Completed()
	Reset()
}

// Timeline defines a struct to manage the behaviour of a animation frame.
type Timeline struct {
	stat  Stat
	timer Timeable
	tb    TimelineBehaviour

	start    time.Time
	end      time.Time
	progress time.Time

	beating int64
	paused  int64

	repeatCount int

	endOnce sync.Once

	timeline time.Duration
}

// NewTimeline returns a new timeline to manage the lifetime of a animation.
func NewTimeline(tc Timeable, t TimelineBehaviour, stat Stat) *Timeline {
	tm := Timeline{stat: stat, tb: t, timer: tc}
	tc.Use(&tm)
	return &tm
}

// Resume unpauses the timeline operations if its started.
func (t *Timeline) Resume() {
	if atomic.LoadInt64(&t.beating) < 1 {
		return
	}

	atomic.StoreInt64(&t.paused, 0)
}

// Pause pauses the timeline operations if its started.
func (t *Timeline) Pause() {
	if atomic.LoadInt64(&t.beating) < 1 {
		return
	}

	atomic.StoreInt64(&t.paused, 1)
}

// Start loads the timeline animation to the run loop.
func (t *Timeline) Start() {
	stopCache.Add(t.timer, engine.Loop(func(delta float64) {
		t.timer.Update()
	}, 0))
}

// Begin sets the timeline ready to begin to clocking its behaviours
// update and render cycles.
func (t *Timeline) Begin(begin time.Time) {
	t.start = begin
	t.progress = begin.Add(t.stat.Delay)
	t.timeline = t.stat.Duration + t.stat.Delay
	t.end = t.start.Add(t.timeline)

	if fb, ok := t.tb.(*SeqBev); ok {
		fb.begins.Emit(time.Since(begin).Seconds())
	}
}

// Render implements the TimeBehaviour interface Render() function.
func (t *Timeline) Render(delta float64) {
	if atomic.LoadInt64(&t.paused) > 0 {
		return
	}

	t.tb.Render(delta)
	if fb, ok := t.tb.(*SeqBev); ok {
		fb.progressing.Emit(time.Since(t.progress).Seconds())
	}
}

// Update implements the TimeBehaviour interface Update() function.
func (t *Timeline) Update(delta float64, progress float64) {
	if atomic.LoadInt64(&t.paused) > 0 {
		return
	}

	atomic.StoreInt64(&t.beating, 1)
	t.progress = t.progress.Add(time.Duration(progress) * time.Second)

	fmt.Println("Will stop: ", time.Since(t.progress))
	if t.timeline < time.Since(t.progress) {
		fmt.Println("Stopping")
		t.timer.Stop()
		stop(t.timer)

		t.endOnce.Do(func() {
			if fb, ok := t.tb.(*SeqBev); ok {
				fb.ended.Emit(progress)
			}
		})

		return
	}

	t.tb.Update(delta, progress)
}

//==============================================================================

// Timer defines a interface for definining a timer.
type Timer interface {
	Update()
	Stop()
}

// Timeable defines an interface that defines a Timer confirming
// structure with the ability to set the TimeBehaviour to use.
type Timeable interface {
	Timer
	Use(TimeBehaviour)
}

// maxMSPerUpdate defines the maximum tick for which our updates
// are called within our game loop.
var maxMSPerUpdate = 0.01
var maxUpdateRuns = 0.25

// ModeTimer defines a configuration for seting the behaviour of a
// timer loop.
type ModeTimer struct {
	Delay             time.Duration
	MaxMSPerUpdate    float64
	MaxDeltaPerUpdate float64
}

// NewTimer returns a new timer struct which calculates the delta and elapse time
// each calls of run.
func NewTimer(b TimeBehaviour, mod ModeTimer) Timeable {
	tm := timer{behaviour: b, mode: mod}
	return &tm
}

//==============================================================================

// timer defines a internal clock which calculates appropriate
// elapsed time for animations.
type timer struct {
	ml        sync.RWMutex
	behaviour TimeBehaviour
	mode      ModeTimer

	start    time.Time
	initial  time.Time
	end      time.Time
	previous time.Time
	progress time.Time

	elapsed     float64
	accumulator float64
	totaldelta  float64

	prevState float64
	curState  float64

	delta      time.Duration
	lastDelta  time.Duration
	totalDelta time.Duration

	run      int64
	stop     int64
	skipTick float64
}

// Use sets the behaviour to be used by the timer for its update
// process.
func (t *timer) Use(d TimeBehaviour) {
	t.ml.Lock()
	defer t.ml.Unlock()
	t.behaviour = d
}

// Update updates the timers internal clocks, calculating the necessary durations
// and delta values
func (t *timer) Update() {
	t.ml.RLock()
	defer t.ml.RUnlock()

	if t.behaviour == nil {
		return
	}

	if !t.hasBegun() {
		t.init()
	}

	if atomic.LoadInt64(&t.stop) > 0 {
		return
	}

	now := time.Now()

	t.lastDelta = t.delta
	t.delta = now.Sub(t.previous)
	t.previous = now

	t.progress = t.progress.Add(t.delta)

	if t.progress.Before(t.initial) {
		return
	}

	var dt float64

	if t.delta.Seconds() > t.mode.MaxDeltaPerUpdate {
		dt = t.mode.MaxDeltaPerUpdate
	} else {
		dt = t.delta.Seconds()
	}

	t.accumulator += dt

	for t.accumulator >= t.mode.MaxMSPerUpdate {
		// t.prevState = t.curState
		t.behaviour.Update(t.mode.MaxMSPerUpdate, t.totaldelta)
		t.totaldelta += t.mode.MaxMSPerUpdate
		t.accumulator -= t.mode.MaxMSPerUpdate
	}

	interpolate := t.accumulator / t.mode.MaxMSPerUpdate
	// t.curState = t.curState*interpolate + t.prevState*(1.0-interpolate)

	t.behaviour.Render(interpolate)
}

// Stop sets the timer loop as inactive.
func (t *timer) Stop() {
	atomic.StoreInt64(&t.stop, 1)
}

// init initializes the details of the time for work.
func (t *timer) init() {
	t.start = time.Now()
	t.previous = t.start
	t.progress = t.start
	t.initial = t.start.Add(t.mode.Delay)
	atomic.StoreInt64(&t.run, 1)

	t.ml.RLock()
	defer t.ml.RUnlock()
	if t.behaviour != nil {
		if so, ok := t.behaviour.(StartableBehaviour); ok {
			so.Begin(t.start)
		}
	}
}

// hasBegun returns true/false if the clock has begun running.
func (t *timer) hasBegun() bool {
	return atomic.LoadInt64(&t.run) > 0
}

//==============================================================================
