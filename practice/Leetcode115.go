package main

import "fmt"

func numDistinct(s string, t string) int {
	ns := len(s)
	nt := len(t)
	if ns < nt {
		return 0
	}
	if ns == 0 && nt == 0 {
		return 1
	}
	temp := make([][]int, nt+1)
	for i := 0; i < nt+1; i++ {
		temp[i] = make([]int, ns)
	}
	for i := 0; i < ns; i++ {
		temp[0][i] = 1
	}
	for i := 1; i < nt+1; i++ {
		for j := i - 1; j < ns; j++ {
			if j == 0 {
				if t[i-1] == s[j] {
					temp[i][j] = 1
				}
				continue
			}
			if t[i-1] == s[j] {
				temp[i][j] = temp[i-1][j-1] + temp[i][j-1]
			} else {
				temp[i][j] = temp[i][j-1]
			}

		}
	}
	fmt.Print("  ")
	for i := 0; i < ns; i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	for i := 1; i < nt+1; i++ {
		fmt.Printf("%c ", t[i-1])
		for j := 0; j < ns; j++ {
			fmt.Printf("%d ", temp[i][j])
		}
		fmt.Println()
	}
	return temp[nt][ns-1]
}

func main() {
	s := "b"
	t := "a"
	fmt.Println(numDistinct(s, t))
}
