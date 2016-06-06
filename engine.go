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
func Animate(stat Stat, b Values, elems Elementals) *Timeline {
	frame := NewSeqBev(elems, stat, b)
	return NewTimeline(ModeTimer{
		Delay:             stat.Delay,
		MaxMSPerUpdate:    0.01,
		MaxDeltaPerUpdate: 2.5,
	}, frame, stat)
}

//==============================================================================

var engine loop.GameEngine

// Init initializes the animation system with the necessary loop engine,
// desired to be used in running the animation. This is runned by default
// by the runtime using init() functions, but you can reset the animation
// looper using this.
func Init(gear loop.EngineGear) {
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
