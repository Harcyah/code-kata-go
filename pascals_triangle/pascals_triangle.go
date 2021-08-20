package pascal_triangle

func Triangle(height int) [][]int {
	if height <= 0 {
		panic("Height must be > 0")
	}

	var triangle = make([][]int, height)
	triangle[0] = make([]int, 1)
	triangle[0][0] = 1

	for i := 1; i < height; i++ {
		fillLevel(triangle, i)
	}

	return triangle
}

func fillLevel(triangle [][]int, level int) {
	triangle[level] = make([]int, level + 1)
	triangle[level][0] = 1
	triangle[level][level] = 1

	for i := 1; i < level; i++ {
		above := triangle[level - 1]
		left := above[i - 1]
		right := above[i]
		triangle[level][i] = left + right
	}
}

