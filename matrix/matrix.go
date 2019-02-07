package matrix

import (
	"github.com/anolson/rtc/util"
)

type Matrix struct {
	rows int
	cols int
	data []float64
}

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

func (m *Matrix) At(i, j int) float64 {
	index := m.cols*i + j

	if index >= len(m.data) {
		return 0
	}

	return m.data[index]
}

func (m *Matrix) Set(i, j int, value float64) {
	index := m.cols*i + j

	if index < len(m.data) {
		m.data[index] = value
	}
}

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

func (m *Matrix) Row(index int) []float64 {
	offset := m.cols * index

	if offset >= len(m.data) {
		return nil
	}

	return m.data[offset : offset+m.cols]
}

func (m *Matrix) Col(index int) []float64 {
	col := []float64{}

	for i := index; i < len(m.data); i += m.cols {
		col = append(col, m.data[i])
	}

	return col
}

func (m *Matrix) Multiply(other *Matrix) *Matrix {
	result := New(m.rows, m.cols, nil)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
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
