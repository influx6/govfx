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
}

// Block represents a single state instance for rendering at a specific moment
// in time.
type Block struct {
	Elem *Element
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

	ended       Listener
	progressing Listener
	begins      Listener
	blocks      []BlockMoment

	elems Elementals
	ideas Values

	flymode  int64
	flyIndex int64
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
		Stat:        stat,
		elems:       elems,
		ended:       &listener{},
		begins:      &listener{},
		progressing: &listener{},
	}

	f.ended.Add(func(_ float64) {
		atomic.StoreInt64(&f.flymode, 1)
	})

	for _, elem := range elems {
		// Add the sequence into the element tree.
		elem.Add(GenerateSequence(ideas)...)

		// Init the properties with the element.
		elem.Init()
	}

	return &f
}

// Reset is called by the timer to tell the frame its animation period as finished.
func (f *SeqBev) Reset() {
	atomic.StoreInt64(&f.flymode, 0)
	for _, elem := range f.elems {
		elem.Reset()
	}
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
		block.Do()
	}

	f.blocks[ind] = blocks
	atomic.AddInt64(&f.flyIndex, 1)
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
