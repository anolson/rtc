package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
