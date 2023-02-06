package gold

// https://leetcode.com/problems/path-with-maximum-gold

func GetMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	maxGold := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// try starting at every non-0 position
			if grid[i][j] > 0 {
				// find max gold starting at i, j
				maxGold = maxOf(maxGold, dfs(grid, i, j, 0))
			}
		}
	}
	return maxGold
}

var directions = [4][2]int{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // up
	{-1, 0}, // down
}

func dfs(grid [][]int, i, j, gold int) int {
	m, n := len(grid), len(grid[0])

	prevGold := grid[i][j]
	gold += grid[i][j]
	grid[i][j] = 0

	maxGold := gold
	for _, d := range directions {
		neigI, neigJ := i+d[0], j+d[1]
		if neigI >= 0 && neigI < m &&
			neigJ >= 0 && neigJ < n &&
			grid[neigI][neigJ] > 0 {
			maxGold = maxOf(maxGold, dfs(grid, neigI, neigJ, gold))
		}
	}

	grid[i][j] = prevGold

	return maxGold
}

func maxOf(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
