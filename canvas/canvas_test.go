package canvas

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/anolson/rtc/color"
	"github.com/stretchr/testify/assert"
	"gotest.tools/golden"
)

func TestNew(t *testing.T) {
	black := color.RGB(0, 0, 0)
	c := New(10, 20)

	assert.Equal(t, 10, c.Width)
	assert.Equal(t, 20, c.Height)

	fmt.Println(len(c.Pixels))
	fmt.Println(c.Pixels[0])

	for _, row := range c.Pixels {
		for _, pixel := range row {
			assert.Equal(t, black, pixel)
		}
	}
}

func TestWritePixel(t *testing.T) {
	red := color.RGB(1, 0, 0)
	c := New(10, 20)

	c.WritePixel(2, 3, red)

	pixel, err := c.PixelAt(2, 3)
	assert.Nil(t, err)
	assert.Equal(t, red, pixel)
}

func TestSave(t *testing.T) {
	t.Run("Constructing the PPM header", func(t *testing.T) {
		c := New(5, 3)
		buffer := bytes.NewBuffer([]byte{})

		c.Save(buffer)

		golden.Assert(t, buffer.String(), "ppm_header.golden")
	})

	t.Run("Constructing the PPM pixel data", func(t *testing.T) {
		c := New(5, 3)
		c1 := color.RGB(1.5, 0, 0)
		c2 := color.RGB(0, 0.5, 0)
		c3 := color.RGB(-0.5, 0, 1)
		c.WritePixel(0, 0, c1)
		c.WritePixel(2, 1, c2)
		c.WritePixel(4, 2, c3)
		buffer := bytes.NewBuffer([]byte{})

		c.Save(buffer)

		golden.Assert(t, buffer.String(), "ppm_canvas.golden")
	})

	t.Run("Splitting long lines in PPM files", func(t *testing.T) {
		c := New(10, 2)
		color := color.RGB(1, 0.8, 0.6)

		for y := 0; y < 2; y++ {
			for x := 0; x < 10; x++ {
				c.WritePixel(x, y, color)
			}
		}
		buffer := bytes.NewBuffer([]byte{})

		c.Save(buffer)

		golden.Assert(t, buffer.String(), "ppm_canvas_line_wrap.golden")
	})
}
