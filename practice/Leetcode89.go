package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	result := []int{0, 1}
	for i := 2; i <= n; i++ {
		m := 1 << uint(i)
		for j := 0; j < m; j++ {

		}
	}
	return result
}

func grayCode2(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	result := []int{0, 1}
	for i := 2; i <= n; i++ {
		m := len(result)
		for j := m - 1; j >= 0; j-- {
			result = append(result, result[j]+(1<<uint(i-1)))
		}
	}
	return result
}

func main() {
	n := 4
	t := grayCode2(n)
	for i := 0; i < len(t); i++ {
		fmt.Printf("%5b\n", t[i])
	}
}
