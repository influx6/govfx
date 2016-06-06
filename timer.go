package govfx

import (
	"sync"
	"sync/atomic"
	"time"
)

//==============================================================================

// StartableBehaviour defines an interface for calling a start function
// by a timer.
type StartableBehaviour interface {
	Begin(time.Time)
}

//==============================================================================

// TimelineEmitable defines an interface that provides a notification standard for
// structures with the desire to receive notification through a timeline cycle.
type TimelineEmitable interface {
	EmitEnd(float64)
	EmitBegin(float64)
	EmitProgress(float64)
}

// TimelineBehaviour defines a interface for callable structures from a timeline
// provider.
type TimelineBehaviour interface {
	Done() bool
	Reset()
	Completed(int)
	Render(interpolate float64)
	RenderReverse(interpolate float64)
	UpdateReverse(delta float64)
	Update(delta, progress float64, timeline float64)
}

// Timeline defines a struct to manage the behaviour of a animation frame.
type Timeline struct {
	stat Stat
	tb   TimelineBehaviour

	tmMod ModeTimer
	timer Timeable

	start time.Time
	// end   time.Time

	progress float64

	beating int64
	paused  int64
	dead    int64

	loopInfinite bool
	loops        bool

	reversed     bool
	reversedDone bool
	completed    bool

	loop        int
	repeatCount int

	endOnce sync.Once

	timeline time.Duration
}

// NewTimeline returns a new timeline to manage the lifetime of a animation.
func NewTimeline(mt ModeTimer, t TimelineBehaviour, stat Stat) *Timeline {
	tm := Timeline{tmMod: mt, stat: stat, tb: t}

	// Setup loop flags.
	tm.loop = stat.Loop
	tm.loops = (stat.Loop < 0 || stat.Loop > 0)
	tm.loopInfinite = stat.Loop < 0

	return &tm
}

// Resume unpauses the timeline operations if its started.
func (t *Timeline) Resume() {
	if atomic.LoadInt64(&t.beating) < 1 {
		return
	}

	atomic.StoreInt64(&t.paused, 0)
	t.timer.Resume()
}

// Pause pauses the timeline operations if its started.
func (t *Timeline) Pause() {
	if atomic.LoadInt64(&t.beating) < 1 {
		return
	}

	atomic.StoreInt64(&t.paused, 1)
	t.timer.Pause()
}

// Start loads the timeline animation to the run loop.
func (t *Timeline) Start() {
	if atomic.LoadInt64(&t.paused) > 0 {
		return
	}

	t.timer = NewTimer(t, t.tmMod)
	stopCache.Add(t.timer, engine.Loop(func(delta float64) {
		go t.timer.Update()
	}, 0))
}

// Begin sets the timeline ready to begin to clocking its behaviours
// update and render cycles.
func (t *Timeline) Begin(begin time.Time) {
	t.start = begin
	t.timeline = t.stat.Duration + t.stat.Delay
	// t.end = t.start.Add(t.timeline)

	if fb, ok := t.tb.(TimelineEmitable); ok {
		fb.EmitBegin(time.Since(begin).Seconds())
	}
}

// Render implements the TimeBehaviour interface Render() function.
func (t *Timeline) Render(delta float64) {
	if atomic.LoadInt64(&t.paused) > 0 {
		return
	}

	if t.reversed {
		t.tb.RenderReverse(delta)
	} else {
		t.tb.Render(delta)
	}

	if atomic.LoadInt64(&t.dead) < 1 {
		if fb, ok := t.tb.(TimelineEmitable); ok {
			fb.EmitProgress(t.progress)
		}
	}
}

// loopRun calls the looping phase for the timeline.
func (t *Timeline) loopRun() {
	// Pause and stop the current timer, we need a fresh timer
	// to ensure our sequence end time checks works.
	t.timer.Pause()
	StopTimer(t.timer)

	// Reset the behaviour for recall.
	t.tb.Reset()

	// Reset the reverse switches.
	t.reversed = false
	t.reversedDone = false

	// Set progress to 0.
	t.progress = 0

	// Create a new timer and run the clock.
	t.timer = NewTimer(t, t.tmMod)
	stopCache.Add(t.timer, engine.Loop(func(delta float64) {
		go t.timer.Update()
	}, 0))
}

// Update implements the TimeBehaviour interface Update() function.
func (t *Timeline) Update(delta float64, progress float64) {
	if atomic.LoadInt64(&t.paused) > 0 {
		return
	}

	atomic.StoreInt64(&t.beating, 1)

	t.progress = progress

	if t.completed {
		if t.stat.Reverse {

			// If we have reversed and do not loop then end.
			if t.reversed && t.reversedDone && !t.loops {
				return
			}

			// If we have reversed and do loop but the loop is done then end.
			if t.reversed && t.reversedDone && t.loops && t.loop == 0 {
				return
			}
		}

		// If the loops and its not infinite and the loop is done then
		// end.
		if t.loops && !t.loopInfinite && t.loop == 0 {
			return
		}

	}

	// We check the timelines to see if its matching what's expected.
	if t.timeline.Seconds() < progress {

		if !t.completed {
			t.completed = true
			t.tb.Completed(0)
		}

		if t.stat.Reverse {
			if !t.reversed && !t.tb.Done() {
				t.reversed = true
			}

			if t.reversed && !t.tb.Done() {
				t.tb.UpdateReverse(delta)
				return
			}

			t.reversedDone = true
			t.tb.Reset()
		}

		if t.loops {
			if t.loopInfinite {
				t.endOnce.Do(func() {
					atomic.StoreInt64(&t.dead, 1)
					if fb, ok := t.tb.(TimelineEmitable); ok {
						fb.EmitEnd(progress)
					}
				})

				// Pause and stop the current timer, we need a fresh timer
				// to ensure our sequence end time checks works.
				t.timer.Pause()
				t.loopRun()
				return
			}

			t.loop--

			if t.loop > 0 {
				// Pause and stop the current timer, we need a fresh timer
				// to ensure our sequence end time checks works.
				t.timer.Pause()
				t.loopRun()
				return
			}
		}

		t.endOnce.Do(func() {
			atomic.StoreInt64(&t.dead, 1)
			if fb, ok := t.tb.(TimelineEmitable); ok {
				fb.EmitEnd(progress)
			}
		})

		t.timer.Pause()
		StopTimer(t.timer)
		return
	}

	t.tb.Update(delta, progress, progress/t.timeline.Seconds())
}

//==============================================================================

// TimeBehaviour defines an interface for timeable structures which want to
// both render and update, it allows timer to effectively call the appropriate
// method for each step.
type TimeBehaviour interface {
	Render(interpolate float64)
	Update(delta, progress float64)
}

// Timer defines a interface for definining a timer.
type Timer interface {
	Update()
	Pause()
	Resume()
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
		t.behaviour.Update(t.mode.MaxMSPerUpdate, t.totaldelta)
		t.totaldelta += t.mode.MaxMSPerUpdate
		t.accumulator -= t.mode.MaxMSPerUpdate
	}

	interpolate := t.accumulator / t.mode.MaxMSPerUpdate

	t.behaviour.Render(interpolate)
}

// Pause sets the timer loop as inactive.
func (t *timer) Pause() {
	atomic.StoreInt64(&t.stop, 1)
}

// Resume resets the timer loop as active.
func (t *timer) Resume() {
	atomic.StoreInt64(&t.stop, 0)
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
