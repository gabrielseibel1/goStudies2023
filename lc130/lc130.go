package lc130

func solve(board [][]byte) {
	var p position

	m, n := len(board), len(board[0])

	for r := 0; r < m; r++ {
		p = position{r, 0}
		if board[p.r][p.c] == 'O' && p.isEdge(board) {
			expandRegion(board, p)
		}
		p = position{r, n - 1}
		if board[p.r][p.c] == 'O' && p.isEdge(board) {
			expandRegion(board, p)
		}
	}
	for c := 0; c < n; c++ {
		p = position{0, c}
		if board[p.r][p.c] == 'O' && p.isEdge(board) {
			expandRegion(board, p)
		}
		p = position{m - 1, c}
		if board[p.r][p.c] == 'O' && p.isEdge(board) {
			expandRegion(board, p)
		}
	}
	for r := range board {
		for c := range board[r] {
			switch board[r][c] {
			case 'O':
				board[r][c] = 'X'
			case '-':
				board[r][c] = 'O'
			}
		}
	}
}

func expandRegion(board [][]byte, p position) {
	// expandRegionBFS(board, p)
	expandRegionDFS(board, p)
}

func expandRegionDFS(board [][]byte, p position) {
	board[p.r][p.c] = '-'
	for _, d := range p.directions(board) {
		expandRegionDFS(board, d)
	}
}

// func expandRegionBFS(board [][]byte, p position) {
// 	border := make([]position, 1)
// 	border[0] = p

// 	for batchSize := len(border); batchSize > 0; batchSize = len(border) {
// 		for i := 0; i < batchSize; i++ {
// 			pos := border[i]
// 			border = append(border, pos.directions(board)...)
// 			board[pos.r][pos.c] = '-'
// 		}
// 		border = border[batchSize:]
// 	}
// }

type position struct {
	r int
	c int
}

func (p position) isWithin(board [][]byte) bool {
	return p.r >= 0 && p.r < len(board) && p.c >= 0 && p.c < len(board[p.r])
}

func (p position) isEdge(board [][]byte) bool {
	return p.r == 0 || p.r == len(board)-1 || p.c == 0 || p.c == len(board[p.r])-1
}

func (p position) directions(board [][]byte) []position {
	dirs := make([]position, 0, 4)
	u, r, d, l := p, p, p, p
	u.r -= 1
	r.c += 1
	d.r += 1
	l.c -= 1
	for _, dir := range []position{u, r, d, l} {
		if dir.isWithin(board) && board[dir.r][dir.c] == 'O' {
			dirs = append(dirs, dir)
		}
	}
	return dirs
}
