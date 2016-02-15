package govfx

//==============================================================================

// EaseIn provides a struct for 'easing-in' based animation.
type EaseIn struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseIn) Ease(c EaseConfig) float64 {
	return (c.DeltaValue * c.Stats.DeltaIteration()) + c.CurrentValue
}

//==============================================================================

// EaseInQuad provides a struct for 'easing-in-quad' based animation.
type EaseInQuad struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseInQuad) Ease(c EaseConfig) float64 {
	ms := c.Stats.DeltaIteration()
	return (c.DeltaValue * ms * ms) + c.CurrentValue
}

//==============================================================================

// EaseOutQuad provides a struct for 'easing-out-quad' based animation.
type EaseOutQuad struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseOutQuad) Ease(c EaseConfig) float64 {
	ms := (c.Stats.DeltaIteration()) * float64(c.Stats.CurrentIteration()-2)
	return ((c.DeltaValue * -1) * ms) + c.CurrentValue
}

//==============================================================================

// EaseInOutQuad provides a struct for 'easing-in-out-quad' based animation.
type EaseInOutQuad struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseInOutQuad) Ease(c EaseConfig) float64 {
	diff := c.Stats.DeltaIteration()

	if diff < 1 {
		return (c.DeltaValue / 2) * diff * diff
	}

	diff--

	return (-1*c.DeltaValue)*((diff)*(diff-2)-1) + c.CurrentValue
}

//==============================================================================
