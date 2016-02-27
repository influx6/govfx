package govfx

import (
	"errors"
	"fmt"

	"github.com/influx6/faux/reflection"
)

//==============================================================================

// DelaySequence returns a new sequence type that checks if a sequence is allowed
// to be runnable during a sequence iteration.
type DelaySequence interface {
	Continue() bool
}

// StoppableSequence defines a interface for sequences that can be stopped.
type StoppableSequence interface {
	Stop()
}

// Sequence defines a series of animation step which will be runned
// through by calling its .Next() method continousely until the
// sequence is done(if its not a repetitive sequence).
// Sequence when calling their next method, all sequences must return a
// DeferWriter.
type Sequence interface {
	Init(Stats, Elementals) DeferWriters
	Next(Stats, Elementals) DeferWriters
}

// SequenceList defines a lists of animatable sequence.
type SequenceList []Sequence

//==============================================================================

// NewSequence returns a new sequence tagged by the giving name, using the
// values map to initialize the attributes accordingly, else returns an
// error if the sequence name does not exists.
func NewSequence(name string, m Value) (Sequence, error) {
	ani, defaults := animationProviders.Get(name)
	if ani == nil {
		return nil, fmt.Errorf("No Sequence with Name[%s]", name)
	}

	return ani(defaults, m), nil
}

// RegisterSequence adds a sequence by taking a sample value type of the real struct
// that provides that and generating a new one when requested.
func RegisterSequence(name string, structType interface{}) error {
	if !reflection.IsStruct(structType) {
		return errors.New("Not a Struct")
	}

	d, _ := reflection.ToMap(VFXTag, structType, false)

	animationProviders.Add(name, func(d, m Value) Sequence {
		newSeq, _ := reflection.MakeNew(structType)
		return Merge(newSeq, d, m)
	}, d)

	return nil
}

//==============================================================================
