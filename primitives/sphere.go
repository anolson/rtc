package primitives

import (
	"math"

	"github.com/anolson/rtc/ray"
	"github.com/anolson/rtc/tuple"
)

// Sphere represents a 3D Spherical shape
type Sphere struct{}

// NewSphere returns a new Sphere object
func NewSphere() *Sphere {
	return &Sphere{}
}

// Intersect returns the intersection of a ray through sphere
func (s *Sphere) Intersect(r *ray.Ray) []*Intersection {
	sphereToRay := tuple.Subtract(r.Origin, tuple.Point(0, 0, 0))

	a := tuple.Dot(r.Direction, r.Direction)
	b := 2 * tuple.Dot(r.Direction, sphereToRay)
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
