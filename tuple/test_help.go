package tuple

import "math"

func floatEqual(a, b float64) bool {
	epsilon := 0.00001

	return math.Abs(a-b) < epsilon
}
