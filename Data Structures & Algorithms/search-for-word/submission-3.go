func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var searchWord func(int, int, int) bool
	searchWord = func(r, c, i int) bool {
	
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[i] {
			return false
		}

		if i == len(word)-1 {
			return true
		}


		original := board[r][c]
		board[r][c] = '#'

		exists := searchWord(r+1, c, i+1) ||
			   searchWord(r-1, c, i+1) ||
			   searchWord(r, c+1, i+1) ||
			   searchWord(r, c-1, i+1)
		
		board[r][c] = original
		return exists
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if searchWord(i, j, 0) {
				return true
			}
		}
	}

	return false
}
