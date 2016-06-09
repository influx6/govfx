package govfx

import (
	"bytes"
	"sync"
	"sync/atomic"
	"time"

	"honnef.co/go/js/dom"
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
	Begin    Listener
	End      Listener
	Progress Listener
}

// Block represents a single state instance for rendering at a specific moment
// in time.
type Block struct {
	Elem Elemental
	Buf  *bytes.Buffer
}

// Do writes the giving buffer into the style attribute of the element.
func (b *Block) Do() {
	b.Elem.SetAttribute("style", b.Buf.String())
}

// BlockMoment represents a full moment or rendering of the state of a element
// in time.
type BlockMoment []Block

// Run calls all the Do() methods of its internal blocks.
func (b BlockMoment) Run() {
	for _, block := range b {
		// TODO: should we Go-routine this, to ensure elements update asynchronousely?
		block.Do()
	}
}

// SeqBev defines a sequence producer interface.
type SeqBev struct {
	Stat

	blocks []BlockMoment

	reversing bool
	reversed  bool

	elems Elementals
	ideas Values

	flymode  int64
	flyIndex int64
	simMode  int64
}

// QuerySequence uses a selector to retrieve the desired elements needed
// to be animated, returning the frame for the animation sequence.
func QuerySequence(selector string, stat Stat, vs Values) *SeqBev {
	return ElementalSequence(TransformElements(QuerySelectorAll(selector)), stat, vs)
}

// DOMSequence returns a new SeqBev transforming the lists of
// accordingly dom.Elements into its desired elementals for the animation
// sequence.
func DOMSequence(elems []dom.Element, stat Stat, vs Values) *SeqBev {
	return ElementalSequence(TransformElements(elems), stat, vs)
}

// ElementalSequence returns a new frame using the selected Elementals for
// the animation sequence.
func ElementalSequence(elems Elementals, stat Stat, id Values) *SeqBev {
	ani := NewSeqBev(elems, stat, id)
	return ani
}

// NewSeqBev returns a new instance of a SeqBev.
func NewSeqBev(elems Elementals, stat Stat, ideas Values) *SeqBev {
	f := SeqBev{
		Stat:  stat,
		elems: elems,
	}

	for _, elem := range elems {
		// Add the sequence into the element tree.
		elem.Add(GenerateSequence(ideas)...)

		// Init the properties with the element.
		elem.Init()
	}

	return &f
}

// SimulationOFF puts off the sequence frame simulation mode returning things
// back to normal operations.
func (f *SeqBev) SimulationOFF() {
	atomic.StoreInt64(&f.simMode, 0)
}

// SimulationON puts the sequence frame into simulation mode where
// the operations are performed but not written to display.
func (f *SeqBev) SimulationON() {
	atomic.StoreInt64(&f.simMode, 1)
}

// Completed signifies the completion of the sequence.
func (f *SeqBev) Completed(cycle int) {
	atomic.StoreInt64(&f.flymode, 1)
}

// Done returns true/false if the sequence has completed a full run.
// Where a full run is a completed cycle + reveres run.
func (f *SeqBev) Done() bool {
	flymod := int(atomic.LoadInt64(&f.flymode))

	if atomic.LoadInt64(&f.flyIndex) <= 0 {
		f.reversed = true
	}

	if flymod < 1 {
		return false
	}

	if f.Stat.Reverse && !f.reversed {
		return false
	}

	return true
}

// Reset is called by the timer to tell the frame its animation period as finished.
func (f *SeqBev) Reset() {
	f.reversed = false
	f.reversing = false
	atomic.StoreInt64(&f.flyIndex, 0)
}

// RenderReverse reverses the rendering of the sequence by calling the
// index in reverese.
func (f *SeqBev) RenderReverse(delta float64) {
	ind := int(atomic.LoadInt64(&f.flyIndex))
	total := len(f.blocks)

	if ind >= total {
		ind = total - 1
	}

	if ind < 0 {
		ind = 0
	}

	blocks := f.blocks[ind]

	if atomic.LoadInt64(&f.simMode) < 1 {
		blocks.Run()
	}

	atomic.AddInt64(&f.flyIndex, -1)
}

// Render renders the current frame feeding the delta value if needed to its
// internals.
func (f *SeqBev) Render(delta float64) {
	flymod := int(atomic.LoadInt64(&f.flymode))

	ind := atomic.LoadInt64(&f.flyIndex)

	if int(ind) >= len(f.blocks) {
		f.blocks = append(f.blocks, []Block{})
	}

	blocks := f.blocks[ind]

	if flymod > 0 {
		blocks.Run()
		atomic.AddInt64(&f.flyIndex, 1)
		return
	}

	// Build the blocks list for this current index.
	for _, elem := range f.elems {
		elem.Blend(delta)

		var buf bytes.Buffer
		elem.CSS(&buf)

		block := Block{
			Elem: elem,
			Buf:  &buf,
		}

		blocks = append(blocks, block)

		if int(atomic.LoadInt64(&f.simMode)) < 1 {
			block.Do()
		}
	}

	f.blocks[ind] = blocks
	atomic.AddInt64(&f.flyIndex, 1)
}

//==============================================================================

// EmitBegin emits the begin signal to the listener supplied in the stat.
func (f *SeqBev) EmitBegin(delta float64) {
	if f.Stat.Begin != nil {
		f.Stat.Begin.Emit(delta)
	}
}

// EmitProgress emits the progress signal to the listener supplied in the stat.
func (f *SeqBev) EmitProgress(delta float64) {
	if f.Stat.Progress != nil {
		f.Stat.Progress.Emit(delta)
	}
}

// EmitEnd emits the ending signal to the listener supplied in the stat.
func (f *SeqBev) EmitEnd(delta float64) {
	if f.Stat.End != nil {
		f.Stat.End.Emit(delta)
	}
}

//==============================================================================

// Update generates the next frame sequence to be rendered and stacks them for
// rendering for the system.
func (f *SeqBev) Update(delta, total float64, timeline float64) {
	if atomic.LoadInt64(&f.flymode) > 0 {
		return
	}

	for _, elem := range f.elems {
		elem.Update(delta, timeline)
	}
}

// UpdateReverse calls a reverse procedure on the sequence being runned.
func (f *SeqBev) UpdateReverse(delta float64) {
}

//==============================================================================

// Listener defines an interface that provides callback hooks.
type Listener interface {
	Add(fn func(float64))
	Emit(float64)
}

// NewListener returns a new instance of a structure that matches the Listener
// interface.
func NewListener(cbs ...func(float64)) Listener {
	var lm listener

	for _, item := range cbs {
		lm.Add(item)
	}

	return &lm
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
