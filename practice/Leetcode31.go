package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	s := n - 1
	i := n - 2
	for ; i > -1; i-- {
		if nums[i] >= nums[s] {
			s = i
			continue
		} else {
			for j := n - 1; j >= s; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
			sort.Ints(nums[i+1:])
			break
		}
	}
	if i == -1 {
		sort.Ints(nums)
	}
}

func isEquals(nums []int, nums2 []int) bool {
	if len(nums) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			return false
		}
	}
	return true
}

func main() {
	nums := [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2},
		{2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1},
		{3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1},
		{4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1}}
	for i := 0; i < len(nums)-1; i++ {
		fmt.Printf("%v, %v  ", nums[i], nums[i+1])
		nextPermutation(nums[i])
		fmt.Println(isEquals(nums[i], nums[i+1]))
	}
}
