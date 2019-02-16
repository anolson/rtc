package matrix

import (
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
