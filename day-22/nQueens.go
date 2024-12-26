package main

func TotalNQueens(n int) int {
    board := make([][]byte, n)
    var res [][]string
    for i := range board {
        board[i] = make([]byte, n)
        for j := range board[i] {
            board[i][j] = '.'
        }
    }
    backtrack(board, 0, &res)
    return len(res)
}

func backtrack(board [][]byte, row int, res *[][]string) {
    if row == len(board) {
        sol := make([]string, len(board))
        for i := range board {
            sol[i] = string(board[i])
        }
        *res = append(*res, sol)
        return
    }
    for col := 0; col < len(board); col++ {
        if isSafe(board, row, col) {
            board[row][col] = 'Q'
            backtrack(board, row+1, res)
            board[row][col] = '.'
        }
    }
}

func isSafe(board [][]byte, row, col int) bool {
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    return true
}

// 0ms
// 4.12mb