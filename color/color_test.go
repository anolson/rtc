package color

import (
	"testing"

	"github.com/anolson/rtc/util"
	"github.com/stvp/assert"
)

func TestRGB(t *testing.T) {
	c := RGB(-0.5, 0.4, 1.7)

	assert.Equal(t, -0.5, c.Red)
	assert.Equal(t, 0.4, c.Green)
	assert.Equal(t, 1.7, c.Blue)
}

func TestAdd(t *testing.T) {
	c1 := RGB(0.9, 0.6, 0.75)
	c2 := RGB(0.7, 0.1, 0.25)

	result := Add(c1, c2)
	assert.True(t, util.Approx(1.6, result.Red))
	assert.True(t, util.Approx(0.7, result.Green))
	assert.True(t, util.Approx(1.0, result.Blue))
}

func TestSubtract(t *testing.T) {
	c1 := RGB(0.9, 0.6, 0.75)
	c2 := RGB(0.7, 0.1, 0.25)

	result := Subtract(c1, c2)
	assert.True(t, util.Approx(0.2, result.Red))
	assert.True(t, util.Approx(0.5, result.Green))
	assert.True(t, util.Approx(0.5, result.Blue))
}

func TestMutiply(t *testing.T) {
	c := RGB(0.2, 0.3, 0.4)

	result := Multiply(c, 2)
	assert.True(t, util.Approx(0.4, result.Red))
	assert.True(t, util.Approx(0.6, result.Green))
	assert.True(t, util.Approx(0.8, result.Blue))
}

func TestHadamardProduct(t *testing.T) {
	c1 := RGB(1, 0.2, 0.4)
	c2 := RGB(0.9, 1, 0.1)

	result := HadamardProduct(c1, c2)
	assert.True(t, util.Approx(0.9, result.Red))
	assert.True(t, util.Approx(0.2, result.Green))
	assert.True(t, util.Approx(0.04, result.Blue))
}
