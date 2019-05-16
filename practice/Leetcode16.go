package main

import (
	"sort"
	"fmt"
)

func threeSumClosest(nums []int, target int) int {
	dis := 1 << 31
	sum := 1 << 31
	n := len(nums)
	sort.Ints(nums)
	for i:=0;i<n-2;i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		s := i+1
		e := n-1
		for s < e {
			if nums[i]+nums[s]+nums[e] > target  {
				if nums[i]+nums[s]+nums[e] - target < dis {
					dis = nums[i]+nums[s]+nums[e] - target
					sum = nums[i]+nums[s]+nums[e]
				}
				e --
			}else if nums[i]+nums[s]+nums[e] < target {
				if target-nums[i]-nums[s]-nums[e] < dis {
					sum = nums[i]+nums[s]+nums[e]
					dis = target-nums[i]-nums[s]-nums[e]
				}
				s ++
			}else if nums[i]+nums[s]+nums[e] == target {
				return target
			}
		}
	}
	return sum
}

func main() {
	n := []int{0,1,2}
	fmt.Println(threeSumClosest(n,3))
}
