package main

import "fmt"

func searchInsert(nums []int, target int) int {
	var result int
	start := 0
	n := len(nums)
	end := n - 1
	mid := (start + end) / 2
	for start <= end {
		if nums[mid] > target {
			if mid > 0 && nums[mid-1] >= target {
				end = mid - 1
			} else {
				result = mid
				break
			}
		} else if nums[mid] < target {
			if mid < n-1 && nums[mid+1] <= target {
				start = mid + 1
			} else {
				result = mid + 1
				break
			}
		} else {
			result = mid
			break
		}
		mid = (start + end) / 2
	}
	return result
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 1
	fmt.Println(searchInsert(nums, target))
}
