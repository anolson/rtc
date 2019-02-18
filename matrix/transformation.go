package matrix

// Translation returns a matrix for moving a point
func Translation(x, y, z float64) *Matrix {
	return New(4, 4, []float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	})
}
