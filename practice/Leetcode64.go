package main

import "fmt"

func minPathSum(grid [][]int) int {
	n := len(grid)
	if n < 1 {
		return 0
	}
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				if j == 0 {
					continue
				}
				grid[i][j] += grid[i][j-1]
			} else {
				if j == 0 {
					grid[i][j] += grid[i-1][j]
				} else {
					if grid[i][j-1] < grid[i-1][j] {
						grid[i][j] += grid[i][j-1]
					} else {
						grid[i][j] += grid[i-1][j]
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%2d ", grid[i][j])
		}
		fmt.Println()
	}
	return grid[n-1][m-1]
}

func main() {
	t := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(t))
}
