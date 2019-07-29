package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{{}}
	}
	binarySort(nums, 0, n-1)
	result := [][]int{{}}
	temp := []int{}
	var m int
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			temp = append(temp, nums[i])
		} else {
			m = len(result)
			temp = []int{nums[i]}
		}
		for i := 0; i < m; i++ {
			t := []int{}
			t = append(t, result[i]...)
			result = append(result, append(t, temp...))
		}
	}
	return result
}

func binarySort(nums []int, start, end int) {
	if start >= end {
		return
	}
	temp := nums[end]
	s, e := start, end-1
	for s <= e {
		if nums[s] <= temp {
			s++
			continue
		}
		if nums[e] >= temp {
			e--
			continue
		}
		nums[s], nums[e] = nums[e], nums[s]
	}
	if nums[s] > nums[end] {
		nums[s], nums[end] = nums[end], nums[s]
	}
	binarySort(nums, start, s-1)
	binarySort(nums, s+1, end)
}

func main() {
	k := []int{9, 0, 3, 5, 7}
	r := subsetsWithDup(k)
	fmt.Println(r)
}
