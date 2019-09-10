package main

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 1
	}
	n := len(dungeon[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	if dungeon[m-1][n-1] > 0 {
		dp[m-1][n-1] = 1
	} else {
		dp[m-1][n-1] = 1 - dungeon[m-1][n-1]
	}
	return calculate(dungeon, 0, 0, dp)
}

func calculate(dungeon [][]int, x, y int, dp [][]int) int {
	if x >= len(dungeon) || y >= len(dungeon[0]) {
		return math.MaxInt32
	}
	if dp[x][y] != 0 {
		return dp[x][y]
	}
	down := calculate(dungeon, x+1, y, dp)
	right := calculate(dungeon, x, y+1, dp)
	min := down
	if right < down {
		min = right
	}
	if min-dungeon[x][y] <= 0 {
		dp[x][y] = 1
	} else {
		dp[x][y] = min - dungeon[x][y]
	}
	return dp[x][y]
}

func main() {
}
