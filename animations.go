package govfx

import "time"

//==============================================================================

// AnimateAttributeName defines the property used to identify the Animator
// referred to by a Animate item.
const AnimateAttributeName = "animate"

// Animation defines a new API structure that helps to ease the creation of a
// new sequence frames and simplifes the creation of animations.
type Animation struct {
	Loop     int
	Reverse  bool
	Duration time.Duration
	Delay    time.Duration
	Animates Values
}

// Animations defines a lists of Animation structs.
type Animations []Animation

// B returns a new animation frame build from the data provided by the animation
// structure.
func (a Animation) B(e ...Elemental) Frame {

	// Loop through all Animate data building the sequence
	// using the Animator registry
	var seqs SequenceList

	for _, prop := range a.Animates {
		seq, err := NewSequence(prop[AnimateAttributeName].(string), prop)
		if err != nil {
			// fmt.Printf("Error occured %s\n", err)
			panic(err)
			// continue
		}

		seqs = append(seqs, seq)
	}

	as := NewAnimationSequence(NewStat(StatConfig{
		Duration: a.Duration,
		Delay:    a.Delay,
		Loop:     a.Loop,
		Reverse:  a.Reverse,
		Optimize: true,
	}), seqs...)

	as.Use(e)

	return as
}

//==============================================================================
