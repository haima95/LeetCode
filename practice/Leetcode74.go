package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	start := 0
	end := n - 1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if matrix[mid][0] < target && target < matrix[mid][m-1] {
			break
		} else if matrix[mid][0] > target {
			end = mid - 1
		} else if target > matrix[mid][m-1] {
			start = mid + 1
		} else if target == matrix[mid][m-1] || target == matrix[mid][0] {
			return true
		}
	}
	temp := mid
	start = 0
	end = m - 1
	for start <= end {
		mid = (start + end) / 2
		if matrix[temp][mid] > target {
			end = mid - 1
		} else if target > matrix[temp][mid] {
			start = mid + 1
		} else if target == matrix[temp][mid] {
			return true
		}
	}
	return false
}

func main() {
	test := [][][]int{
		{{}},
		{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		{{1}},
		{{1, 2, 3}},
		{{1}, {2}, {3}},
	}
	result := []int{0, 3, 13, 1, 4, 1}
	for i := 0; i < len(test); i++ {
		fmt.Println(searchMatrix(test[i], result[i]))
	}
}
