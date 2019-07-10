package main

import "fmt"

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	temp1 := 1
	temp2 := 2
	for i := 3; i < n+1; i++ {
		temp1, temp2 = temp2, temp1
		temp2 = temp1 + temp2
	}
	return temp2
}

func main() {
	n := 5
	fmt.Println(climbStairs(n))
}
