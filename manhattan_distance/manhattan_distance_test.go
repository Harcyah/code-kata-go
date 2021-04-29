package manhattan_distance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	assert.Equal(t, 0, ManhattanDistance(1, 1, 1, 1))
	assert.Equal(t, 4, ManhattanDistance(5, 4, 3, 2))
	assert.Equal(t, 3, ManhattanDistance(1, 1, 0, 3))
}
