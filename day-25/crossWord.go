package main

func Exist(board [][]byte, word string) bool {
	var b func(r, c, i int) bool
	b = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		if r < 0 || r > len(board[0]) || c < 0 || c > len(board) || board[r][c] != word[i] {
			return false
		}
		t := board[r][c]
		board[r][c] = '.'
		f := b(r+1, c, i+1) ||
			b(r, c+1, i+1) ||
			b(r-1, c, i+1) ||
			b(r, c-1, i+1)
		board[r][c] = t
		return f
	}
	for r := 0; r < len(board[0]); r++ {
		for c := 0; c < len(board); c++ {
			if b(r, c, 0) {
				return true
			}
		}
	}
	return false
}

// 99ms
// 4.05mb