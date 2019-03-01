package tuple

import (
	"math"

	"github.com/anolson/rtc/util"
)

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

// New returns a new a Tuple object
func New(x, y, z, w float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func (t *Tuple) isPoint() bool {
	return t.W == pointType
}

func (t *Tuple) isVector() bool {
	return t.W == vectorType
}

// Equal returns true if a Tuple is equal to another, otherwise false
func (t *Tuple) Equal(other *Tuple) bool {
	return util.Approx(t.X, other.X) &&
		util.Approx(t.Y, other.Y) &&
		util.Approx(t.Z, other.Z) &&
		util.Approx(t.W, other.W)
}

// Add a Tuple to another one
func Add(a, b *Tuple) *Tuple {
	return &Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

// Subtract a Tuple from another one
func Subtract(a, b *Tuple) *Tuple {
	return &Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

// Negate a Tuple
func Negate(t *Tuple) *Tuple {
	return &Tuple{
		X: float64(0) - t.X,
		Y: float64(0) - t.Y,
		Z: float64(0) - t.Z,
		W: float64(0) - t.W,
	}
}

// Multiply a Tuple by a value
func Multiply(t *Tuple, value float64) *Tuple {
	return &Tuple{
		X: t.X * value,
		Y: t.Y * value,
		Z: t.Z * value,
		W: t.W * value,
	}
}

// Divide a Tuple by a value
func Divide(t *Tuple, value float64) *Tuple {
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

// Cross calulates the cross product of two vectors
func Cross(a, b *Tuple) *Tuple {
	return Vector(
		(a.Y*b.Z)-(a.Z*b.Y),
		(a.Z*b.X)-(a.X*b.Z),
		(a.X*b.Y)-(a.Y*b.X),
	)
}
