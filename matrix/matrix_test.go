package matrix

import (
	"testing"

	"github.com/stvp/assert"
)

func TestNew(t *testing.T) {
	t.Run("Creating a 2x2 matrix", func(t *testing.T) {
		m := New(2, 2, []float64{
			-3.0, 5.0,
			1.0, -2.0,
		})

		assert.Equal(t, -3.0, m.At(0, 0))
		assert.Equal(t, 5.0, m.At(0, 1))
		assert.Equal(t, 1.0, m.At(1, 0))
		assert.Equal(t, -2.0, m.At(1, 1))
	})

	t.Run("Creating a 3x3 matrix", func(t *testing.T) {
		m := New(3, 3, []float64{
			-3.0, 5.0, 0.0,
			1.0, -2.0, -7.0,
			0.0, 1.0, 1.0,
		})

		assert.Equal(t, -3.0, m.At(0, 0))
		assert.Equal(t, -2.0, m.At(1, 1))
		assert.Equal(t, 1.0, m.At(2, 2))
	})

	t.Run("Creating a 4x4 matrix", func(t *testing.T) {
		m := New(4, 4, []float64{
			1.0, 2.0, 3.0, 4.0,
			5.5, 6.5, 7.5, 8.5,
			9.0, 10.0, 11.0, 12.0,
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
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	})

	other := New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.000001,
	})

	otherNotEqual := New(4, 4, []float64{
		2.0, 3.0, 4.0, 5.0,
		6.0, 7.0, 8.0, 9.0,
		8.0, 7.0, 6.0, 5.0,
		4.0, 3.0, 2.0, 1.0,
	})

	assert.True(t, m.Equal(other))
	assert.False(t, m.Equal(otherNotEqual))
}

func TestRow(t *testing.T) {
	m := New(3, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	})

	assert.Equal(t, []float64{1.0, 2.0, 3.0, 4.0}, m.Row(0))
	assert.Equal(t, []float64{5.0, 6.0, 7.0, 8.0}, m.Row(1))
	assert.Equal(t, []float64{9.0, 8.0, 7.0, 6.0}, m.Row(2))
	assert.Equal(t, []float64{5.0, 4.0, 3.0, 2.0}, m.Row(3))
}

func TestCol(t *testing.T) {
	m := New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	})

	assert.Equal(t, []float64{1.0, 5.0, 9.0, 5.0}, m.Col(0))
	assert.Equal(t, []float64{2.0, 6.0, 8.0, 4.0}, m.Col(1))
	assert.Equal(t, []float64{3.0, 7.0, 7.0, 3.0}, m.Col(2))
	assert.Equal(t, []float64{4.0, 8.0, 6.0, 2.0}, m.Col(3))
}

func TestMultiply(t *testing.T) {
	a := New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	})

	b := New(4, 4, []float64{
		-2.0, 1.0, 2.0, 3.0,
		3.0, 2.0, 1.0, -1.0,
		4.0, 3.0, 6.0, 5.0,
		1.0, 2.0, 7.0, 8.0,
	})

	result := New(4, 4, []float64{
		20.0, 22.0, 50.0, 48.0,
		44.0, 54.0, 114.0, 108.0,
		40.0, 58.0, 110.0, 102.0,
		16.0, 26.0, 46.0, 42.0,
	})

	assert.Equal(t, result, a.Multiply(b))
}
