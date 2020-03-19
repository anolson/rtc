package primitives

import (
	"testing"

	"github.com/anolson/rtc/color"
	"github.com/anolson/rtc/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNewPointLight(t *testing.T) {
	t.Run("A point light has a position and an intensity", func(t *testing.T) {
		intensity := color.RGB(1, 1, 1)
		position := tuple.Point(0, 0, 0)

		light := NewPointLight(position, intensity)

		assert.Equal(t, intensity, light.Intensity)
		assert.Equal(t, position, light.Position)
	})
}
