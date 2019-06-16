package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	if n < 1 {
		return 0
	}
	left_max := make([]int, n)
	right_max := make([]int, n)
	left_max[0] = height[0]
	for i := 1; i < n; i++ {
		if height[i] > left_max[i-1] {
			left_max[i] = height[i]
		} else {
			left_max[i] = left_max[i-1]
		}
	}
	right_max[n-1] = height[n-1]
	for i := n - 2; i > -1; i-- {
		if height[i] > right_max[i+1] {
			right_max[i] = height[i]
		} else {
			right_max[i] = right_max[i+1]
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		if left_max[i] > right_max[i] {
			result += right_max[i] - height[i]
		} else {
			result += left_max[i] - height[i]
		}
	}
	return result
}

func trap2(height []int) int {
	var left_max, right_max int
	n := len(height)
	anx := 0
	var left, right = 0, n - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= left_max {
				left_max = height[left]
			} else {
				anx += left_max - height[left]
			}
			left++
		} else {
			if height[right] >= right_max {
				right_max = height[right]
			} else {
				anx += right_max - height[right]
			}
			right--
		}
	}
	return anx
}

func main() {
	nums := [][]int{
		{2, 1, 0},
		{1, 0, 0, 1, 0, 0, 1},
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{5, 4, 1, 2}}
	for _, v := range nums {
		fmt.Println(trap2(v))

	}
}
