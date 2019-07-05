package main

import "fmt"

func uniquePaths(m int, n int) int {
	temp := make([][]int, m)
	for i := 0; i < m; i++ {
		temp[i] = make([]int, n)
		temp[i][0] = 1
	}

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if i > 0 {
				temp[i][j] = temp[i][j-1] + temp[i-1][j]
			} else {
				temp[i][j] = temp[i][j-1]
			}
		}
	}
	return temp[m-1][n-1]
}

func main() {
	m, n := 7, 3
	fmt.Println(uniquePaths(m, n))

}
