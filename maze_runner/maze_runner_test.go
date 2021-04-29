package maze_runner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var maze = [][]int{
	{1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 3},
	{1, 0, 1, 0, 1, 0, 1},
	{0, 0, 1, 0, 0, 0, 1},
	{1, 0, 1, 0, 1, 0, 1},
	{1, 0, 0, 0, 0, 0, 1},
	{1, 2, 1, 0, 1, 0, 1},
}

func TestMazeFixture1(t *testing.T) {
	var moves = []string{"N", "N", "N", "N", "N", "E", "E", "E", "E", "E"}

	var result = Move(maze, moves)

	assert.Equal(t, "Finish", result)
}

func TestMazeFixture2(t *testing.T) {
	var moves = []string{"N", "N", "N", "W", "W"}

	var result = Move(maze, moves)

	assert.Equal(t, "Dead", result)
}

func TestMazeFixture3(t *testing.T) {
	var moves = []string{"N", "E", "E", "E", "E"}

	var result = Move(maze, moves)

	assert.Equal(t, "Dead", result)
}
