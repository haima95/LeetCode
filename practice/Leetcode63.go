package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m < 1 {
		return 0
	}
	n := len(obstacleGrid[0])
	temp := make([][]int, m)
	for i := 0; i < m; i++ {
		temp[i] = make([]int, n)
		if i == 0 {
			if obstacleGrid[i][0] == 0 {
				temp[0][0] = 1
			}
		} else {
			if obstacleGrid[i][0] == 0 {
				temp[i][0] = temp[i-1][0]
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if i > 0 {
				if obstacleGrid[i][j] == 0 {
					temp[i][j] = temp[i][j-1] + temp[i-1][j]
				} else {
					temp[i][j] = 0
				}
			} else {
				if obstacleGrid[i][j] == 0 {
					temp[i][j] = temp[i][j-1]
				} else {
					temp[i][j] = 0
				}
			}
		}
	}
	return temp[m-1][n-1]
}

func main() {
	t := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	//t := [][]int{{1}, {0}}

	fmt.Println(uniquePathsWithObstacles(t))
}
