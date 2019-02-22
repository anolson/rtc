package matrix

import "math"

// Translation returns a Matrix for moving a point
func Translation(x, y, z float64) *Matrix {
	return New(4, 4, []float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	})
}

// Scaling returns a Matrix for scaling a vector or point
func Scaling(x, y, z float64) *Matrix {
	return New(4, 4, []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	})
}

// RotationX returns a Matrix for rotating a point around the x axis
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

// RotationY returns a Matrix for rotating a point around the y axis
func RotationY(radians float64) *Matrix {
	cos := math.Cos(radians)
	sin := math.Sin(radians)

	return New(4, 4, []float64{
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1,
	})
}
