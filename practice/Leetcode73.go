package main

import "fmt"

func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	var flag [][2]int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				temp := [2]int{i, j}
				flag = append(flag, temp)
			}
		}
	}
	for k := 0; k < len(flag); k++ {
		for i := 0; i < n; i++ {
			matrix[i][flag[k][1]] = 0
		}
		for j := 0; j < m; j++ {
			matrix[flag[k][0]][j] = 0
		}
	}
}

func main() {
	test := [][][]int{
		{{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1}},

		{{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5}},
	}
	for i := 0; i < len(test); i++ {
		setZeroes(test[i])
		for j := 0; j < len(test[i]); j++ {
			for k := 0; k < len(test[i][j]); k++ {
				fmt.Printf("%d ", test[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
}
