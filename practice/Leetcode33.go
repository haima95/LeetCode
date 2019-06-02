package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	start := 0
	end := n - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[start] > target {
				start = start + 1
			} else {
				end = mid - 1
			}
		} else {
			if nums[end] < target {
				end = end - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}

func main() {
	//nums := []int{4, 5, 6, 7, 8, 1, 2}
	nums := []int{8, 1, 2, 4, 5, 6, 7}
	target := 8
	fmt.Println(search(nums, target))
}
