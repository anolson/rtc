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
func (c *Color) Add(other *Color) *Color {
	return &Color{
		Red:   c.Red + other.Red,
		Green: c.Green + other.Green,
		Blue:  c.Blue + other.Blue,
	}
}

// Subtract a Color to another one
func (c *Color) Subtract(other *Color) *Color {
	return &Color{
		Red:   c.Red - other.Red,
		Green: c.Green - other.Green,
		Blue:  c.Blue - other.Blue,
	}
}

// Multiply a Color by a scalar value
func (c *Color) Multiply(value float64) *Color {
	return &Color{
		Red:   c.Red * value,
		Green: c.Green * value,
		Blue:  c.Blue * value,
	}
}

// HadamardProduct multiplies each color component
func (c *Color) HadamardProduct(other *Color) *Color {
	return &Color{
		Red:   c.Red * other.Red,
		Green: c.Green * other.Green,
		Blue:  c.Blue * other.Blue,
	}
}
