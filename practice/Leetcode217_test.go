package main

import (
	"fmt"
	"testing"
)

func containsDuplicate(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	v := make(map[int]bool)
	for i := 0; i < n; i++ {
		if _, ok := v[nums[i]]; ok {
			return true
		} else {
			v[nums[i]] = true
		}
	}
	return false
}

func Test217(t *testing.T) {
	s := []int{3, 3}
	fmt.Println(containsDuplicate(s))
}
