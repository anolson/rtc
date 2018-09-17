package tuple

const (
	pointType  = float64(1)
	vectorType = float64(0)
)

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

// NewPoint returns a Tuple that represents a Point
func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: pointType,
	}
}

// NewVector returns a Tuple that represents a Vector
func NewVector(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: vectorType,
	}
}
