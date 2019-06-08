package main

import "fmt"

func solveSudoku(board [][]byte) {

	var mRow [9][9]bool
	var mCol [9][9]bool
	var mSquare [9][9]bool

	for aRow := 0; aRow < 9; aRow++ {
		for aCol := 0; aCol < 9; aCol++ {

			if board[aRow][aCol] != '.' {
				aVal := board[aRow][aCol] - '1'
				mRow[aRow][aVal] = true
				mCol[aCol][aVal] = true
				aBoxRaw := aRow / 3
				aBoxCol := aCol / 3
				mSquare[aBoxRaw*3+aBoxCol][aVal] = true
			}
		}
	}
	solveSudokuInternal(board, 0, 0, mRow, mCol, mSquare)
}

func solveSudokuInternal(board [][]byte, pRow, pCol int, mRow, mCol, mSquare [9][9]bool) bool {
	var dic = map[int]byte{
		1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9',
	}
	if pRow == -1 {
		return true
	}
	aNextRow := pRow + 1
	aNextCol := pCol
	if aNextRow == 9 {
		aNextRow = 0
		aNextCol++
	}
	if aNextCol == 9 {
		aNextRow = -1
		aNextCol = -1
	}
	if board[pRow][pCol] != '.' {
		return solveSudokuInternal(board, aNextRow, aNextCol, mRow, mCol, mSquare)
	}
	aBoxRow := pRow / 3
	aBoxCol := pCol / 3
	aBox := aBoxRow*3 + aBoxCol
	for aTry := 0; aTry < 9; aTry++ {
		if mRow[pRow][aTry] {
			continue
		}
		if mCol[pCol][aTry] {
			continue
		}
		if mSquare[aBox][aTry] {
			continue
		}
		board[pRow][pCol] = dic[aTry+1]
		mRow[pRow][aTry] = true
		mCol[pCol][aTry] = true
		mSquare[aBox][aTry] = true

		if solveSudokuInternal(board, aNextRow, aNextCol, mRow, mCol, mSquare) {
			return true
		}

		mRow[pRow][aTry] = false
		mCol[pCol][aTry] = false
		mSquare[aBox][aTry] = false
	}
	board[pRow][pCol] = '.'
	return false
}

func main() {
	board := [][]byte{{'.', '.', '9', '7', '4', '8', '.', '.', '.'}, {'7', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '2', '.', '1', '.', '9', '.', '.', '.'}, {'.', '.', '7', '.', '.', '.', '2', '4', '.'}, {'.', '6', '4', '.', '1', '.', '5', '9', '.'}, {'.', '9', '8', '.', '.', '.', '3', '.', '.'}, {'.', '.', '.', '8', '.', '3', '.', '2', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '6'}, {'.', '.', '.', '2', '7', '5', '9', '.', '.'}}

	//board := [][]byte{
	//	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}

	solveSudoku2(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				fmt.Printf("0 ")
			} else {
				fmt.Printf("%d ", board[i][j]-'0')
			}
		}
		fmt.Println()
	}
	fmt.Println("============ end ============")
}
