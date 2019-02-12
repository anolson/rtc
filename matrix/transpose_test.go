package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	a := New(4, 4, []float64{
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8,
	})

	result := New(4, 4, []float64{
		0, 9, 1, 0,
		9, 8, 8, 0,
		3, 0, 5, 5,
		0, 8, 3, 8,
	})

	assert.Equal(t, result, Transpose(a))
}

func TestTransposeIdentityMatrix(t *testing.T) {
	identityMatrix := New(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})

	assert.Equal(t, identityMatrix, Transpose(identityMatrix))
}
