package main

import "fmt"

func exist(board [][]byte, word string) bool {
	n := len(board)
	if len(word) == 0 {
		return true
	}
	if n == 0 {
		return false
	}
	m := len(board[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != word[0] {
				continue
			}
			temp := board[i][j]
			board[i][j] = '.'
			if checkNeibor(i, j, board, word[1:]) {
				return true
			}
			board[i][j] = temp
		}
	}
	return false
}

func checkNeibor(x, y int, board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if x+1 < len(board) && board[x+1][y] != '.' {
		if board[x+1][y] == word[0] {
			temp := board[x+1][y]
			board[x+1][y] = '.'
			if checkNeibor(x+1, y, board, word[1:]) {
				return true
			}
			board[x+1][y] = temp
		}
	}
	if y+1 < len(board[0]) && board[x][y+1] != '.' {
		if board[x][y+1] == word[0] {
			temp := board[x][y+1]
			board[x][y+1] = '.'
			if checkNeibor(x, y+1, board, word[1:]) {
				return true
			}
			board[x][y+1] = temp
		}
	}
	if x-1 > -1 && board[x-1][y] != '.' {
		if board[x-1][y] == word[0] {
			temp := board[x-1][y]
			board[x-1][y] = '.'
			if checkNeibor(x-1, y, board, word[1:]) {
				return true
			}
			board[x-1][y] = temp
		}
	}
	if y-1 > -1 && board[x][y-1] != '.' {
		if board[x][y-1] == word[0] {
			temp := board[x][y-1]
			board[x][y-1] = '.'
			if checkNeibor(x, y-1, board, word[1:]) {
				return true
			}
			board[x][y-1] = temp
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	//word := "ABCCED"
	//word := "SEE"
	word := "ABCB"
	fmt.Println(exist(board, word))
}
