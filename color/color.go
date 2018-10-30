package color

import "math"

// Color represents an rgb color
type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

// RGB returns a new Color object
func RGB(r, g, b float64) *Color {
	return &Color{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}

// Add a Color to another one
func Add(c1, c2 *Color) *Color {
	return &Color{
		Red:   c1.Red + c2.Red,
		Green: c1.Green + c2.Green,
		Blue:  c1.Blue + c2.Blue,
	}
}

// Subtract a Color from another one
func Subtract(c1, c2 *Color) *Color {
	return &Color{
		Red:   c1.Red - c2.Red,
		Green: c1.Green - c2.Green,
		Blue:  c1.Blue - c2.Blue,
	}
}

// Multiply a Color by a scalar value
func Multiply(c *Color, value float64) *Color {
	return &Color{
		Red:   c.Red * value,
		Green: c.Green * value,
		Blue:  c.Blue * value,
	}
}

// HadamardProduct multiplies each color component
func HadamardProduct(c1, c2 *Color) *Color {
	return &Color{
		Red:   c1.Red * c2.Red,
		Green: c1.Green * c2.Green,
		Blue:  c1.Blue * c2.Blue,
	}
}

// Scaled returns each color component scaled to lie between (min-max)
func Scaled(c *Color, min, max float64) *Color {
	return &Color{
		Red:   clamp(c.Red, min, max),
		Green: clamp(c.Green, min, max),
		Blue:  clamp(c.Blue, min, max),
	}
}

func clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(math.Round(value*max), max))
}
