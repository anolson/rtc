package primitives

import "sort"

// Intersection aggregates the object and t values
type Intersection struct {
	t      float64
	object Object
}

// NewIntersection returns a new Intersection object
func NewIntersection(t float64, object Object) *Intersection {
	return &Intersection{
		t:      t,
		object: object,
	}
}

// Hit identifies which intersection is visible from the ray's origin
func Hit(intersections []*Intersection) *Intersection {
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].t < intersections[j].t
	})

	for _, intersection := range intersections {
		if intersection.t > 0 {
			return intersection
		}
	}

	return nil
}
