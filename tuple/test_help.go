package tuple

import "math"

// Approx returns true if two float values are approximately equal
func Approx(a, b float64) bool {
	epsilon := 0.00001

	return math.Abs(a-b) < epsilon
}
