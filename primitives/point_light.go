package primitives

import (
	"github.com/anolson/rtc/color"
	"github.com/anolson/rtc/tuple"
)

// PointLight represents a light source
type PointLight struct {
	Position  *tuple.Tuple
	Intensity *color.Color
}

// NewPointLight returns a new point light source object
func NewPointLight(position *tuple.Tuple, intensity *color.Color) *PointLight {
	return &PointLight{
		Position:  position,
		Intensity: intensity,
	}
}
