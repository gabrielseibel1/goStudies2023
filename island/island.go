package island

// solve https://leetcode.com/problems/shortest-bridge

import "math"

type coordinate struct {
	row int
	col int
}

func ShortestBridge(grid [][]int) int {
	start := firstLand(grid)

	q := make([]coordinate, 0)
	dfsFindLand(grid, start, &q)

	levels := bfsCountingLevels(grid, q)
	return levels
}

func firstLand(grid [][]int) coordinate {
	for i, row := range grid {
		for j, ele := range row {
			if ele == 1 {
				return coordinate{i, j}
			}
		}
	}
	return coordinate{}
}

func dfsFindLand(grid [][]int, pos coordinate, q *[]coordinate) {
	*q = append(*q, pos)
	grid[pos.row][pos.col] = math.MinInt

	for _, next := range neighbors(grid, pos) {
		if grid[next.row][next.col] == 1 {
			dfsFindLand(grid, next, q)
		}
	}
}

func bfsCountingLevels(grid [][]int, q []coordinate) int {
	level := 2
	for len(q) > 0 {
		batch := len(q)
		for i := 0; i < batch; i++ {
			pos := q[0]
			q = q[1:]

			if grid[pos.row][pos.col] > 0 {
				continue
			}
			grid[pos.row][pos.col] = level

			for _, next := range neighbors(grid, pos) {
				switch grid[next.row][next.col] {
				case 0:
					q = append(q, next)
				case 1:
					return level - 2
				}
			}
		}
		level++
	}
	return math.MaxInt
}

func neighbors(grid [][]int, pos coordinate) []coordinate {
	validCoords := make([]coordinate, 0)
	coords := []coordinate{
		{pos.row, pos.col - 1},
		{pos.row, pos.col + 1},
		{pos.row - 1, pos.col},
		{pos.row + 1, pos.col},
	}
	for _, c := range coords {
		if c.row >= 0 &&
			c.row < len(grid) &&
			c.col >= 0 &&
			c.col < len(grid) {
			validCoords = append(validCoords, c)
		}
	}
	return validCoords
}
