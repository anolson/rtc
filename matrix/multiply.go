package matrix

import "github.com/anolson/rtc/tuple"

// MultiplyByMatrix multiplies two matrices together
func Multiply(m, other *Matrix) *Matrix {
	result := New(m.rows, m.cols, nil)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.cols; j++ {
			row := m.Row(i)
			col := other.Col(j)

			var accumulator float64
			for k := 0; k < m.rows; k++ {
				accumulator = accumulator + (row[k] * col[k])
			}

			result.Set(i, j, accumulator)
		}
	}

	return result
}

// MultiplyByTuple multiplies a Matrix and a Tuple together
func MultiplyByTuple(m *Matrix, t *tuple.Tuple) *tuple.Tuple {
	result := make([]float64, 4)

	for i := 0; i < m.rows; i++ {
		row := m.Row(i)
		col := []float64{t.X, t.Y, t.Z, t.W}

		var accumulator float64
		for j := 0; j < len(row); j++ {
			accumulator = accumulator + (row[j] * col[j])
		}

		result[i] = accumulator
	}

	return tuple.New(result[0], result[1], result[2], result[3])
}
