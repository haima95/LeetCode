package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	re :=[][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i:=0;i<n-2;i++ {
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		s := i+1
		e := n-1
		sum := -nums[i]
		for s < e {
			if nums[s] + nums[e] > sum {
				e --
			}else if nums[s] + nums[e] < sum {
				s ++
			}else {
				if !(s > 0 && nums[s]== nums[s-1]) || !(e < n-1 && nums[e] == nums[e+1]) {
					re = append(re,[]int{nums[i],nums[s],nums[e]})
				}
				s++
				e--
			}
		}
	}
	return re
}

func threeSum2(nums []int) [][]int {
	re :=[][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i:=0;i<n-1;i++ {
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		temp := make(map[int]int)
		for j:=i+1;j<n;j++ {
			if _,ok := temp[nums[j]];ok {
				re = append(re,[]int{nums[i],-nums[i]-nums[j],nums[j]})
			}else {
				temp[-nums[i]-nums[j]] = 1
			}
		}
	}
	count := len(re)
	for i:=0;i<count-1;i++ {
		if re[i][0]==re[i+1][0] && re[i][1]==re[i+1][1] && re[i][2]==re[i+1][2] {
			for j:=i;j<count-1;j++ {
				re[j][0],re[j][1],re[j][2] = re[j+1][0],re[j+1][1],re[j+1][2]
			}
			i--
			count --
		}
	}
	return re[:count]
}


func main() {
	//[[-4 -2 6] [-4 0 4] [-4 1 3] [-4 2 2] [-2 -2 4] [-2 0 2]]
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSum2(nums))
}
