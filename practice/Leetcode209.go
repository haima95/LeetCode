package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sum := nums[0]
	if sum >= s {
		return 1
	}
	n := 1
	min := len(nums) + 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= s {
			return 1
		}
		n++
		sum += nums[i]
		if sum >= s {
			sum -= nums[i-n+1]
			n--
			for j := n - 1; j > 0 && sum >= s; j-- {
				sum -= nums[i-j]
				n--
			}
			if min > n+1 {
				min = n + 1
			}
		}
	}
	if min > len(nums) {
		return 0
	}
	return min
}

func main() {
	t := []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}
	s := 20
	fmt.Println(minSubArrayLen(s, t))
}
