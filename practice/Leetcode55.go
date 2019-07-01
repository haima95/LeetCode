package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	maxJ := 0
	for j := 0; j < len(nums); {
		max := 0
		for i := 1; i <= nums[j]; i++ {
			if i+j >= len(nums)-1 {
				return true
			}
			if i+j+nums[i+j] > max {
				max = i + nums[i+j] + j
				maxJ = i + j
			}
		}
		if max == 0 {
			return false
		}
		j = maxJ
	}
	return true
}

func main() {
	arr := []int{2, 0, 0}
	fmt.Println(canJump(arr))
}
