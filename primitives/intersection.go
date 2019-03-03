package primitives

type Intersection struct {
	t      float64
	object *Sphere
}

func NewIntersection(t float64, object *Sphere) *Intersection {
	return &Intersection{
		t:      t,
		object: object,
	}
}
