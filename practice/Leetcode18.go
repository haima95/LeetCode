package main

import (
	"sort"
	"fmt"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var result [][]int
	for i:=0;i<n-3;i++ {
		if i>0 && nums[i]==nums[i-1] {
			continue
		}
		for j:=i+1;j<n-2;j++ {
			if j>i+1 && nums[j]==nums[j-1]{
				continue
			}
			s := j+1
			e := n-1
			temp := target - nums[i]-nums[j]
			for s < e {
				if nums[s]+nums[e] == temp {
					m := len(result)
					if m < 1 || !(result[m-1][0]==nums[i]&&result[m-1][1]==nums[j]&&result[m-1][2]==nums[s]&&result[m-1][3]==nums[e]) {
						result = append(result,[]int{nums[i],nums[j],nums[s],nums[e]})
					}
					s ++
					e --
				}else if nums[s]+nums[e] > temp {
					e --
				}else {
					s ++
				}
			}
		}
	}
	return result
}

func main() {
	//nums := []int{-1,0,1,2,-1,-4}
	nums := []int{0,0,0,0}
	//nums := []int{-5,-4,-3,-2,-1,0,0,1,2,3,4,5}
	//nums:= []int{0,4,-5,2,-2,4,2,-1,4}
	target := 0
	fmt.Println(fourSum(nums,target))
}
