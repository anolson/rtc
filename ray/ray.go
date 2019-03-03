package ray

import (
	"math"

	"github.com/anolson/rtc/primitives"
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

// Intersect returns the intersection
func (r *Ray) Intersect(s *primitives.Sphere) []float64 {
	sphereToRay := tuple.Subtract(r.Origin, tuple.Point(0, 0, 0))

	a := tuple.Dot(r.Direction, r.Direction)
	b := 2 * tuple.Dot(r.Direction, sphereToRay)
	c := tuple.Dot(sphereToRay, sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return []float64{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return []float64{t1, t2}
}
