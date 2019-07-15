package main

import "fmt"

func combine(n int, k int) [][]int {
	if n < k || k == 0 {
		return [][]int{}
	}
	temp := [][]int{}
	for i := 1; i <= n-k+1; i++ {
		t := []int{i}
		temp = append(temp, createQueue(i+1, n, k-1, t)...)
	}
	return temp
}

func createQueue(start, end, k int, t []int) [][]int {
	var result [][]int
	if k == 0 {
		result = append(result, t)
		return result
	}
	for i := start; i <= end-k+1; i++ {
		var temp []int
		temp = append(temp, t...)
		temp = append(temp, i)
		result = append(result, createQueue(i+1, end, k-1, temp)...)
	}
	return result
}

func main() {
	n, k := 6, 4
	fmt.Println(combine(n, k))
}
