package govfx

import (
	"errors"
	"fmt"

	"honnef.co/go/js/dom"

	"github.com/influx6/faux/reflection"
)

//==============================================================================

// Sequence defines a series of animation step which will be runned
// through by calling its .Next() method continousely until the
// sequence is done(if its not a repetitive sequence).
// Sequence when calling their next method, all sequences must return a
// DeferWriter.
type Sequence interface {
	Update(float64)
	Init(Elemental) DeferWriter
	Write(Elemental) DeferWriter
}

// BlendingSequence defines a sequence with a Blend() function which
// allows the timeline to deliver the blender factor to instead of using
// the update method, because we are using a more refined timing mechanism
// during render time the blending factor helps to ensure the properties
// are rendered at the accurate positions without affecting the updated
// value.
type BlendingSequence interface {
	Sequence
	Blend(float64)
}

// SequenceList defines a lists of animatable sequence.
type SequenceList []Sequence

//==============================================================================

// QuerySequence uses a selector to retrieve the desired elements needed
// to be animated, returning the frame for the animation sequence.
func QuerySequence(selector string, stat Stat, s ...Sequence) *SeqBev {
	return ElementalSequence(TransformElements(QuerySelectorAll(selector)), stat, s...)
}

//==============================================================================

// DOMSequence returns a new SeqBev transforming the lists of
// accordingly dom.Elements into its desired elementals for the animation
// sequence.
func DOMSequence(elems []dom.Element, stat Stat, s ...Sequence) *SeqBev {
	return ElementalSequence(TransformElements(elems), stat, s...)
}

//==============================================================================

// ElementalSequence returns a new frame using the selected Elementals for
// the animation sequence.
func ElementalSequence(elems Elementals, stat Stat, s ...Sequence) *SeqBev {
	ani := NewSeqBev(stat, s)
	ani.Use(elems)
	return ani
}

//==============================================================================

// AnimateAttributeName defines the property used to identify the Animator
// referred to by a Animate item.
const AnimateAttributeName = "animate"

// GenerateSequence takes a map of animation properties and builds a sequence list
// from this map.
func GenerateSequence(vals Values) []Sequence {
	var seqs []Sequence

	for _, prop := range vals {
		seq, err := NewSequence(prop[AnimateAttributeName].(string), prop)
		if err != nil {
			panic(err)
		}

		seqs = append(seqs, seq)
	}

	return seqs
}

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
