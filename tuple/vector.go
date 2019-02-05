package tuple

// Vector returns a Tuple that represents a Vector
func Vector(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: vectorType,
	}
}
