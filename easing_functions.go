package govfx

//==============================================================================

// Linear provides a struct for 'linear' based animation.
type Linear struct{}

// Ease returns a new value base on the EaseConfig received.
func (l Linear) Ease(c EaseConfig) float64 {
	return c.DeltaValue*LinearSpline.X(c.Stat.DeltaIteration()) + c.CurrentValue
}

//==============================================================================

// EaseIn provides a struct for 'easing-in' based animation.
type EaseIn struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseIn) Ease(c EaseConfig) float64 {
	return c.DeltaValue*EaseInSpline.X(c.Stat.DeltaIteration()) + c.CurrentValue
}

//==============================================================================

// EaseOut provides a struct for 'easing-out' based animation.
type EaseOut struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseOut) Ease(c EaseConfig) float64 {
	return c.DeltaValue*EaseOutSpline.X(c.Stat.DeltaIteration()) + c.CurrentValue
}

//==============================================================================

// EaseInOut provides a struct for 'easing-in-out' based animation.
type EaseInOut struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseInOut) Ease(c EaseConfig) float64 {
	return c.DeltaValue*EaseInOutSpline.X(c.Stat.DeltaIteration()) + c.CurrentValue
}

//==============================================================================
