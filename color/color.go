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

// ScaledRGBCompents returns each color component scaled to lie between 0 and 255
func ScaledRGB(c *Color) *Color {
	return &Color{
		Red:   scaledRGBComponent(c.Red),
		Green: scaledRGBComponent(c.Green),
		Blue:  scaledRGBComponent(c.Blue),
	}
}

func scaledRGBComponent(value float64) float64 {
	return clamp(math.Round(value * 255))
}

func clamp(value float64) float64 {
	return math.Max(0, math.Min(value, 255))
}
