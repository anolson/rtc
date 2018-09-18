package tuple

import "math"

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

// Negate a Tuple
func (t *Tuple) Negate() *Tuple {
	return &Tuple{
		X: float64(0) - t.X,
		Y: float64(0) - t.Y,
		Z: float64(0) - t.Z,
		W: float64(0) - t.W,
	}
}

// Multiply a Tuple by a value
func (t *Tuple) Multiply(value float64) *Tuple {
	return &Tuple{
		X: t.X * value,
		Y: t.Y * value,
		Z: t.Z * value,
		W: t.W * value,
	}
}

// Divide a Tuple by a value
func (t *Tuple) Divide(value float64) *Tuple {
	return &Tuple{
		X: t.X / value,
		Y: t.Y / value,
		Z: t.Z / value,
		W: t.W / value,
	}
}

// Magnitude calculate the length of a vector
func (t *Tuple) Magnitude() float64 {
	squares := math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2)

	return math.Sqrt(squares)
}

// Normalize convert a vector to a unit vector
func (t *Tuple) Normalize() *Tuple {
	magnitude := t.Magnitude()
	return &Tuple{
		X: t.X / magnitude,
		Y: t.Y / magnitude,
		Z: t.Z / magnitude,
		W: t.W / magnitude,
	}
}

// Dot calulates the dot product of two vectors
func Dot(a, b *Tuple) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z) + (a.W * b.W)
}

// Point returns a Tuple that represents a Point
func Point(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: pointType,
	}
}

// Vector returns a Tuple that represents a Vector
func Vector(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: vectorType,
	}
}
