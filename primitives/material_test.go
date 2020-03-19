package primitives

import (
	"math"
	"testing"

	"github.com/anolson/rtc/color"
	"github.com/anolson/rtc/light"
	"github.com/anolson/rtc/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNewMaterial(t *testing.T) {
	t.Run("The default material", func(t *testing.T) {
		m := DefaultMaterial()

		assert.Equal(t, color.RGB(1, 1, 1), m.Color)
		assert.Equal(t, defaultAmbient, m.Ambient)
		assert.Equal(t, defaultDiffuse, m.Diffuse)
		assert.Equal(t, defaultSpecular, m.Specular)
		assert.Equal(t, defaultShininess, m.Shininess)
	})
}

func TestLighting(t *testing.T) {
	m := DefaultMaterial()
	position := tuple.Point(0, 0, 0)

	t.Run("Lighting with the eye between the light and the surface", func(t *testing.T) {
		eyev := tuple.Vector(0, 0, -1)
		normalv := tuple.Vector(0, 0, -1)
		light := light.NewPointLight(tuple.Point(0, 0, -10), color.RGB(1, 1, 1))

		result := m.Lighting(light, position, eyev, normalv)

		assert.True(t, result.Equal(color.RGB(1.9, 1.9, 1.9)))
	})

	t.Run("Lighting with the eye between light and surface, eye offset 45°", func(t *testing.T) {
		eyev := tuple.Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
		normalv := tuple.Vector(0, 0, -1)
		light := light.NewPointLight(tuple.Point(0, 0, -10), color.RGB(1, 1, 1))

		result := m.Lighting(light, position, eyev, normalv)

		assert.True(t, result.Equal(color.RGB(1.0, 1.0, 1.0)))
	})

	t.Run("Lighting with the eye opposite surface, eye offset 45°", func(t *testing.T) {
		eyev := tuple.Vector(0, 0, -1)
		normalv := tuple.Vector(0, 0, -1)
		light := light.NewPointLight(tuple.Point(0, 10, -10), color.RGB(1, 1, 1))

		result := m.Lighting(light, position, eyev, normalv)

		assert.True(t, result.Equal(color.RGB(0.7364, 0.7364, 0.7364)))
	})

	t.Run("Lighting with the eye in the path of the refletion vector", func(t *testing.T) {
		eyev := tuple.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
		normalv := tuple.Vector(0, 0, -1)
		light := light.NewPointLight(tuple.Point(0, 10, -10), color.RGB(1, 1, 1))

		result := m.Lighting(light, position, eyev, normalv)

		assert.True(t, result.Equal(color.RGB(1.6364, 1.6364, 1.6364)))
	})

	t.Run("Lighting with the light behind the surface", func(t *testing.T) {
		eyev := tuple.Vector(0, 0, -1)
		normalv := tuple.Vector(0, 0, -1)
		light := light.NewPointLight(tuple.Point(0, 0, 10), color.RGB(1, 1, 1))

		result := m.Lighting(light, position, eyev, normalv)

		assert.Equal(t, color.RGB(0.1, 0.1, 0.1), result)
		assert.True(t, result.Equal(color.RGB(0.1, 0.1, 0.1)))
	})
}
