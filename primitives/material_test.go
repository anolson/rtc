package primitives

import (
	"testing"

	"github.com/anolson/rtc/color"
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
