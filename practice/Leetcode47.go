package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 2 {
		return append(result, nums)
	}
	result = append(result, []int{nums[0]})
	for i := 1; i < len(nums); i++ {
		result = dep2(result, nums[i])
	}
	temp := make(map[string][]int)
	for i := 0; i < len(result); i++ {
		str := ""
		for j := 0; j < len(result[i]); j++ {
			str += fmt.Sprint("%d_", result[i][j])
		}
		temp[str] = result[i]
	}
	result2 := [][]int{}
	for _, v := range temp {
		result2 = append(result2, v)
	}
	return result2
}

func dep2(nums [][]int, n int) [][]int {
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
	v := []int{1, 1, 2}
	fmt.Println(permuteUnique(v))
}
