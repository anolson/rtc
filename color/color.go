package color

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
