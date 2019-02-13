package matrix

import (
	"testing"

	"github.com/anolson/rtc/util"
	"github.com/stretchr/testify/assert"
)

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

		determinant := Determinant(m)
		assert.Equal(t, float64(532), determinant)

		assert.Equal(t, float64(-160), Cofactor(m, 2, 3))
		assert.True(t, util.Approx(inverse.At(3, 2), Cofactor(m, 2, 3)/determinant))

		assert.Equal(t, float64(105), Cofactor(m, 3, 2))
		assert.True(t, util.Approx(inverse.At(2, 3), Cofactor(m, 3, 2)/determinant))

		result, err := Inverse(m)
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

		c := Multiply(a, b)
		inverse, err := Inverse(b)
		assert.Nil(t, err)

		result := Multiply(c, inverse)
		assert.True(t, a.Equal(result))
	})
}

func TestDeterminant(t *testing.T) {
	t.Run("Calculating the determinant of a 2x2 matrix", func(t *testing.T) {
		m := New(2, 2, []float64{
			1, 5,
			-3, 2,
		})

		assert.Equal(t, float64(17), Determinant(m))
	})

	t.Run("Calculating the determinant of a 3x3 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			1, 2, 6,
			-5, 8, -4,
			2, 6, 4,
		})

		assert.Equal(t, float64(56), Cofactor(m, 0, 0))
		assert.Equal(t, float64(12), Cofactor(m, 0, 1))
		assert.Equal(t, float64(-46), Cofactor(m, 0, 2))
		assert.Equal(t, float64(-196), Determinant(m))
	})

	t.Run("Calculating the determinant of a 4x4 matrix", func(t *testing.T) {
		m := New(4, 4, []float64{
			-2, -8, 3, 5,
			-3, 1, 7, 3,
			1, 2, -9, 6,
			-6, 7, 7, -9,
		})

		assert.Equal(t, float64(690), Cofactor(m, 0, 0))
		assert.Equal(t, float64(447), Cofactor(m, 0, 1))
		assert.Equal(t, float64(210), Cofactor(m, 0, 2))
		assert.Equal(t, float64(51), Cofactor(m, 0, 3))
		assert.Equal(t, float64(-4071), Determinant(m))
	})
}

func TestMinor(t *testing.T) {
	t.Run("Calulating the minor of a 3x3 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			3, 5, 0,
			2, -1, -7,
			6, -1, 5,
		})

		submatrix := Submatrix(m, 1, 0)
		assert.Equal(t, float64(25), Determinant(submatrix))
		assert.Equal(t, float64(25), Minor(m, 1, 0))
	})
}

func TestCofactor(t *testing.T) {
	t.Run("Calulating the cofactor of a 3x3 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			3, 5, 0,
			2, -1, -7,
			6, -1, 5,
		})

		assert.Equal(t, float64(-12), Minor(m, 0, 0))
		assert.Equal(t, float64(-12), Cofactor(m, 0, 0))
		assert.Equal(t, float64(25), Minor(m, 1, 0))
		assert.Equal(t, float64(-25), Cofactor(m, 1, 0))
	})
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

		assert.Equal(t, result, Submatrix(m, 0, 2))
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

		assert.Equal(t, result, Submatrix(m, 2, 1))
	})
}
