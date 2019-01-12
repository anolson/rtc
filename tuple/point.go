package tuple

// Point returns a Tuple that represents a Point
func Point(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: pointType,
	}
}
