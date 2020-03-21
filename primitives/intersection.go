package primitives

import "sort"

// Intersection aggregates the object and t values
type Intersection struct {
	T      float64
	Object Object
}

// NewIntersection returns a new Intersection object
func NewIntersection(t float64, object Object) *Intersection {
	return &Intersection{
		T:      t,
		Object: object,
	}
}

// Hit identifies which intersection is visible from the ray's origin
func Hit(intersections []*Intersection) *Intersection {
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].T < intersections[j].T
	})

	for _, intersection := range intersections {
		if intersection.T > 0 {
			return intersection
		}
	}

	return nil
}
