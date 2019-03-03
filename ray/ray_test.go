package ray

import (
	"testing"

	"github.com/anolson/rtc/primitives"
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

func TestIntersection(t *testing.T) {
	t.Run("A ray intsects a sphere at two points", func(t *testing.T) {
		origin := tuple.Point(0, 0, -5)
		direction := tuple.Vector(0, 0, 1)

		r := New(origin, direction)
		s := primitives.NewSphere()
		intersection := r.Intersect(s)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, 4.0, intersection[0])
		assert.Equal(t, 6.0, intersection[1])
	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		origin := tuple.Point(0, 1, -5)
		direction := tuple.Vector(0, 0, 1)

		r := New(origin, direction)
		s := primitives.NewSphere()
		intersection := r.Intersect(s)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, 5.0, intersection[0])
		assert.Equal(t, 5.0, intersection[1])
	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		origin := tuple.Point(0, 2, -5)
		direction := tuple.Vector(0, 0, 1)

		r := New(origin, direction)
		s := primitives.NewSphere()
		intersection := r.Intersect(s)

		assert.Equal(t, 0, len(intersection))
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		origin := tuple.Point(0, 0, 0)
		direction := tuple.Vector(0, 0, 1)

		r := New(origin, direction)
		s := primitives.NewSphere()
		intersection := r.Intersect(s)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, -1.0, intersection[0])
		assert.Equal(t, 1.0, intersection[1])
	})

	t.Run("A sphere is behind a ray", func(t *testing.T) {
		origin := tuple.Point(0, 0, 5)
		direction := tuple.Vector(0, 0, 1)

		r := New(origin, direction)
		s := primitives.NewSphere()
		intersection := r.Intersect(s)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, -6.0, intersection[0])
		assert.Equal(t, -4.0, intersection[1])
	})
}
