package ray

import (
	"testing"

	"github.com/anolson/rtc/matrix"
	"github.com/anolson/rtc/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	origin := tuple.Point(1, 2, 3)
	direction := tuple.Vector(4, 5, 6)

	r := New(origin, direction)
	assert.Equal(t, origin, r.Origin)
	assert.Equal(t, direction, r.Direction)
}

func TestPosition(t *testing.T) {
	origin := tuple.Point(2, 3, 4)
	direction := tuple.Vector(1, 0, 0)

	r := New(origin, direction)

	assert.Equal(t, origin, r.Position(0))
	assert.Equal(t, tuple.Point(3, 3, 4), r.Position(1))
	assert.Equal(t, tuple.Point(1, 3, 4), r.Position(-1))
	assert.Equal(t, tuple.Point(4.5, 3, 4), r.Position(2.5))
	assert.Equal(t, direction, r.Direction)
}

func TestTransform(t *testing.T) {
	t.Run("Translating a ray", func(t *testing.T) {
		origin := tuple.Point(1, 2, 3)
		direction := tuple.Vector(0, 1, 0)

		r := New(origin, direction)
		m := matrix.Translation(3, 4, 5)

		r2 := r.Transform(m)

		assert.Equal(t, tuple.Point(4, 6, 8), r2.Origin)
		assert.Equal(t, tuple.Vector(0, 1, 0), r2.Direction)
	})

	t.Run("Translating a ray", func(t *testing.T) {
		origin := tuple.Point(1, 2, 3)
		direction := tuple.Vector(0, 1, 0)

		r := New(origin, direction)
		m := matrix.Scaling(2, 3, 4)

		r2 := r.Transform(m)

		assert.Equal(t, tuple.Point(2, 6, 12), r2.Origin)
		assert.Equal(t, tuple.Vector(0, 3, 0), r2.Direction)
	})
}
