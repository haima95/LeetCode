package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	result := math.MinInt32
	cump, cumn := 1, 1
	var temp int
	for _, v := range nums {
		if v >= 0 {
			if v > cump*v {
				cump = v
			} else {
				cump = cump * v
			}
			if v < cumn*v {
				cumn = v
			} else {
				cumn = cumn * v
			}
		} else {
			if v > cumn*v {
				temp = v
			} else {
				temp = cumn * v
			}
			if v < cump*v {
				cumn = v
			} else {
				cumn = cump * v
			}
			cump = temp
		}
		if cump > result {
			result = cump
		}
	}
	return result
}

func main() {
	i := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(i))
}
