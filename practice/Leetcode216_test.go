package main

import (
	"fmt"
	"testing"
)

func combinationSum3(k int, n int) [][]int {
	if n < k || n > 9*n {
		return [][]int{}
	}
	result := [][]int{}
	for i := 1; i < 10; i++ {
		temp := createSum([]int{i}, n-i, k-1)
		if len(temp) > 0 {
			result = append(result, temp...)
		}
	}
	return result
}

func createSum(arr []int, n, k int) [][]int {
	if k == 0 {
		if n == 0 {
			return [][]int{arr}
		} else {
			return [][]int{}
		}
	}
	t := len(arr)
	if arr[t-1] >= n {
		return [][]int{}
	}
	var result [][]int
	for j := arr[t-1] + 1; j < 10; j++ {
		temp := append([]int{}, arr...)
		temp = append(temp, j)
		tt := createSum(temp, n-j, k-1)
		if len(tt) > 0 {
			result = append(result, tt...)
		}

	}
	return result
}

func Test216(t *testing.T) {
	k := 2
	n := 3
	fmt.Println(combinationSum3(k, n))
}
