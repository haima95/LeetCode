package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
	if n1 == 0 {
		return n2
	}
	if n2 == 0 {
		return n1
	}

	result := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		result[i] = make([]int, n2+1)
	}

	for i := 0; i < n1+1; i++ {
		for j := 0; j < n2+1; j++ {
			if j == 0 {
				result[i][j] = i
			} else if i == 0 {
				result[i][j] = j
			} else if word1[i-1] == word2[j-1] {
				result[i][j] = result[i-1][j-1]
			} else {
				result[i][j] = min(result[i-1][j-1], result[i-1][j], result[i][j-1]) + 1
			}
		}
	}
	return result[n1][n2]
}

func min(a, b, c int) int {
	if a < b {
		a, b = b, a
	}
	if c > b {
		return b
	} else {
		return c
	}
}

func main() {
	word1 := "horse"
	word2 := "r"
	fmt.Println(minDistance(word1, word2))
}
