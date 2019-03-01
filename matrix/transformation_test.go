package matrix

import (
	"math"
	"testing"

	"github.com/anolson/rtc/tuple"
	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	t.Run("Applying a translation matrix to a point", func(t *testing.T) {
		transform := Translation(5, -3, 2)
		p := tuple.Point(-3, 4, 5)

		assert.Equal(t, tuple.Point(2, 1, 7), Transform(transform, p))
	})

	t.Run("Applying a translation matrix to a vector - has no effect", func(t *testing.T) {
		transform := Translation(5, -3, 2)
		v := tuple.Vector(-3, 4, 5)

		assert.Equal(t, v, Transform(transform, v))
	})

	t.Run("Applying the inverse of a translation matrix", func(t *testing.T) {
		transform := Translation(5, -3, 2)
		inverse, err := Inverse(transform)
		assert.Nil(t, err)
		p := tuple.Point(-3, 4, 5)

		assert.Equal(t, tuple.Point(-8, 7, 3), Transform(inverse, p))
	})
}

func TestScaling(t *testing.T) {
	t.Run("Applying a scaling matrix to a point", func(t *testing.T) {
		transform := Scaling(2, 3, 4)
		p := tuple.Point(-4, 6, 8)

		assert.Equal(t, tuple.Point(-8, 18, 32), Transform(transform, p))
	})

	t.Run("Applying by a scaling matrix to a vector", func(t *testing.T) {
		transform := Scaling(2, 3, 4)
		v := tuple.Vector(-4, 6, 8)

		assert.Equal(t, tuple.Vector(-8, 18, 32), Transform(transform, v))
	})

	t.Run("Applying the inverse of a scaling matrix", func(t *testing.T) {
		transform := Scaling(2, 3, 4)
		inverse, err := Inverse(transform)
		assert.Nil(t, err)
		v := tuple.Vector(-4, 6, 8)

		assert.Equal(t, tuple.Vector(-2, 2, 2), Transform(inverse, v))
	})

	t.Run("Reflection is scaling by a negative value", func(t *testing.T) {
		transform := Scaling(-1, 1, 1)
		p := tuple.Point(2, 3, 4)

		assert.Equal(t, tuple.Point(-2, 3, 4), Transform(transform, p))
	})
}

func TestRotationX(t *testing.T) {
	t.Run("Rotating a point around the x axis", func(t *testing.T) {
		halfQuarter := RotationX(math.Pi / 4)
		fullQuarter := RotationX(math.Pi / 2)
		p := tuple.Point(0, 1, 0)

		assert.True(t, tuple.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2).Equal(Transform(halfQuarter, p)))
		assert.True(t, tuple.Point(0, 0, 1).Equal(Transform(fullQuarter, p)))
	})

	t.Run("Rotating a point around the x axis - in the opposite direction", func(t *testing.T) {
		halfQuarter := RotationX(math.Pi / 4)
		inverse, err := Inverse(halfQuarter)
		assert.Nil(t, err)
		p := tuple.Point(0, 1, 0)

		assert.True(t, tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2).Equal(Transform(inverse, p)))
	})
}

func TestRotationY(t *testing.T) {
	t.Run("Rotating a point around the y axis", func(t *testing.T) {
		halfQuarter := RotationY(math.Pi / 4)
		fullQuarter := RotationY(math.Pi / 2)
		p := tuple.Point(0, 0, 1)

		assert.True(t, tuple.Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2).Equal(Transform(halfQuarter, p)))
		assert.True(t, tuple.Point(1, 0, 0).Equal(Transform(fullQuarter, p)))
	})

	t.Run("Rotating a point around the y axis - in the opposite direction", func(t *testing.T) {
		halfQuarter := RotationY(math.Pi / 4)
		inverse, err := Inverse(halfQuarter)
		assert.Nil(t, err)
		p := tuple.Point(0, 0, 1)

		assert.True(t, tuple.Point(-math.Sqrt(2)/2, 0, math.Sqrt(2)/2).Equal(Transform(inverse, p)))
	})
}

func TestRotationZ(t *testing.T) {
	t.Run("Rotating a point around the z axis", func(t *testing.T) {
		halfQuarter := RotationZ(math.Pi / 4)
		fullQuarter := RotationZ(math.Pi / 2)
		p := tuple.Point(0, 1, 0)

		assert.True(t, tuple.Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0).Equal(Transform(halfQuarter, p)))
		assert.True(t, tuple.Point(-1, 0, 0).Equal(Transform(fullQuarter, p)))
	})

	t.Run("Rotating a point around the z axis - in the opposite direction", func(t *testing.T) {
		halfQuarter := RotationZ(math.Pi / 4)
		inverse, err := Inverse(halfQuarter)
		assert.Nil(t, err)
		p := tuple.Point(0, 1, 0)

		assert.True(t, tuple.Point(math.Sqrt(2)/2, math.Sqrt(2)/2, 0).Equal(Transform(inverse, p)))
	})
}

func TestShearing(t *testing.T) {
	t.Run("A shearing transformation moves x in proportion to y", func(t *testing.T) {
		transform := Shearing(1, 0, 0, 0, 0, 0)
		p := tuple.Point(2, 3, 4)

		assert.Equal(t, tuple.Point(5, 3, 4), Transform(transform, p))
	})

	t.Run("A shearing transformation moves x in proportion to z", func(t *testing.T) {
		transform := Shearing(0, 1, 0, 0, 0, 0)
		p := tuple.Point(2, 3, 4)

		assert.Equal(t, tuple.Point(6, 3, 4), Transform(transform, p))
	})

	t.Run("A shearing transformation moves y in proportion to x", func(t *testing.T) {
		transform := Shearing(0, 0, 1, 0, 0, 0)
		p := tuple.Point(2, 3, 4)

		assert.Equal(t, tuple.Point(2, 5, 4), Transform(transform, p))
	})

	t.Run("A shearing transformation moves y in proportion to z", func(t *testing.T) {
		transform := Shearing(0, 0, 0, 1, 0, 0)
		p := tuple.Point(2, 3, 4)

		assert.Equal(t, tuple.Point(2, 7, 4), Transform(transform, p))
	})

	t.Run("A shearing transformation moves z in proportion to x", func(t *testing.T) {
		transform := Shearing(0, 0, 0, 0, 1, 0)
		p := tuple.Point(2, 3, 4)

		assert.Equal(t, tuple.Point(2, 3, 6), Transform(transform, p))
	})

	t.Run("A shearing transformation moves z in proportion to y", func(t *testing.T) {
		transform := Shearing(0, 0, 0, 0, 0, 1)
		p := tuple.Point(2, 3, 4)

		assert.Equal(t, tuple.Point(2, 3, 7), Transform(transform, p))
	})
}

func TestChainingTransformations(t *testing.T) {
	t.Run("Individual transformations are applied in sequence", func(t *testing.T) {
		p := tuple.Point(1, 0, 1)
		a := RotationX(math.Pi / 2)
		b := Scaling(5, 5, 5)
		c := Translation(10, 5, 7)

		p2 := Transform(a, p)
		assert.True(t, tuple.Point(1, -1, 0).Equal(p2))

		p3 := Transform(b, p2)
		assert.True(t, tuple.Point(5, -5, 0).Equal(p3))

		p4 := Transform(c, p3)
		assert.True(t, tuple.Point(15, 0, 7).Equal(p4))
	})

	t.Run("Chained transformations must be applied in reverse order", func(t *testing.T) {
		p := tuple.Point(1, 0, 1)
		a := RotationX(math.Pi / 2)
		b := Scaling(5, 5, 5)
		c := Translation(10, 5, 7)

		p2 := Chain(p, a, b, c)
		assert.True(t, tuple.Point(15, 0, 7).Equal(p2))
	})
}
