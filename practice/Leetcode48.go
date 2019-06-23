package main

import "fmt"

func rotate(matrix [][]int) {
	row := len(matrix)
	if row < 2 {
		return
	}
	for i := 0; i < row/2; i++ {
		for j := i; j < row-i-1; j++ {
			r := i
			c := j
			temp := matrix[r][c]

			// 向右
			if c > 0 {
				matrix[c][row-r-1], temp = temp, matrix[c][row-r-1]
				r, c = c, row-r-1
			} else if c == 0 {
				matrix[r][row-1], temp = temp, matrix[r][row-1]
				c = row - 1
			}
			// 向下
			if r > 0 {
				matrix[c][row-r-1], temp = temp, matrix[c][row-r-1]
				r, c = c, row-r-1
			} else if r == 0 {
				matrix[row-1][c], temp = temp, matrix[row-1][c]
				r = row - 1
			}
			// 向左
			if c < row-1 {
				matrix[c][row-r-1], temp = temp, matrix[c][row-r-1]
				r, c = c, row-r-1
			} else if c == row-1 {
				matrix[r][row-1-c], temp = temp, matrix[r][row-1-c]
				c = 0
			}
			// 向上
			if r < row-1 {
				matrix[c][row-r-1], temp = temp, matrix[c][row-r-1]
				r, c = c, row-r-1
			} else if r == row-1 {
				matrix[row-1-r][c], temp = temp, matrix[row-1-r][c]
				r = 0
			}
		}
	}
}

func printArr(step int, arr [][]int) {
	fmt.Println("step #", step)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%2d ", arr[i][j])
		}
		fmt.Println()
	}
}

func main() {

	//matrix := [][]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//	{13, 14, 15, 16},
	//}
	matrix2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix2)
	fmt.Println(matrix2)
}
