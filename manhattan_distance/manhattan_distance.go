package manhattan_distance

import (
    "github.com/Harcyah/code-kata-go/maths"
)

func ManhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	return maths.Abs(x2-x1) + maths.Abs(y2-y1)
}
