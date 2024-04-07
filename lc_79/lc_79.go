package lc79

func exist(board [][]byte, word string) bool {
	letters := []byte(word)

	for i := range board {
		for j := range board[i] {
			if hasWord(board, i, j, letters) {
				return true
			}
		}
	}
	return false
}

func hasWord(board [][]byte, i, j int, letters []byte) bool {
	here := board[i][j]
	if here != letters[0] {
		return false
	}
	if len(letters)-1 == 0 {
		return true
	}
	board[i][j] = '-'
	defer func() { board[i][j] = here }()
	for _, dir := range [4][2]int{
		{i - 1, j},
		{i + 1, j},
		{i, j - 1},
		{i, j + 1},
	} {
		if (dir[0] >= 0 && dir[0] < len(board)) &&
			(dir[1] >= 0 && dir[1] < len(board[dir[0]])) &&
			board[dir[0]][dir[1]] != '-' {
			if hasWord(board, dir[0], dir[1], letters[1:]) {
				return true
			}
		}
	}
	return false
}
