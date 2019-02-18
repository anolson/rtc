package matrix

import (
	"testing"

	"github.com/anolson/rtc/tuple"
	"github.com/stretchr/testify/assert"
)

func TestTranslationMatrix(t *testing.T) {
	t.Run("Multiplying by a translation matrix", func(t *testing.T) {
		transform := Translation(5, -3, 2)
		p := tuple.Point(-3, 4, 5)

		assert.Equal(t, tuple.Point(2, 1, 7), MultiplyByTuple(transform, p))
	})

	t.Run("Multiplying by inverse of a translation matrix", func(t *testing.T) {
		transform := Translation(5, -3, 2)
		inverse, err := Inverse(transform)
		assert.Nil(t, err)
		p := tuple.Point(-3, 4, 5)

		assert.Equal(t, tuple.Point(-8, 7, 3), MultiplyByTuple(inverse, p))
	})

	t.Run("Translation does not affect vector", func(t *testing.T) {
		transform := Translation(5, -3, 2)
		v := tuple.Vector(-3, 4, 5)

		assert.Equal(t, v, MultiplyByTuple(transform, v))
	})
}

func TestScalingMatrix(t *testing.T) {
	t.Run("Applying a scaling matrix to a point", func(t *testing.T) {
		transform := Scaling(2, 3, 4)
		p := tuple.Point(-4, 6, 8)

		assert.Equal(t, tuple.Point(-8, 18, 32), MultiplyByTuple(transform, p))
	})

	t.Run("Applying by a scaling matrix to a vector", func(t *testing.T) {
		transform := Scaling(2, 3, 4)
		v := tuple.Vector(-4, 6, 8)

		assert.Equal(t, tuple.Vector(-8, 18, 32), MultiplyByTuple(transform, v))
	})

	t.Run("Applying the inverse of a scaling matrix", func(t *testing.T) {
		transform := Scaling(2, 3, 4)
		inverse, err := Inverse(transform)
		assert.Nil(t, err)
		v := tuple.Vector(-4, 6, 8)

		assert.Equal(t, tuple.Vector(-2, 2, 2), MultiplyByTuple(inverse, v))
	})
}
