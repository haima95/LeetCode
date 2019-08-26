package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	if n < 2 {
		return n
	}
	temp := make([]int, n)
	temp[0] = 1
	numCount := 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			temp[i] = temp[i-1] + 1
			numCount += (temp[i-1] + 1)
		} else {
			numCount++
			temp[i] = 1
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", temp[i])
	}
	fmt.Println()
	fmt.Println(numCount)
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && temp[i] >= temp[i-1] {
			numCount += temp[i] - temp[i-1] + 1
			temp[i-1] = temp[i] + 1
		}
	}
	return numCount
}

func main() {
	n := []int{1, 0, 2}
	fmt.Println(candy(n))
}
