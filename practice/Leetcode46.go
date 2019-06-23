package main

import "fmt"

func permute(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 2 {
		return append(result, nums)
	}
	result = append(result, []int{nums[0]})
	for i := 1; i < len(nums); i++ {
		result = dep(result, nums[i])
	}
	return result
}

func dep(nums [][]int, n int) [][]int {
	var result [][]int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i])+1; j++ {
			var temp []int
			if j == 0 {
				temp = append(temp, n)
				temp = append(temp, nums[i][j:]...)
			} else if j == len(nums[i]) {
				temp = append(temp, nums[i][:j]...)
				temp = append(temp, n)
			} else {
				temp = append(temp, nums[i][:j]...)
				temp = append(temp, n)
				temp = append(temp, nums[i][j:]...)
			}
			result = append(result, temp)
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Printf("%v\n", permute(arr))
}
