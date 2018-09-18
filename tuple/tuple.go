package tuple

const (
	pointType  = float64(1)
	vectorType = float64(0)
)

// Tuple represents a position
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (t *Tuple) isPoint() bool {
	return t.W == pointType
}

func (t *Tuple) isVector() bool {
	return t.W == vectorType
}

// Add a Tuple to another one
func (t *Tuple) Add(other *Tuple) *Tuple {
	return &Tuple{
		X: t.X + other.X,
		Y: t.Y + other.Y,
		Z: t.Z + other.Z,
		W: t.W + other.W,
	}
}

// Subtract a Tuple to another one
func (t *Tuple) Subtract(other *Tuple) *Tuple {
	return &Tuple{
		X: t.X - other.X,
		Y: t.Y - other.Y,
		Z: t.Z - other.Z,
		W: t.W - other.W,
	}
}

// NewPoint returns a Tuple that represents a Point
func Point(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: pointType,
	}
}

// NewVector returns a Tuple that represents a Vector
func Vector(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: vectorType,
	}
}
