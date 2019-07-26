package main

import "fmt"

func largestRectangleArea(heights []int) int {
	n := len(heights)
	var stack []int
	max := 0
	index := 0
	for index < n {
		if len(stack) == 0 || heights[index] >= heights[stack[len(stack)-1]] {
			stack = append(stack, index)
			index++
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var area int
			if len(stack) == 0 {
				area = heights[temp] * index
			} else {
				area = heights[temp] * (index - 1 - stack[len(stack)-1])
			}
			if area > max {
				max = area
			}
		}
	}
	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var area int
		if len(stack) == 0 {
			area = heights[temp] * index
		} else {
			area = heights[temp] * (index - 1 - stack[len(stack)-1])
		}
		if area > max {
			max = area
		}
	}
	return max
}

func main() {
	height := []int{2, 1, 5, 6, 2, 3}
	//height := []int{1, 2, 3, 4, 5}
	//height := []int{5, 4, 3, 2, 1}
	//height := []int{1, 5, 1, 4, 1, 1, 1}
	//height := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//height := []int{8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(largestRectangleArea(height))
}
