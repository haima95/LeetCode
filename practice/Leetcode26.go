package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	count := 1
	last := nums[0]
	for i:=1;i<n;i++ {
		if nums[i] == last {
			continue
		}
		nums[count] = nums[i]
		count ++
		last = nums[i]
	}
	return count
}

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
