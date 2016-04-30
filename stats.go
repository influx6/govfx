package govfx

import "time"

//==============================================================================

// GetIterations returns the total iterations for a specific time.Duration that
// is supplied.
func GetIterations(ms time.Duration) int64 {
	return AnimationStepsPerSec * int64(ms.Seconds())
}

//==============================================================================
