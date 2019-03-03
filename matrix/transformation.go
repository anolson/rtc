package matrix

import (
	"math"

	"github.com/anolson/rtc/tuple"
)

// Transform returns the result of applying a transformation to a Tuple
func Transform(m *Matrix, t *tuple.Tuple) *tuple.Tuple {
	return MultiplyByTuple(m, t)
}

// Chain returns the result of applying a multiple transformations to a Tuple
func Chain(t *tuple.Tuple, transforms ...*Matrix) *tuple.Tuple {
	transformed := t
	for _, transform := range transforms {
		transformed = Transform(transform, transformed)
	}

	return transformed
}

// Identity returns the identiry Matrix
func Identity() *Matrix {
	return New(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
}

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

// RotationZ returns a Matrix for rotating a point around the z axis
func RotationZ(radians float64) *Matrix {
	cos := math.Cos(radians)
	sin := math.Sin(radians)

	return New(4, 4, []float64{
		cos, -sin, 0, 0,
		sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
}

// Shearing returns a Matrix for rotating a point around the z axis
func Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	return New(4, 4, []float64{
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1,
	})
}
