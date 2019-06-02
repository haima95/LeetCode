package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	result := []int{-1, -1}
	if n < 1 {
		return result
	}
	start := 0
	end := n - 1
	mid := -1
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			result[0] = mid
			result[1] = mid
			break
		}
	}

	start = 0
	end = result[0] - 1
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			result[0] = mid
			end = mid - 1
		}
	}
	start = result[1] + 1
	end = n - 1
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			result[1] = mid
			start = mid + 1
		}
	}
	return result
}

func main() {
	nums := []int{5, 7, 7, 7, 8, 8, 10}
	target := 7
	fmt.Println(searchRange(nums, target))
}
