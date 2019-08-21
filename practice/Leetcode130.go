package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	temp := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		temp[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(temp); i++ {
		solveO(i, 0, temp, board)
		solveO(i, len(temp[0])-1, temp, board)
	}
	for i := 0; i < len(temp[0]); i++ {
		solveO(0, i, temp, board)
		solveO(len(temp)-1, i, temp, board)
	}
	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(temp[i]); j++ {
			if !temp[i][j] && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solveO(x, y int, temp [][]bool, board [][]byte) {
	if x < 0 || x >= len(temp) || y < 0 || y >= len(temp[0]) {
		return
	}
	if !temp[x][y] && board[x][y] == 'O' {
		temp[x][y] = true
		solveO(x+1, y, temp, board)
		solveO(x-1, y, temp, board)
		solveO(x, y-1, temp, board)
		solveO(x, y+1, temp, board)
	}
}

func main() {
	//x := [][]byte{
	//	{'X', 'X', 'X', 'X'},
	//	{'X', 'O', 'O', 'X'},
	//	{'X', 'X', 'O', 'X'},
	//	{'X', 'O', 'X', 'X'},
	//}
	x := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'O', 'O', 'X', 'O', 'X'}}
	//x := [][]byte{{'O'}}
	solve(x)
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Printf("%c ", x[i][j])
		}
		fmt.Println()
	}
}
