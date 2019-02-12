package matrix

import (
	"testing"

	"github.com/anolson/rtc/tuple"
	"github.com/stretchr/testify/assert"
)

func TestMultiplyByMatrix(t *testing.T) {
	a := New(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	b := New(4, 4, []float64{
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	})

	result := New(4, 4, []float64{
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	})

	assert.Equal(t, result, Multiply(a, b))
}

func TestMultiplyByTuple(t *testing.T) {
	a := New(4, 4, []float64{
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	})

	b := tuple.New(1, 2, 3, 1)
	result := tuple.New(18.0, 24.0, 33.0, 1.0)

	assert.Equal(t, result, MultiplyByTuple(a, b))
}

func TestMultiplyByIdentityMatrix(t *testing.T) {
	a := New(4, 4, []float64{
		0, 1, 2, 4,
		1, 2, 4, 8,
		2, 4, 8, 16,
		4, 8, 16, 32,
	})

	identityMatrix := New(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})

	assert.Equal(t, a, Multiply(a, identityMatrix))
}

func TestMultiplyIdentityMatrixByTuple(t *testing.T) {
	a := tuple.New(1, 2, 3, 1)
	identityMatrix := New(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})

	assert.Equal(t, a, MultiplyByTuple(identityMatrix, a))
}
