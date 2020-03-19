package primitives

import (
	"math"

	"github.com/anolson/rtc/color"
	"github.com/anolson/rtc/light"
	"github.com/anolson/rtc/tuple"
)

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

// Lighting shades an objects so it appears three-dimensional
func (m *Material) Lighting(
	light *light.PointLight,
	point *tuple.Tuple,
	eyev *tuple.Tuple,
	normalv *tuple.Tuple,
) *color.Color {

	effectiveColor := color.HadamardProduct(m.Color, light.Intensity)
	lightv := tuple.Subtract(light.Position, point).Normalize()
	ambient := color.Multiply(effectiveColor, m.Ambient)
	black := color.RGB(0, 0, 0)

	var diffuse, specular *color.Color

	lightDotNormal := tuple.Dot(lightv, normalv)
	if lightDotNormal < 0 {
		diffuse = black
		specular = black
	} else {
		diffuse = color.Multiply(effectiveColor, (m.Diffuse * lightDotNormal))

		reflectv := tuple.Reflect(tuple.Negate(lightv), normalv)
		reflectDotEye := tuple.Dot(reflectv, eyev)

		if reflectDotEye <= 0 {
			specular = black
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = color.Multiply(light.Intensity, (m.Specular * factor))
		}
	}

	return color.Add(ambient, color.Add(diffuse, specular))
}
