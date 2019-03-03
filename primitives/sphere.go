package primitives

import (
	"math"

	"github.com/anolson/rtc/matrix"
	"github.com/anolson/rtc/ray"
	"github.com/anolson/rtc/tuple"
)

// Sphere represents a 3D Spherical shape
type Sphere struct {
	Transform *matrix.Matrix
}

// NewSphere returns a new Sphere object
func NewSphere() *Sphere {
	return &Sphere{
		Transform: matrix.Identity(),
	}
}

// SetTransform sets the Transform on a Sphere
func (s *Sphere) SetTransform(transform *matrix.Matrix) {
	s.Transform = transform
}

// Intersect returns the intersection of a ray through sphere
func (s *Sphere) Intersect(r *ray.Ray) []*Intersection {
	inverseTransform, _ := matrix.Inverse(s.Transform)
	r2 := r.Transform(inverseTransform)

	sphereToRay := tuple.Subtract(r2.Origin, tuple.Point(0, 0, 0))

	a := tuple.Dot(r2.Direction, r2.Direction)
	b := 2 * tuple.Dot(r2.Direction, sphereToRay)
	c := tuple.Dot(sphereToRay, sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return []*Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return []*Intersection{
		NewIntersection(t1, s),
		NewIntersection(t2, s),
	}
}
