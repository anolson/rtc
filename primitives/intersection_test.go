package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntersection(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, s)

	assert.Equal(t, s, i.object)
	assert.Equal(t, 3.5, i.t)
}

func TestIntersectionAggregation(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)

	intersections := []*Intersection{i1, i2}

	assert.Equal(t, float64(1), intersections[0].t)
	assert.Equal(t, float64(2), intersections[1].t)
}

func TestHit(t *testing.T) {
	t.Run("The hit, when all intersections have positive t values", func(t *testing.T) {
		s := NewSphere()
		i1 := NewIntersection(1, s)
		i2 := NewIntersection(2, s)

		intersections := []*Intersection{i2, i1}
		assert.Equal(t, i1, Hit(intersections))
	})

	t.Run("The hit, when some intersections have negative t values", func(t *testing.T) {
		s := NewSphere()
		i1 := NewIntersection(-1, s)
		i2 := NewIntersection(1, s)

		intersections := []*Intersection{i2, i1}
		assert.Equal(t, i2, Hit(intersections))
	})

	t.Run("The hit, when all intersections have negative t values", func(t *testing.T) {
		s := NewSphere()
		i1 := NewIntersection(-2, s)
		i2 := NewIntersection(-1, s)

		intersections := []*Intersection{i2, i1}
		assert.Nil(t, Hit(intersections))
	})

	t.Run("Is always the non-negative t value", func(t *testing.T) {
		s := NewSphere()
		i1 := NewIntersection(5, s)
		i2 := NewIntersection(7, s)
		i3 := NewIntersection(-3, s)
		i4 := NewIntersection(2, s)

		intersections := []*Intersection{i1, i2, i3, i4}
		assert.Equal(t, i4, Hit(intersections))
	})
}
