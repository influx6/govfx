package govfx

//==============================================================================

// EaseIn provides a struct for 'easing-in' based animation.
type EaseIn struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseIn) Ease(d float64) float64 {
	return d
}

//==============================================================================

// EaseInQuad provides a struct for 'easing-in-quad' based animation.
type EaseInQuad struct{}

// Ease returns a new value base on the EaseConfig received.
func (e EaseInQuad) Ease(ms float64) float64 {
	return ms * ms
}

//==============================================================================
