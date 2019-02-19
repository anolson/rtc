package matrix

import "math"

// Translation returns a matrix for moving a point
func Translation(x, y, z float64) *Matrix {
	return New(4, 4, []float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	})
}

// Scaling returns a matrix for scaling a vector or point
func Scaling(x, y, z float64) *Matrix {
	return New(4, 4, []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	})
}

func RotationX(radians float64) *Matrix {
	cos := math.Cos(radians)
	sin := math.Sin(radians)

	return New(4, 4, []float64{
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1,
	})
}
