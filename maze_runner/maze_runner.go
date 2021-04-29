package maze_runner

// From https://dev.to/thepracticaldev/daily-challenge-288-maze-runner-1bnd

type Result = string

const Dead = "Dead"
const Finish = "Finish"

type Point struct {
	x int
	y int
}

func Move(maze [][]int, moves []string) Result {
	var start = findStart(maze)
	var finish = findFinish(maze)
	var current = Point{start.x, start.y}
	for _, move := range moves {
		switch move {
		case "N":
			current.y--
			break
		case "S":
			current.y++
			break
		case "W":
			current.x--
			break
		case "E":
			current.x++
			break
		default:
			panic("Unsupported move " + move)
		}

		if current.x == finish.x && current.y == finish.y {
			return Finish
		}

		if current.x < 0 || current.x > len(maze) || current.y < 0 || current.y > len(maze) {
			return Dead
		}

	}
	return Dead
}

func findStart(maze [][]int) Point {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze); j++ {
			if maze[i][j] == 2 {
				return Point{j, i}
			}
		}
	}
	panic("Maze has no start point !")
}

func findFinish(maze [][]int) Point {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze); j++ {
			if maze[i][j] == 3 {
				return Point{j, i}
			}
		}
	}
	panic("Maze has no finish point !")
}
