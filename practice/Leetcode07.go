package main

import "fmt"

func reverse(x int) int {
	flag := true    // 默认是正书
	if x < 0 {
		flag = false
		x = -x
	}
	temp := [12]int{}
	n := 0
	for x > 0 {
		temp[n] = x%10
		n ++
		x /= 10
	}
	result := 0
	for i:=0;i<n;i++ {
		result = result*10+temp[i]
	}
	if !flag {
		result = -result
	}
	if result > 1 << 31 || -result >= 1<<31 {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(1534236469))
}
