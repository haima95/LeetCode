package main

import "fmt"

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	temp := make(map[int]int)
	result := 0
	for i := 0; i < n; i++ {
		if _, ok := temp[nums[i]]; ok {
			continue
		}
		left := 0
		right := 0
		if v, ok := temp[nums[i]-1]; ok {
			left = v
		}
		if v, ok := temp[nums[i]+1]; ok {
			right = v
		}
		val := left + right + 1
		temp[nums[i]-left] = val
		temp[nums[i]+right] = val
		temp[nums[i]] = val
		if result < val {
			result = val
		}
	}
	return result
}

func main() {
	input := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(input))
}
