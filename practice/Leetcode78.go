package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{{}}

	for k := 1; k <= len(nums); k++ {
		createSubs(&result, k, []int{nums[k-1]}, nums)
	}
	return result
}

func createSubs(result *[][]int, k int, org, nums []int) {
	*result = append(*result, org)
	if k == len(nums) {
		return
	}
	for i := k; i < len(nums); i++ {
		var temp []int
		temp = append(temp, org...)
		temp = append(temp, nums[i])
		createSubs(result, i+1, temp, nums)
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(subsets(nums))
	fmt.Println(len(subsets(nums)))
}
