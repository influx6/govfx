package govfx

import (
	"sync"
	"sync/atomic"
	"time"
)

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

// WriteLink defines a linked chain of writers that stacks the previous link
// to create a run chain of DeferWriters
type WriteLink struct {
	Writers []DeferWriters
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
	seqs   SequenceList
	inited int64
	elems  Elementals
	bl     sync.RWMutex
	blocks []WriteLink
}

// NewFrame returns a new instance of a Frame.
func NewFrame(stat Stat, seqs SequenceList) *Frame {
	f := Frame{Stat: stat, seqs: seqs}
	return &f
}

// Use sets the giving elementals to be used for animation by the frame.
func (f *Frame) Use(elems Elementals) {
	f.bl.Lock()
	defer f.bl.Unlock()

	f.elems = elems
	f.blocks = make([]WriteLink, len(f.elems))
}

// Then stacks the next frame to be called by this frame when it ends or after
// its first complete run if its a infinite looped frame.
func (f *Frame) Then(fr *Frame) {

}

// Reset is called by the timer to tell the frame its animation period as finished.
func (f *Frame) Reset() {
}

//==============================================================================

// Render renders the current frame feeding the delta value if needed to its
// internals.
func (f *Frame) Render(delta float64, current int) {
	f.bl.RLock()
	defer f.bl.RUnlock()

	// Loop through all the blocks and lunch their internal writers.
	for _, wl := range f.blocks {
		block := lastBlock(wl.Writers)
		block.Write()
	}

}

func lastBlock(m []DeferWriters) DeferWriters {
	size := len(m)

	if size > 0 {
		return m[size-1]
	}

	return nil
}

func lastWriter(m []DeferWriter) DeferWriter {
	size := len(m)
	if size > 0 {
		return m[size-1]
	}

	return nil
}

//==============================================================================

// Update generates the next frame sequence to be rendered and stacks them for
// rendering for the system.
func (f *Frame) Update(delta float64, current, total int) {

	// If we are just starting with our frame call then initialize initiial state
	// of the frames by calling the .Init() for sequences and record for each
	// element.
	if atomic.LoadInt64(&f.inited) < 0 {
		var ws DeferWriters

		for ind, elem := range f.elems {
			for _, seq := range f.seqs {
				ws = append(ws, seq.Init(delta, elem))
			}

			wl := f.blocks[ind]
			wl.Writers = append(wl.Writers, ws)
		}

		atomic.StoreInt64(&f.inited, 1)
		// TODO: Should we return here or let the initialization be part of the
		// first startup.
		// return
	}

	// Call the .Next() method call for the sequences for each element to have
	// them begin intializing their update cycle.
	var ws DeferWriters

	for ind, elem := range f.elems {
		for _, seq := range f.seqs {
			ws = append(ws, seq.Next(delta, elem))
		}

		wl := f.blocks[ind]
		wl.Writers = append(wl.Writers, ws)
	}

}

//==============================================================================
