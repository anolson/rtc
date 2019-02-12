package matrix

// Inverse caclulates the inverse of a matrix
func (m *Matrix) Inverse() (*Matrix, error) {
	determinant := m.Determinant()

	if determinant == 0 {
		return nil, ErrNotInvertible
	}

	result := New(m.rows, m.cols, nil)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			cofactor := m.Cofactor(i, j)
			result.Set(j, i, (cofactor / determinant))
		}
	}

	return result, nil
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

// Cofactor caclulates the cofactor of a matrix at row, col
func (m *Matrix) Cofactor(i, j int) float64 {
	minor := m.Minor(i, j)

	if ((i + j) % 2) != 0 {
		return -minor
	}

	return minor
}

// Minor caclulates the minor of a matrix at row, col
func (m *Matrix) Minor(i, j int) float64 {
	return m.Submatrix(i, j).Determinant()
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
