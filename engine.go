package govfx

import (
	"strings"

	"github.com/fatih/camelcase"
	"github.com/influx6/faux/loop"
	"github.com/influx6/faux/loop/web"
)

//==============================================================================

// Animate provides the central engine for managing all animation calls.
// Animate uses writer batching to reduce layout trashing. Hence  each frame
// assigned for each animation call, will have all their writes batched
// into one call.
func Animate(frame *Frame) {

	// create the timer for which the frame will be animated.
	timer := NewTimer(frame, frame.Stat.Duration, frame.Stat.Delay)

	// Return this frame subscription ender, initialized and run its writers.
	stopCache.Add(frame, engine.Loop(func(delta float64) {
		timer.Update()
	}, 0))
}

// Stop stops the frame within the animation step, removing its registered
// loopere.
func Stop(frame *Frame) {
	looper := stopCache.Get(frame)
	if looper != nil {
		stopCache.Delete(frame)
		looper.End()
	}
}

//==============================================================================

// engine is the global gameloop engine used in managing animations within the
// global loop.
var engine loop.GameEngine

// stopCache contains all loop.Loopers that pertain to any frame, to allow
// stopping any frame immediately
var stopCache *loopCache

// Init initializes the animation system with the necessary loop engine,
// desired to be used in running the animation. This is runned by default
// by the runtime using init() functions, but you can reset the animation
// looper using this.
func Init(gear loop.EngineGear) {
	stopCache = newLoopCache()
	wcache = NewDeferWriterCache()
	easingProviders = NewEasingRegister()
	animationProviders = NewAnimatorsRegister()
	engine = loop.New(gear)
}

// init initializes the selector code before the start of the animators.
func init() {
	Init(web.Loop)

	// Register all our easing providers.
	for name, vals := range EasingValues {
		cased := strings.ToLower(strings.Join(camelcase.Split(name), "-"))
		RegisterEasing(cased, NewSpline(vals[0], vals[1], vals[2], vals[3]))
	}
}

//==============================================================================
