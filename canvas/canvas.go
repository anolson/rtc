package canvas

import (
	"errors"
	"fmt"

	"github.com/anolson/rtc/color"
)

// Canvas represents a rectagular grid of pixels
type Canvas struct {
	Width  int
	Height int
	Pixels [][]*color.Color
}

var (
	ErrorOutOfBounds = errors.New("out of bounds")
)

// New returns a new Canvas object
func New(w, h int) *Canvas {
	c := &Canvas{
		Width:  w,
		Height: h,
	}

	c.Pixels = make([][]*color.Color, h)

	for i := 0; i < h; i++ {
		c.Pixels[i] = make([]*color.Color, w)
		for j := 0; j < w; j++ {
			c.Pixels[i][j] = color.RGB(0, 0, 0)
		}
	}

	return c
}

// WritePixel sets a pixel at x,y to the provided Color
func (c *Canvas) WritePixel(x, y int, color *color.Color) error {
	if x < 0 || y < 0 || x >= c.Width || y >= c.Height {
		return ErrorOutOfBounds
	}

	c.Pixels[y][x] = color

	return nil
}

// PixelAt returns the pixel at x,y
func (c *Canvas) PixelAt(x, y int) (*color.Color, error) {
	if x < 0 || y < 0 || x >= c.Width || y >= c.Height {
		return nil, fmt.Errorf("out of bounds")
	}

	return c.Pixels[y][x], nil
}
