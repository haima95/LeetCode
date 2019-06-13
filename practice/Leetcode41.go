package main

import "fmt"

func firstMissingPositive(nums []int) int {
	temp := make([]int, len(nums)+1)
	n := 1
	for _, v := range nums {
		if v == n {
			temp[v] = 1
			for n < len(nums)+1 && temp[n] != 0 {
				n = n + 1
			}
		} else if v > -1 && v < len(nums)+1 {
			temp[v] = 1
		}
	}
	return n
}

func main() {
	nums := []int{2, 1}
	fmt.Println(firstMissingPositive(nums))
}
