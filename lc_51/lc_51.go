package lc51

func solveNQueens(n int) [][]string {
	solutions := make(map[string][]string)

	// try every start position
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			q := [][2]int{{i, j}}
			fillNQueens(n, &q, solutions)
		}
	}

	boards := make([][]string, len(solutions))
	i := 0
	for _, v := range solutions {
		boards[i] = v
	}

	return boards
}

func fillNQueens(n int, q *[][2]int, solutions map[string][]string) {
	for {
		if len(*q) == n {
			v := boardWithNQueens(n, *q)
			solutions[key(v)] = v
			break
		}
		poss := availablePositions(n, *q)
		if len(poss) == 0 {
			break
		}
		for _, pos := range poss {
			*q = append(*q, pos)
			fillNQueens(n, q, solutions)
			*q = (*q)[:len(*q)-1]
		}
	}
}

func key(board []string) string {
	k := make([]byte, 0, len(board)*len(board))
	for _, s := range board {
		k = append(k, []byte(s)...)
	}
	return string(k)
}

func boardWithNQueens(n int, q [][2]int) []string {
	board := make([]string, n)
	for _, p := range q {
		row := make([]byte, n)
		for i := 0; i < n; i++ {
			row[i] = '.'
		}
		row[p[1]] = 'Q'
		board[p[0]] = string(row)
	}
	return board
}

func availablePositions(n int, q [][2]int) [][2]int {
	panic("ni")
}
