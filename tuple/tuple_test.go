package tuple

import (
	"testing"

	"github.com/stvp/assert"
)

func TestIsPoint(t *testing.T) {
	t.Run("returns true for tuples that are points (W=1)", func(t *testing.T) {
		tuple := &Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}
		assert.True(t, tuple.isPoint())
	})

	t.Run("returns false for tuples that are vectors (W=0)", func(t *testing.T) {
		tuple := &Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}
		assert.False(t, tuple.isPoint())
	})
}

func TestIsVector(t *testing.T) {
	t.Run("returns true for tuples that are vectors (W=0)", func(t *testing.T) {
		tuple := &Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}
		assert.True(t, tuple.isVector())
	})

	t.Run("returns false for tuples that are points (W=1)", func(t *testing.T) {
		tuple := &Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}
		assert.False(t, tuple.isVector())
	})
}

func TestNewPoint(t *testing.T) {
	x, y, z, w := 4.3, -4.2, 3.1, 1.0
	tuple := &Tuple{X: x, Y: y, Z: z, W: w}

	point := NewPoint(x, y, z)
	assert.Equal(t, tuple, point)
	assert.True(t, point.isPoint())
	assert.False(t, point.isVector())

}

func TestNewVector(t *testing.T) {
	x, y, z, w := 4.3, -4.2, 3.1, 0.0
	tuple := &Tuple{X: x, Y: y, Z: z, W: w}

	vector := NewVector(x, y, z)
	assert.Equal(t, tuple, vector)
	assert.True(t, tuple.isVector())
	assert.False(t, tuple.isPoint())
}

func TestAdd(t *testing.T) {
	t1 := &Tuple{X: 3, Y: -2, Z: 5, W: 1}
	t2 := &Tuple{X: -2, Y: 3, Z: 1, W: 0}

	result := t1.Add(t2)
	assert.Equal(t, &Tuple{X: 1, Y: 1, Z: 6, W: 1}, result)
}
