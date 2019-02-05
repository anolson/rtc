package matrix

import "github.com/anolson/rtc/util"

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
