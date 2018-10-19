package canvas

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/anolson/rtc/color"
	"github.com/stvp/assert"
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
	c := New(5, 3)

	buffer := bytes.NewBuffer([]byte{})
	c.Save(buffer)
	assert.Equal(t, "P3\n5 3\n255", buffer.String())
}
