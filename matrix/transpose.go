package matrix

// Transpose transposes a matrix, convert the columns in to rows (and vice versa)
func Transpose(m *Matrix) *Matrix {
	result := []float64{}

	for i := 0; i < m.cols; i++ {
		result = append(result, m.Col(i)...)
	}

	return New(m.rows, m.cols, result)
}
