package light

import (
	"github.com/anolson/rtc/color"
	"github.com/anolson/rtc/tuple"
)

// Light represents a light source
type Light struct {
	Position  *tuple.Tuple
	Intensity *color.Color
}

// NewPoint returns a new point light source object
func NewPoint(position *tuple.Tuple, intensity *color.Color) *Light {
	return &Light{
		Position:  position,
		Intensity: intensity,
	}
}
