package main

import "fmt"

func search(nums []int, target int) bool {
	s := 0
	e := len(nums) - 1
	return binarySearch(s, e, target, nums)
}

func binarySearch(s, e, target int, nums []int) bool {
	if s > e {
		return false
	}
	m := (s + e) / 2
	if target == nums[m] {
		return true
	} else if target > nums[m] {
		if nums[e] >= target {
			return binarySearch(m+1, e, target, nums)
		} else {
			if nums[e] > nums[m] {
				return binarySearch(s, m-1, target, nums)
			} else if nums[e] == nums[m] {
				if binarySearch(s, m-1, target, nums) {
					return true
				} else {
					return binarySearch(m+1, e, target, nums)
				}
			} else {
				return binarySearch(m+1, e, target, nums)
			}
		}
	} else {
		if nums[s] <= target {
			return binarySearch(s, m-1, target, nums)
		} else {
			if nums[s] < nums[m] {
				return binarySearch(m+1, e, target, nums)
			} else if nums[s] == nums[m] {
				if binarySearch(s, m-1, target, nums) {
					return true
				} else {
					return binarySearch(m+1, e, target, nums)
				}
			} else {
				return binarySearch(s, m-1, target, nums)
			}
		}
	}
}

func main() {
	//nums := []int{1, 1, 3, 1}
	//target := 0
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 4
	//nums := []int{3, 1}
	//target := 1
	fmt.Println(search(nums, target))
}
