package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var result []int
	rows := len(matrix)
	if rows < 1 {
		return result
	}
	cols := len(matrix[0])
	for i := 0; i <= (rows-1)/2; i++ {
		for j := i; j < cols-i && i <= (rows-1)/2; j++ {
			result = append(result, matrix[i][j])
		}
		for j := i + 1; j < rows-1-i && cols-i-1 >= cols/2; j++ {
			result = append(result, matrix[j][cols-i-1])
		}
		for j := cols - 1 - i; j > i-1 && rows-i-1 > (rows-1)/2; j-- {
			result = append(result, matrix[rows-i-1][j])
		}
		for j := rows - 2 - i; j > i && i < cols/2; j-- {
			result = append(result, matrix[j][i])
		}
	}
	return result
}

func main() {
	arr := [][]int{{1}, {4}, {7}}
	fmt.Println(spiralOrder(arr))
}
