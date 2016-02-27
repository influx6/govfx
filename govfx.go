package govfx

import "honnef.co/go/js/dom"

//==============================================================================

// X defines a function that allows the construction of a Animation by Animation
// sequence where each Animation within the lists begins after the ending of
// the previous animation, it allows a simple way of stacking multiple sequences
// easily.
func X(ems Elementals, a ...Animation) Frame {
	var initial, current Frame

	for _, ani := range a {

		// If we are just starting the loop, set the initial and current frame
		// as needed.
		if initial == nil {
			initial = ani.B(ems...)
			current = initial
			continue
		}

		// Generate the new frame and apply it to the last frame as its next
		// Frame call, then set this as the current frame.
		f := ani.B(ems...)
		current.Then(f)
		current = f
	}

	return initial
}

//==============================================================================

// QuerySequence uses a selector to retrieve the desired elements needed
// to be animated, returning the frame for the animation sequence.
func QuerySequence(selector string, stat Stats, s ...Sequence) Frame {
	return ElementalSequence(TransformElements(QuerySelectorAll(selector)), stat, s...)
}

//==============================================================================

// DOMSequence returns a new Frame transforming the lists of
// accordingly dom.Elements into its desired elementals for the animation
// sequence.
func DOMSequence(elems []dom.Element, stat Stats, s ...Sequence) Frame {
	return ElementalSequence(TransformElements(elems), stat, s...)
}

//==============================================================================

// ElementalSequence returns a new frame using the selected Elementals for
// the animation sequence.
func ElementalSequence(elems Elementals, stat Stats, s ...Sequence) Frame {
	ani := NewAnimationSequence(stat, s...)
	ani.Use(elems)
	return ani
}
