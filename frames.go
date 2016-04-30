package govfx

import "time"

//==============================================================================

// AnimationStepsPerSec defines the total steps taking per second of each clock
// tick.
var AnimationStepsPerSec int64 = 60

// FramePhase defines a animation phase type.
type FramePhase int

// const contains sets of Frame phase that identify the current frame animation
// phase.
const (
	NOPHASE FramePhase = iota
	STARTPHASE
	OPTIMISEPHASE
)

// Listeners defines an interface that provides callback hooks for animations.
type Listeners interface {
	OnEnd(func(Frame))
	OnBegin(func(Frame))
	OnProgress(func(Frame))
	ResetListeners()
}

//==============================================================================

// Stat provides a configuration for building a Stats object for animators.
type Stat struct {
	Duration time.Duration
	Delay    time.Duration
	Loop     int
	Reverse  bool
	Optimize bool
}

// Frame defines a sequence producer interface.
type Frame struct {
	Stat
	seqs  SequenceList
	elems Elementals
}

// NewFrame returns a new instance of a Frame.
func NewFrame(stat Stat, seqs SequenceList) *Frame {
	f := Frame{Stat: stat, seqs: seqs}
	return &f
}

// Use sets the giving elementals to be used for animation by the frame.
func (f *Frame) Use(elems Elementals) {
	f.elems = elems
}

// Then stacks the next frame to be called by this frame when it ends or after
// its first complete run if its a infinite looped frame.
func (f *Frame) Then(fr *Frame) {}

// Render renders the current frame feeding the delta value if needed to its
// internals.
func (f *Frame) Render(delta float64) {}

// Update generates the next frame sequence to be rendered and stacks them for
// rendering for the system.
func (f *Frame) Update(delta, runs float64) {}

//==============================================================================
