package lc695

const (
	seen  = -1
	water = 0
	land  = 1
)

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == seen {
				continue
			}
			if a := areaOfLand(grid, i, j); a > max {
				max = a
			}
		}
	}
	return max
}

func areaOfLand(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return 0
	}

	switch grid[i][j] {
	case seen:
		return 0
	case water:
		grid[i][j] = seen
		return 0
	case land:
		grid[i][j] = seen
		up := areaOfLand(grid, i+1, j)
		right := areaOfLand(grid, i, j+1)
		down := areaOfLand(grid, i-1, j)
		left := areaOfLand(grid, i, j-1)
		return 1 + up + right + down + left
	default:
		panic("weird cell")
	}
}
