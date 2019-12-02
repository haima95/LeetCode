package main

import (
	"fmt"
	"testing"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	for i := 0; i < k; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums[k-1]
}

func Test215(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}
