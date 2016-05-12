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

//==============================================================================

// Stat provides a configuration for building a Stats object for animators.
type Stat struct {
	Duration time.Duration
	Delay    time.Duration
	Loop     int
	Reverse  bool
}

// SeqBev defines a sequence producer interface.
type SeqBev struct {
	Stat

	once *sync.Once
	seqs SequenceList

	blocks []WriteLink

	ended       Listener
	progressing Listener
	begins      Listener

	bl    sync.RWMutex
	elems Elementals

	flymode  int64
	flyIndex int
}

// NewSeqBev returns a new instance of a SeqBev.
func NewSeqBev(stat Stat, seqs []Sequence) *SeqBev {
	var one sync.Once

	f := SeqBev{
		Stat:        stat,
		seqs:        seqs,
		once:        &one,
		ended:       &listener{},
		begins:      &listener{},
		progressing: &listener{},
	}

	f.ended.Add(func(_ float64) {
		atomic.StoreInt64(&f.flymode, 1)
	})

	return &f
}

// Use sets the giving elementals to be used for animation by the frame.
func (f *SeqBev) Use(elems Elementals) {
	f.bl.Lock()
	defer f.bl.Unlock()

	f.elems = elems
}

// Reset is called by the timer to tell the frame its animation period as finished.
func (f *SeqBev) Reset() {
	var one sync.Once
	f.once = &one
	// atomic.StoreInt64(&f.flymode, 0)
}

//==============================================================================

// WriteLink defines a linked chain of writers that stacks the previous link
// to create a run chain of DeferWriters
type WriteLink struct {
	nexts []DeferWriters
	inits DeferWriters

	elem    Elemental
	lastRun int
}

// Fire calls the current set of DeferWriters to write out their
// details and calls the element to sync.
func (wl *WriteLink) Fire() {
	if len(wl.nexts) == 0 {
		return
	}

	if wl.lastRun >= len(wl.nexts) {
		wl.lastRun = -1
	}

	if wl.lastRun < 0 {
		wl.inits.Write()
		wl.elem.Sync()
		wl.lastRun++
		return
	}

	dl := wl.nexts[wl.lastRun]
	dl.Write()

	wl.lastRun++

	wl.elem.Sync()
}

// Render renders the current frame feeding the delta value if needed to its
// internals.
func (f *SeqBev) Render(delta float64) {
	f.bl.RLock()
	defer f.bl.RUnlock()

	if atomic.LoadInt64(&f.flymode) > 0 {
		for _, blocks := range f.blocks {
			blocks.Fire()
		}
	}

	// if delta != 0 {
	//
	// }

	for ind, elem := range f.elems {
		wl := f.blocks[ind]

		var ws DeferWriters

		for _, seq := range f.seqs {

			// If we can blend using a blending sequence then call the blend
			// function else let the Update function handle it has normal.
			if bs, ok := seq.(BlendingSequence); ok {
				bs.Blend(delta)
				ws = append(ws, seq.Write(elem))
				continue
			}

			//TODO: Does this really make sense, do we want to take the
			// interpolation value as just another update sequence?
			seq.Update(delta)
			ws = append(ws, seq.Write(elem))
		}

		wl.nexts = append(wl.nexts, ws)
		wl.Fire()
	}

}

//==============================================================================

// Update generates the next frame sequence to be rendered and stacks them for
// rendering for the system.
func (f *SeqBev) Update(delta, total float64) {
	if atomic.LoadInt64(&f.flymode) > 0 {
		return
	}

	// If we are just starting with our frame call then initialize initiial state
	// of the frames by calling the .Init() for sequences and record for each
	// element.
	f.once.Do(func() {
		for _, elem := range f.elems {
			var ws DeferWriters

			for _, seq := range f.seqs {
				ws = append(ws, seq.Init(elem))
			}

			wl := WriteLink{
				elem:  elem,
				inits: ws,
			}

			f.blocks = append(f.blocks, wl)
		}
	})

	for _, seq := range f.seqs {
		seq.Update(delta)
	}
}

//==============================================================================

// Listener defines an interface that provides callback hooks.
type Listener interface {
	Add(fn func(float64))
	Emit(float64)
}

type listener struct {
	rl sync.RWMutex
	fx []func(float64)
}

// Emit fires the functions with the provided value.
func (l *listener) Emit(d float64) {
	l.rl.RLock()
	defer l.rl.RUnlock()
	for _, fx := range l.fx {
		fx(d)
	}
}

// Add adds the function into the lists added.
func (l *listener) Add(fx func(float64)) {
	l.rl.Lock()
	defer l.rl.Unlock()
	l.fx = append(l.fx, fx)
}

//==============================================================================
