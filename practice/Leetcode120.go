package main

import "fmt"

type Node struct {
	i, j, val int
	next      *Node
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	temp := make([]int, len(triangle[n-1]))
	temp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		temp1 := make([]int, len(temp))
		copy(temp1, temp)
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				temp1[0] += triangle[i][j]
			} else if j == len(triangle[i])-1 {
				temp1[j] = temp[j-1] + triangle[i][j]
			} else {
				if temp[j-1] > temp[j] {
					temp1[j] = temp[j] + triangle[i][j]
				} else {
					temp1[j] = temp[j-1] + triangle[i][j]
				}
			}
		}
		temp = temp1
	}
	max := temp[0]
	for i := 1; i < len(temp); i++ {
		if max > temp[i] {
			max = temp[i]
		}
	}
	return max
}

func main() {
	t := [][]int{{-1}, {2, 3}, {1, -1, -3}}
	fmt.Println(minimumTotal(t))
}
