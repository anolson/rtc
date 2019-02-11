package matrix

import (
	"github.com/anolson/rtc/tuple"
	"github.com/anolson/rtc/util"
)

// Matrix is a data structure for storing a grid of numbers
type Matrix struct {
	rows int
	cols int
	data []float64
}

// New returns a new Matrix object
func New(rows, cols int, data []float64) *Matrix {
	if data == nil {
		data = make([]float64, rows*cols)
	}

	return &Matrix{
		rows: rows,
		cols: cols,
		data: data,
	}
}

// At returns the value at i, j
func (m *Matrix) At(i, j int) float64 {
	index := m.cols*i + j

	if index >= len(m.data) {
		return 0
	}

	return m.data[index]
}

// Set assigns the value at i, j
func (m *Matrix) Set(i, j int, value float64) {
	index := m.cols*i + j

	if index < len(m.data) {
		m.data[index] = value
	}
}

// Equal compares a two matrices
func (m *Matrix) Equal(other *Matrix) bool {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if !util.Approx(m.At(i, j), other.At(i, j)) {
				return false
			}
		}
	}

	return true
}

// Row returns the row at the given index
func (m *Matrix) Row(index int) []float64 {
	offset := m.cols * index

	if offset >= len(m.data) {
		return nil
	}

	return m.data[offset : offset+m.cols]
}

// Col returns the column at the given index
func (m *Matrix) Col(index int) []float64 {
	col := []float64{}

	for i := index; i < len(m.data); i += m.cols {
		col = append(col, m.data[i])
	}

	return col
}

// MultiplyMatrix multiplies two matrices together
func (m *Matrix) MultiplyMatrix(other *Matrix) *Matrix {
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

// MultiplyTuple multiplies a Matrix and a Tuple together
func (m *Matrix) MultiplyTuple(t *tuple.Tuple) *tuple.Tuple {
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

// Transpose transposes a matrix, convert the columns in to rows (and vice versa)
func (m *Matrix) Transpose() *Matrix {
	result := []float64{}

	for i := 0; i < m.cols; i++ {
		result = append(result, m.Col(i)...)
	}

	return New(m.rows, m.cols, result)
}

// Submatrix returns the submatrix by removing the provided row, col
func (m *Matrix) Submatrix(i, j int) *Matrix {
	result := []float64{}

	for k := 0; k < m.rows; k++ {
		if k == i {
			continue
		}

		for rowIndex, value := range m.Row(k) {
			if rowIndex != j {
				result = append(result, value)
			}
		}
	}

	return New(m.rows-1, m.cols-1, result)
}

// Minor caclulates the minor of a matrix at row, col
func (m *Matrix) Minor(i, j int) float64 {
	return m.Submatrix(i, j).Determinant()
}

// Cofactor caclulates the cofactor of a matrix at row, col
func (m *Matrix) Cofactor(i, j int) float64 {
	minor := m.Minor(i, j)

	if ((i + j) % 2) != 0 {
		return -minor
	}

	return minor
}

// Determinant caclulates the determinant of a matrix
func (m *Matrix) Determinant() float64 {
	if m.rows == 2 && m.cols == 2 {
		return m.data[0]*m.data[3] - m.data[1]*m.data[2]
	}

	var result float64
	for j, element := range m.Row(0) {
		result = result + (m.Cofactor(0, j) * element)
	}

	return result
}

func (m *Matrix) Inverse() *Matrix {
	result := New(m.rows, m.cols, nil)
	determinant := m.Determinant()

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			cofactor := m.Cofactor(i, j)
			result.Set(j, i, (cofactor / determinant))
		}
	}

	return result
}

// IsInvertible determines if a matrix can be inverted
func (m *Matrix) IsInvertible() bool {
	return m.Determinant() != 0
}
