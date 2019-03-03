package ray

import (
	"github.com/anolson/rtc/tuple"
)

// Ray represnts a ray or line for our ray tracer
type Ray struct {
	Origin    *tuple.Tuple // point
	Direction *tuple.Tuple // vector
}

// New returns a new Ray object
func New(origin, direction *tuple.Tuple) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}

// Position computes the point at the given distance along the ray
func (r *Ray) Position(t float64) *tuple.Tuple {
	return tuple.Add(r.Origin, tuple.Multiply(r.Direction, t))
}
