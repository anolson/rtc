package util

import "math"

// Approx returns true if two float values are approximately equal
func Approx(a, b float64) bool {
	epsilon := 1e-5

	return math.Abs(a-b) < epsilon
}
