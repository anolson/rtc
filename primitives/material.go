package primitives

import "github.com/anolson/rtc/color"

const (
	defaultAmbient   = 0.1
	defaultDiffuse   = 0.9
	defaultSpecular  = 0.9
	defaultShininess = 200.1
)

// Material encapsulates a color the four attributes of the Phong reflection model
type Material struct {
	Color     *color.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

// DefaultMaterial returns the default Material
func DefaultMaterial() *Material {
	return &Material{
		Color:     color.RGB(1, 1, 1),
		Ambient:   defaultAmbient,
		Diffuse:   defaultDiffuse,
		Specular:  defaultSpecular,
		Shininess: defaultShininess,
	}
}
