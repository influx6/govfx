package govfx

// EaseSpline provides a basic "easer" spline to genearte transistion points.
var EaseSpline = NewSpline(0.25, 0.1, 0.25, 1.0)

// LinearSpline provides a basic "linear" spline to genearte transistion points.
var LinearSpline = NewSpline(0, 0, 1.0, 1.0)

// EaseInSpline provides a basic "easein" spline to genearte transistion points.
var EaseInSpline = NewSpline(0.42, 0.0, 1.0, 1.0)

// EaseOutSpline provides a basic "easeout" spline to genearte transistion
// points.
var EaseOutSpline = NewSpline(0, 0, 0.58, 1.0)

// EaseInOutSpline provides a basic "ease-in-out" spline to genearte transistion
// points.
var EaseInOutSpline = NewSpline(0.42, 0, 0.58, 1.0)

// Spline provides a implmenetation of the Keyspline which uses bezier curves
// to generate the new positional value for change in time.
type Spline struct {
	x1       float64
	x2       float64
	y1       float64
	y2       float64
	optimize bool
}

// NewSpline returns a new spline set with the specific control points.
func NewSpline(x, y, x2, y2 float64) *Spline {
	ss := Spline{x1: x, y1: y, x2: x2, y2: y2}
	return &ss
}

// X returns the provided x value for a giving time between 0 and 1.
func (s *Spline) X(t float64) float64 {
	if s.optimize {
		return t
	}

	if s.x1 == s.y1 && s.x2 == s.y2 {
		s.optimize = true
		return t
	}

	return CalculateBezier(s.GetTimeForX(t), s.y1, s.y2)
}

// GetTimeForX returns the giving time value between 0 and 1 for the provided
// x coordinate for a bezier curve.
func (s *Spline) GetTimeForX(aX float64) float64 {
	// Newton raphson iteration
	aGuessT := aX

	for i := 0; i < 4; i++ {

		currentSlope := GetSlope(aGuessT, s.x1, s.y1)

		if currentSlope == 0.0 {
			return aGuessT
		}

		currentX := CalculateBezier(aGuessT, s.x1, s.x2) - aX

		aGuessT -= currentX / currentSlope

	}

	return aGuessT
}

// // Y returns the provided x value for a giving time between 0 and 1.
// func (s *Spline) Y(t float64) float64 {
// 	if s.optimize {
// 		return t
// 	}
//
// 	if s.x1 == s.y1 && s.x2 == s.y2 {
// 		s.optimize = true
// 		return t
// 	}
//
// 	return CalculateBezier(s.GetTimeForX(t), s.x1, s.x2)
// }
//
// // GetTimeForY returns the giving time value between 0 and 1 for the provided
// // x coordinate for a bezier curve.
// func (s *Spline) GetTimeForY(aY float64) float64 {
// 	// Newton raphson iteration
// 	aGuessT := aY
//
// 	for i := 0; i < 4; i++ {
//
// 		currentSlope := GetSlope(aGuessT, s.y1, s.y2)
//
// 		if currentSlope == 0.0 {
// 			return aGuessT
// 		}
//
// 		currentX := CalculateBezier(aGuessT, s.y1, s.y1) - aY
//
// 		aGuessT -= currentX / currentSlope
//
// 	}
//
// 	return aGuessT
// }

// GetSlope returns dx/dt given t, x1, and x2, or dy/dt given t, y1, and y2.
func GetSlope(aT, aA1, aA2 float64) float64 {
	return 3.0*a(aA1, aA2)*aT*aT + 2.0*b(aA1, aA2)*aT + c(aA1)
}

// CalculateBezier returns x(t) given t, x1, and x2, or y(t) given t, y1, and y2.
func CalculateBezier(aT, aA1, aA2 float64) float64 {
	return ((a(aA1, aA2)*aT+b(aA1, aA2))*aT + c(aA1)) * aT
}

func a(aA1, aA2 float64) float64 { return 1.0 - 3.0*aA2 + 3.0*aA1 }
func b(aA1, aA2 float64) float64 { return 3.0*aA2 - 6.0*aA1 }
func c(aA1 float64) float64      { return 3.0 * aA1 }
