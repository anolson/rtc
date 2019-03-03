package ray

import (
	"testing"

	"github.com/anolson/rtc/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	origin := tuple.Point(1, 2, 3)
	direction := tuple.Vector(4, 5, 6)

	r := New(origin, direction)
	assert.Equal(t, origin, r.Origin)
	assert.Equal(t, direction, r.Direction)
}

func TestPosition(t *testing.T) {
	origin := tuple.Point(2, 3, 4)
	direction := tuple.Vector(1, 0, 0)

	r := New(origin, direction)

	assert.Equal(t, origin, r.Position(0))
	assert.Equal(t, tuple.Point(3, 3, 4), r.Position(1))
	assert.Equal(t, tuple.Point(1, 3, 4), r.Position(-1))
	assert.Equal(t, tuple.Point(4.5, 3, 4), r.Position(2.5))
	assert.Equal(t, direction, r.Direction)
}
