package govfx

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
