package main

import "fmt"

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	dic := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x := 0
	y := -1
	t := 0
	for i := 1; i <= n*n; i++ {
		if x+dic[t][0] > -1 && x+dic[t][0] < n && y+dic[t][1] > -1 && y+dic[t][1] < n && result[x+dic[t][0]][y+dic[t][1]] == 0 {
			result[x+dic[t][0]][y+dic[t][1]] = i
			x = x + dic[t][0]
			y = y + dic[t][1]
		} else {
			t = (t + 1) % 4
			i--
		}
	}
	return result
}

func main() {
	//n := 3
	for i := 0; i < 10; i++ {
		result := generateMatrix(i)
		for j := 0; j < i; j++ {
			for k := 0; k < i; k++ {
				fmt.Printf("%2d ", result[j][k])
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
}
