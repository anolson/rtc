package primitives

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
