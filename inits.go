package vfx

import (
	"github.com/influx6/faux/loop"
	"github.com/influx6/faux/loop/web"
)

//==============================================================================

// DefaultEasing defines the default easing key when a invalid easing name is
// given or no easing name is giving.
const DefaultEasing = "ease-in"

// engine is the global gameloop engine used in managing animations within the
// global loop.
var engine loop.GameEngine

// wcache contains all writers cache with respect to each stats for a specific
// frame.
var wcache WriterCache

// easingProviders defines a global registery for easing functions.
var easingProviders EasingProviders

// stopCache contains all loop.Loopers that pertain to any frame, to allow
// stopping any frame immediately
var stopCache *loopCache

//==============================================================================

// Init initializes the animation system with the necessary loop engine,
// desired to be used in running the animation. This is runned by default
// by the runtime using init() functions, but you can reset the animation
// looper using this.
func Init(gear loop.EngineGear) {
	stopCache = newLoopCache()
	wcache = NewDeferWriterCache()
	easingProviders = NewEasingRegister()
	engine = loop.New(gear)
}

// init initializes the selector code before the start of the animators.
func init() {
	Init(web.Loop)
	RegisterEasing("ease-in", EaseIn{})
	RegisterEasing("ease-in-quad", EaseInQuad{})
	RegisterEasing("ease-out-quad", EaseOutQuad{})
	RegisterEasing("ease-in-out-quad", EaseInOutQuad{})
}

//==============================================================================

// GetWriterCache returns the writer cache used by the animation library.
func GetWriterCache() WriterCache {
	return wcache
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
