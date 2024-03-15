package lc994

// import "fmt"

type position struct {
	row int
	col int
}

const (
	empty = iota
	fresh
	rotten
)

func orangesRotting(grid [][]int) int {
	q := make([]position, 0)

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == rotten {
				q = append(q, position{row, col})
			}
		}
	}

	minutes := -1
	size := len(q)
	for size > 0 {
		for i := 0; i < size; i++ {
			here := q[0]
			q = q[1:]
			for _, there := range neighbors(here, grid) {
				grid[there.row][there.col] = rotten
				q = append(q, there)
			}
		}
		minutes++
		size = len(q)
	}

	none := true
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == fresh {
				return -1
			}
			none = none && (grid[row][col] == empty)
		}
	}

	if none {
		return 0
	}

	return minutes
}

func neighbors(here position, grid [][]int) []position {
	n := make([]position, 0, 4)
	up := position{row: here.row - 1, col: here.col}
	dn := position{row: here.row + 1, col: here.col}
	rt := position{row: here.row, col: here.col + 1}
	lt := position{row: here.row, col: here.col - 1}
	for _, pos := range []position{up, dn, rt, lt} {
		if inbound(pos, grid) && grid[pos.row][pos.col] == fresh {
			n = append(n, pos)
		}
	}
	return n
}

func inbound(here position, grid [][]int) bool {
	return here.row >= 0 &&
		here.col >= 0 &&
		here.row < len(grid) &&
		here.col < len(grid[0])
}
