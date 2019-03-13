package primitives

import (
	"math"
	"testing"

	"github.com/anolson/rtc/matrix"
	"github.com/anolson/rtc/ray"
	"github.com/anolson/rtc/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNewSphere(t *testing.T) {
	s := NewSphere()

	assert.Equal(t, matrix.Identity(), s.Transform)
}

func TestSetTransformSphere(t *testing.T) {
	s := NewSphere()
	transform := matrix.Translation(2, 3, 4)

	s.SetTransform(transform)

	assert.Equal(t, transform, s.Transform)
}

func TestIntersection(t *testing.T) {
	t.Run("A ray intsects a sphere at two points", func(t *testing.T) {
		origin := tuple.Point(0, 0, -5)
		direction := tuple.Vector(0, 0, 1)
		r := ray.New(origin, direction)
		s := NewSphere()

		intersection := s.Intersect(r)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, 4.0, intersection[0].t)
		assert.Equal(t, 6.0, intersection[1].t)
	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		origin := tuple.Point(0, 1, -5)
		direction := tuple.Vector(0, 0, 1)
		r := ray.New(origin, direction)
		s := NewSphere()

		intersection := s.Intersect(r)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, 5.0, intersection[0].t)
		assert.Equal(t, 5.0, intersection[1].t)
	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		origin := tuple.Point(0, 2, -5)
		direction := tuple.Vector(0, 0, 1)
		r := ray.New(origin, direction)
		s := NewSphere()

		intersection := s.Intersect(r)

		assert.Equal(t, 0, len(intersection))
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		origin := tuple.Point(0, 0, 0)
		direction := tuple.Vector(0, 0, 1)
		r := ray.New(origin, direction)
		s := NewSphere()

		intersection := s.Intersect(r)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, -1.0, intersection[0].t)
		assert.Equal(t, 1.0, intersection[1].t)
	})

	t.Run("A sphere is behind a ray", func(t *testing.T) {
		origin := tuple.Point(0, 0, 5)
		direction := tuple.Vector(0, 0, 1)
		r := ray.New(origin, direction)
		s := NewSphere()

		intersection := s.Intersect(r)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, -6.0, intersection[0].t)
		assert.Equal(t, -4.0, intersection[1].t)
	})

	t.Run("Intersect sets the object on the intersection", func(t *testing.T) {
		origin := tuple.Point(0, 0, 5)
		direction := tuple.Vector(0, 0, 1)
		r := ray.New(origin, direction)
		s := NewSphere()

		intersection := s.Intersect(r)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, s, intersection[0].object)
		assert.Equal(t, s, intersection[1].object)
	})

	t.Run("Intersecting a scaled sphere a ray", func(t *testing.T) {
		origin := tuple.Point(0, 0, -5)
		direction := tuple.Vector(0, 0, 1)
		r := ray.New(origin, direction)
		s := NewSphere()
		s.SetTransform(matrix.Scaling(2, 2, 2))

		intersection := s.Intersect(r)

		assert.Equal(t, 2, len(intersection))
		assert.Equal(t, float64(3), intersection[0].t)
		assert.Equal(t, float64(7), intersection[1].t)
	})

	t.Run("Intersecting a translated sphere a ray", func(t *testing.T) {
		origin := tuple.Point(0, 0, -5)
		direction := tuple.Vector(0, 0, 1)
		r := ray.New(origin, direction)
		s := NewSphere()
		s.SetTransform(matrix.Translation(5, 0, 0))

		intersection := s.Intersect(r)

		assert.Equal(t, 0, len(intersection))
	})
}

func TestNormalAt(t *testing.T) {
	t.Run("The normal on a sphere at a point on the x axis", func(t *testing.T) {
		s := NewSphere()
		p := tuple.Point(1, 0, 0)

		n := s.NormalAt(p)
		assert.Equal(t, tuple.Vector(1, 0, 0), n)
	})

	t.Run("The normal on a sphere at a point on the y axis", func(t *testing.T) {
		s := NewSphere()
		p := tuple.Point(0, 1, 0)

		n := s.NormalAt(p)
		assert.Equal(t, tuple.Vector(0, 1, 0), n)

	})

	t.Run("The normal on a sphere at a point on the z axis", func(t *testing.T) {
		s := NewSphere()
		p := tuple.Point(0, 0, 1)

		n := s.NormalAt(p)
		assert.Equal(t, tuple.Vector(0, 0, 1), n)
	})

	t.Run("The normal on a sphere at a nonaxial point", func(t *testing.T) {
		s := NewSphere()
		p := tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)

		n := s.NormalAt(p)
		assert.Equal(t, tuple.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), n)
	})

	t.Run("The normal is a normalized vector", func(t *testing.T) {
		s := NewSphere()
		p := tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)

		n := s.NormalAt(p)
		assert.Equal(t, n.Normalize(), n)
	})
}
