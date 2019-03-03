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
