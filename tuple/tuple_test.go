package tuple

import (
	"math"
	"testing"

	"github.com/anolson/rtc/util"
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

	point := Point(x, y, z)
	assert.Equal(t, tuple, point)
	assert.True(t, point.isPoint())
	assert.False(t, point.isVector())

}

func TestNewVector(t *testing.T) {
	x, y, z, w := 4.3, -4.2, 3.1, 0.0
	tuple := &Tuple{X: x, Y: y, Z: z, W: w}

	vector := Vector(x, y, z)
	assert.Equal(t, tuple, vector)
	assert.True(t, tuple.isVector())
	assert.False(t, tuple.isPoint())
}

func TestAdd(t *testing.T) {
	t1 := &Tuple{X: 3, Y: -2, Z: 5, W: 1}
	t2 := &Tuple{X: -2, Y: 3, Z: 1, W: 0}

	result := Add(t1, t2)
	assert.Equal(t, &Tuple{X: 1, Y: 1, Z: 6, W: 1}, result)
}

func TestSubtract(t *testing.T) {
	t.Run("subtract two points", func(t *testing.T) {
		p1 := Point(3, 2, 1)
		p2 := Point(5, 6, 7)

		result := Subtract(p1, p2)
		assert.Equal(t, Vector(-2, -4, -6), result)
	})

	t.Run("subtract a vector from a point", func(t *testing.T) {
		p := Point(3, 2, 1)
		v := Vector(5, 6, 7)

		result := Subtract(p, v)
		assert.Equal(t, Point(-2, -4, -6), result)
	})

	t.Run("subtract two vectors", func(t *testing.T) {
		v1 := Vector(3, 2, 1)
		v2 := Vector(5, 6, 7)

		result := Subtract(v1, v2)
		assert.Equal(t, Vector(-2, -4, -6), result)
	})
}

func TestNegate(t *testing.T) {
	t1 := &Tuple{X: 1, Y: -2, Z: 3, W: 1}

	result := Negate(t1)
	assert.Equal(t, &Tuple{X: -1, Y: 2, Z: -3, W: -1}, result)
}

func TestMultiply(t *testing.T) {
	t1 := &Tuple{X: 1, Y: -2, Z: 3, W: 1}

	result := Multiply(t1, 3.5)
	assert.Equal(t, &Tuple{X: 3.5, Y: -7, Z: 10.5, W: 3.5}, result)
}

func TestDivide(t *testing.T) {
	t1 := &Tuple{X: 1, Y: -2, Z: 3, W: 1}

	result := Divide(t1, 2)
	assert.Equal(t, &Tuple{X: 0.5, Y: -1, Z: 1.5, W: 0.5}, result)
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		description string
		vector      *Tuple
		magnitude   float64
	}{
		{"Magnitude of vector(1, 0, 0)", Vector(1, 0, 0), float64(1)},
		{"Magnitude of vector(0, 1, 0)", Vector(0, 1, 0), float64(1)},
		{"Magnitude of vector(0, 0, 1)", Vector(0, 0, 1), float64(1)},
		{"Magnitude of vector(1, 2, 3)", Vector(1, 2, 3), math.Sqrt(14)},
		{"Magnitude of vector(-1, -2, -3)", Vector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, test := range tests {
		assert.Equal(t, test.magnitude, test.vector.Magnitude())
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		description string
		vector      *Tuple
		normalized  *Tuple
	}{
		{"Normalizing vector(4, 0, 0)", Vector(4, 0, 0), Vector(1, 0, 0)},
		{"Normalizing vector(1, 2, 3)", Vector(1, 2, 3), Vector(0.26726, 0.53452, 0.80178)},
	}

	for _, test := range tests {
		normalized := test.vector.Normalize()
		assert.True(t, util.Approx(normalized.X, test.normalized.X))
		assert.True(t, util.Approx(normalized.Y, test.normalized.Y))
		assert.True(t, util.Approx(normalized.Z, test.normalized.Z))
	}

	t.Run("Magnitude of a normalized vector is 1", func(t *testing.T) {
		v := Vector(1, 2, 3)
		normalized := v.Normalize()

		assert.Equal(t, float64(1), normalized.Magnitude())
	})
}

func TestDot(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)

	assert.Equal(t, float64(20), Dot(v1, v2))
}

func TestCross(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)

	assert.Equal(t, Vector(-1, 2, -1), Cross(v1, v2))
	assert.Equal(t, Vector(1, -2, 1), Cross(v2, v1))
}
