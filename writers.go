package govfx

import "time"

//==============================================================================

// DeferWriter provides an interface that allows deferring the write effects
// of a sequence, these way we can collate all effects of a set of sequences
// together to perform a batch right, reducing layout trashing.
type DeferWriter interface {
	Write()
}

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

// WriterCache provides a interface type for writer cache structures, which catch
// animation produced writers per sequence iteration state.
type WriterCache interface {
	Store(Frame, int, ...DeferWriter)
	Writers(Frame, int) DeferWriters
	ClearIteration(Frame, int)
	Clear(Frame)
}

//==============================================================================

// wcache contains all writers cache with respect to each stats for a specific
// frame.
var wcache WriterCache

// GetWriterCache returns the writer cache used by the animation library.
func GetWriterCache() WriterCache {
	return wcache
}

// NewWriter returns a new DeferWriter which executes the provided function
// on its call to Write().
func NewWriter(d func()) DeferWriter {
	dw := dWriter{fx: d}
	return &dw
}

// DWriter defines a concrete type of a DeferWriter which executes a function
// on its call to write
type dWriter struct {
	fx func()
}

// Write executes the internal function attached to this writer.
func (d *dWriter) Write() {
	if d.fx != nil {
		d.fx()
	}
}

//==============================================================================

// FrameController provides an internal controller for Frames to allow
// controlling of a signal of the beginnning and finishing state of their
// sequence writers. This helps to control the pace of which frames release
// writers for the next sequence.
type frameController interface {
	Frame
	BeginWriting()
	DoneWriting()
}

//==============================================================================

// delayedWriter provides a custom writer that forces a delay in its
// write method. This is useful for inserting good delays around
// animation sequences.
type delayedWriter struct {
	ms time.Duration
	f  frameController
}

// Write calls time.After to perform the necessary delay of operation.
func (d *delayedWriter) Write() {
	if !d.f.Stats().IsFirstDone() {
		<-time.After(d.ms)
	}
}

//==============================================================================

// frameBeginWriter is used to indicate to a frame when its writers have
// begun processing.
type frameBeginWriter struct {
	f frameController
}

func (f *frameBeginWriter) Write() {
	f.f.BeginWriting()
}

// frameEndWriter is used to indicate to a frame when its writers have
// ended processing.
type frameEndWriter struct {
	f frameController
}

// Write calls the DoneWriting method to indicate to a frame, that its sequence
// have finished writing.
func (f *frameEndWriter) Write() {
	f.f.DoneWriting()
}

//==============================================================================
