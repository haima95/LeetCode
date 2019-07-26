package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	max := 0
	hist := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				hist[j]++
			} else {
				hist[j] = 0
			}
		}
		temp := getMaxArea(hist)
		if temp > max {
			max = temp
		}
	}
	return max
}

func getMaxArea(arr []int) int {
	var stack []int
	var max int
	for i := 0; i <= len(arr); i++ {
		var height int
		var cur int
		if i == len(arr) {
			cur = -1
		} else {
			cur = arr[i]
		}
		for len(stack) != 0 && cur < arr[stack[len(stack)-1]] {
			width := arr[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				height = i
			} else {
				height = i - stack[len(stack)-1] - 1
			}
			temp := width * height
			if temp > max {
				max = temp
			}
		}
		stack = append(stack, i)
	}
	return max
}

func printIntArray(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%2d ", arr[i][j])
		}
		fmt.Println()
	}
}

func main() {
	input := [][]byte{
		//{'1', '0', '1', '0', '0'},
		//{'1', '0', '1', '1', '1'},
		//{'1', '1', '1', '1', '1'},
		//{'1', '0', '0', '1', '0'},
		{'1', '1', '1', '1'},
		//{'1'},
	}
	//input := [][]byte{
	//	{'0', '1', '1', '0', '1'},
	//	{'1', '1', '0', '1', '0'},
	//	{'0', '1', '1', '1', '0'},
	//	{'1', '1', '1', '1', '0'},
	//	{'1', '1', '1', '1', '1'},
	//	{'0', '0', '0', '0', '0'}}
	fmt.Println(maximalRectangle(input))
}
