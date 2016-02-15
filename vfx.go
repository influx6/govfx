package vfx

//==============================================================================

// WriterCache provides a interface type for writer cache structures, which catch
// animation produced writers per sequence iteration state.
type WriterCache interface {
	Store(Frame, int, ...DeferWriter)
	Writers(Frame, int) DeferWriters
	ClearIteration(Frame, int)
	Clear(Frame)
}

//==============================================================================

// DeferWriter provides an interface that allows deferring the write effects
// of a sequence, these way we can collate all effects of a set of sequences
// together to perform a batch right, reducing layout trashing.
type DeferWriter interface {
	Write()
}

//==============================================================================

// DeferWriters defines a lists of DeferWriter implementing structures.
type DeferWriters []DeferWriter

// Write calls the internal writers Write() method, within the
// DeferWriters lists.
func (d DeferWriters) Write() {
	for _, dw := range d {
		dw.Write()
	}
}

//==============================================================================

// CascadeDeferWriter provides a specifc writer that can combine multiple writers
// into a single one but also allow flatten its writer sequence into a ordered
// lists of DeferWriters.
type CascadeDeferWriter interface {
	DeferWriter
	Flatten() DeferWriters
}

//==============================================================================

// Sequence defines a series of animation step which will be runned
// through by calling its .Next() method continousely until the
// sequence is done(if its not a repetitive sequence).
// Sequence when calling their next method, all sequences must return a
// DeferWriter.
type Sequence interface {
	Init(Stats, Elementals) DeferWriters
	Next(Stats, Elementals) DeferWriters
}

// DelaySequence returns a new sequence type that checks if a sequence is allowed
// to be runnable during a sequence iteration.
type DelaySequence interface {
	Continue() bool
}

// StoppableSequence defines a interface for sequences that can be stopped.
type StoppableSequence interface {
	Stop()
}

// SequenceList defines a lists of animatable sequence.
type SequenceList []Sequence
