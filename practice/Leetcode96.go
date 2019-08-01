package main

import "fmt"

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	temp := make([]int, n+1)
	temp[0] = 1
	temp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			temp[i] += temp[j-1] * temp[i-j]
		}
	}
	return temp[n]
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(numTrees(i))
	}
	// output
	//0
	//1
	//2
	//5
	//14
	//42
	//132
	//429
	//1430
	//4862

}
