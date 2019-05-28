package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	for _,v := range nums {
		if v == val {
			continue
		}
		nums[result] = v
		result ++
	}
	return result
}

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	fmt.Println(removeElement(nums,val))
	fmt.Println(nums)
}
