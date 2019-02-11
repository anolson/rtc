package matrix

import (
	"testing"

	"github.com/anolson/rtc/tuple"
	"github.com/anolson/rtc/util"
	"github.com/stvp/assert"
)

func TestNew(t *testing.T) {
	t.Run("Creating a 2x2 matrix", func(t *testing.T) {
		m := New(2, 2, []float64{
			-3, 5,
			1, -2,
		})

		assert.Equal(t, -3.0, m.At(0, 0))
		assert.Equal(t, 5.0, m.At(0, 1))
		assert.Equal(t, 1.0, m.At(1, 0))
		assert.Equal(t, -2.0, m.At(1, 1))
	})

	t.Run("Creating a 3x3 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			-3, 5, 0,
			1, -2, -7,
			0, 1, 1,
		})

		assert.Equal(t, -3.0, m.At(0, 0))
		assert.Equal(t, -2.0, m.At(1, 1))
		assert.Equal(t, 1.0, m.At(2, 2))
	})

	t.Run("Creating a 4x4 matrix", func(t *testing.T) {
		m := New(4, 4, []float64{
			1, 2, 3, 4,
			5.5, 6.5, 7.5, 8.5,
			9, 10, 11, 12,
			13.5, 14.5, 15.5, 16.5,
		})

		assert.Equal(t, 1.0, m.At(0, 0))
		assert.Equal(t, 4.0, m.At(0, 3))
		assert.Equal(t, 5.5, m.At(1, 0))
		assert.Equal(t, 7.5, m.At(1, 2))
		assert.Equal(t, 11.0, m.At(2, 2))
		assert.Equal(t, 13.5, m.At(3, 0))
		assert.Equal(t, 15.5, m.At(3, 2))
	})
}

func TestEqual(t *testing.T) {
	m := New(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	other := New(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2.000001,
	})

	otherNotEqual := New(4, 4, []float64{
		2, 3, 4, 5,
		6, 7, 8, 9,
		8, 7, 6, 5,
		4, 3, 2, 1,
	})

	assert.True(t, m.Equal(other))
	assert.False(t, m.Equal(otherNotEqual))
}

func TestRow(t *testing.T) {
	m := New(3, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	assert.Equal(t, []float64{1.0, 2.0, 3.0, 4.0}, m.Row(0))
	assert.Equal(t, []float64{5.0, 6.0, 7.0, 8.0}, m.Row(1))
	assert.Equal(t, []float64{9.0, 8.0, 7.0, 6.0}, m.Row(2))
	assert.Equal(t, []float64{5.0, 4.0, 3.0, 2.0}, m.Row(3))
}

func TestCol(t *testing.T) {
	m := New(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	assert.Equal(t, []float64{1.0, 5.0, 9.0, 5.0}, m.Col(0))
	assert.Equal(t, []float64{2.0, 6.0, 8.0, 4.0}, m.Col(1))
	assert.Equal(t, []float64{3.0, 7.0, 7.0, 3.0}, m.Col(2))
	assert.Equal(t, []float64{4.0, 8.0, 6.0, 2.0}, m.Col(3))
}

func TestMultiplyMatrix(t *testing.T) {
	a := New(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	b := New(4, 4, []float64{
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	})

	result := New(4, 4, []float64{
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	})

	assert.Equal(t, result, a.MultiplyMatrix(b))
}

func TestMultiplyTuple(t *testing.T) {
	a := New(4, 4, []float64{
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	})

	b := tuple.New(1, 2, 3, 1)
	result := tuple.New(18.0, 24.0, 33.0, 1.0)

	assert.Equal(t, result, a.MultiplyTuple(b))
}

func TestMultiplyByIdentityMatrix(t *testing.T) {
	a := New(4, 4, []float64{
		0, 1, 2, 4,
		1, 2, 4, 8,
		2, 4, 8, 16,
		4, 8, 16, 32,
	})

	identityMatrix := New(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})

	assert.Equal(t, a, a.MultiplyMatrix(identityMatrix))
}

func TestMultiplyIdentityMatrixByTuple(t *testing.T) {
	a := tuple.New(1, 2, 3, 1)
	identityMatrix := New(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})

	assert.Equal(t, a, identityMatrix.MultiplyTuple(a))
}

func TestTranspose(t *testing.T) {
	a := New(4, 4, []float64{
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8,
	})

	result := New(4, 4, []float64{
		0, 9, 1, 0,
		9, 8, 8, 0,
		3, 0, 5, 5,
		0, 8, 3, 8,
	})

	assert.Equal(t, result, a.Transpose())
}

func TestTransposeIdentityMatrix(t *testing.T) {
	identityMatrix := New(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})

	assert.Equal(t, identityMatrix, identityMatrix.Transpose())
}

func TestSubmatrix(t *testing.T) {
	t.Run("A submatrix of a 3x3 matrix is a 2x2 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			1, 5, 0,
			-3, 2, 7,
			0, 6, 3,
		})

		result := New(2, 2, []float64{
			-3, 2,
			0, 6,
		})

		assert.Equal(t, result, m.Submatrix(0, 2))
	})

	t.Run("A submatrix of a 4x4 matrix is a 3x3 matrix", func(t *testing.T) {
		m := New(4, 4, []float64{
			-6, 1, 1, 6,
			-8, 5, 8, 6,
			-1, 0, 8, -2,
			-7, 1, -1, 1,
		})

		result := New(3, 3, []float64{
			-6, 1, 6,
			-8, 8, 6,
			-7, -1, 1,
		})

		assert.Equal(t, result, m.Submatrix(2, 1))
	})
}

func TestDeterminant(t *testing.T) {
	t.Run("Calculating the determinant of a 2x2 matrix", func(t *testing.T) {
		m := New(2, 2, []float64{
			1, 5,
			-3, 2,
		})

		assert.Equal(t, float64(17), m.Determinant())
	})

	t.Run("Calculating the determinant of a 3x3 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			1, 2, 6,
			-5, 8, -4,
			2, 6, 4,
		})

		assert.Equal(t, float64(56), m.Cofactor(0, 0))
		assert.Equal(t, float64(12), m.Cofactor(0, 1))
		assert.Equal(t, float64(-46), m.Cofactor(0, 2))
		assert.Equal(t, float64(-196), m.Determinant())
	})

	t.Run("Calculating the determinant of a 4x4 matrix", func(t *testing.T) {
		m := New(4, 4, []float64{
			-2, -8, 3, 5,
			-3, 1, 7, 3,
			1, 2, -9, 6,
			-6, 7, 7, -9,
		})

		assert.Equal(t, float64(690), m.Cofactor(0, 0))
		assert.Equal(t, float64(447), m.Cofactor(0, 1))
		assert.Equal(t, float64(210), m.Cofactor(0, 2))
		assert.Equal(t, float64(51), m.Cofactor(0, 3))
		assert.Equal(t, float64(-4071), m.Determinant())
	})
}

func TestMinor(t *testing.T) {
	t.Run("Calulating the minor of a 3x3 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			3, 5, 0,
			2, -1, -7,
			6, -1, 5,
		})

		submatrix := m.Submatrix(1, 0)
		assert.Equal(t, float64(25), submatrix.Determinant())
		assert.Equal(t, float64(25), m.Minor(1, 0))
	})
}

func TestCofactor(t *testing.T) {
	t.Run("Calulating the cofactor of a 3x3 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			3, 5, 0,
			2, -1, -7,
			6, -1, 5,
		})

		assert.Equal(t, float64(-12), m.Minor(0, 0))
		assert.Equal(t, float64(-12), m.Cofactor(0, 0))
		assert.Equal(t, float64(25), m.Minor(1, 0))
		assert.Equal(t, float64(-25), m.Cofactor(1, 0))
	})
}

func TestIsInvertible(t *testing.T) {
	t.Run("Testing an intvertible matrix for invertibility", func(t *testing.T) {
		m := New(4, 4, []float64{
			6, 4, 4, 4,
			5, 5, 7, 6,
			4, -9, 3, -7,
			9, 1, 7, -6,
		})

		assert.Equal(t, float64(-2120), m.Determinant())
		assert.True(t, m.IsInvertible())
	})

	t.Run("Testing a non-intvertible matrix for invertibility", func(t *testing.T) {
		m := New(4, 4, []float64{
			-4, 2, -2, -3,
			9, 6, 2, 6,
			0, -5, 1, -5,
			0, 0, 0, 0,
		})

		assert.Equal(t, float64(0), m.Determinant())
		assert.False(t, m.IsInvertible())
	})
}

func TestInverse(t *testing.T) {
	t.Run("Calculate the inverse of a matrix", func(t *testing.T) {
		m := New(4, 4, []float64{
			-5, 2, 6, -8,
			1, -5, 1, 8,
			7, 7, -6, -7,
			1, -3, 7, 4,
		})

		inverse := New(4, 4, []float64{
			0.21805, 0.45113, 0.24060, -0.04511,
			-0.80827, -1.45677, -0.44361, 0.52068,
			-0.07895, -0.22368, -0.05263, 0.19737,
			-0.52256, -0.81391, -0.30075, 0.30639,
		})

		determinant := m.Determinant()
		assert.Equal(t, float64(532), determinant)

		assert.Equal(t, float64(-160), m.Cofactor(2, 3))
		assert.True(t, util.Approx(inverse.At(3, 2), m.Cofactor(2, 3)/determinant))

		assert.Equal(t, float64(105), m.Cofactor(3, 2))
		assert.True(t, util.Approx(inverse.At(2, 3), m.Cofactor(3, 2)/determinant))

		result, err := m.Inverse()
		assert.Nil(t, err)
		assert.True(t, inverse.Equal(result))
	})

	t.Run("Multiplying a product by its inverse", func(t *testing.T) {
		a := New(4, 4, []float64{
			3, -9, 7, 3,
			3, -8, 2, -9,
			4, 4, 4, 1,
			6, 5, -1, 1,
		})

		b := New(4, 4, []float64{
			8, 2, 2, 2,
			3, -1, 7, 0,
			7, 0, 5, 4,
			6, -2, 0, 5,
		})

		c := a.MultiplyMatrix(b)
		inverse, err := b.Inverse()
		assert.Nil(t, err)

		result := c.MultiplyMatrix(inverse)
		assert.True(t, a.Equal(result))
	})
}
