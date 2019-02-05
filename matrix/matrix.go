package matrix

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
	return m.data[index]
}
