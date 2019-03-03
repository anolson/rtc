package primitives

import "github.com/anolson/rtc/ray"

type Object interface {
	Intersect(r *ray.Ray) []*Intersection
}
