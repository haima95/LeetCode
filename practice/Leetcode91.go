package main

import "fmt"

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	temp := make([]int, n)
	if s[0] == '0' {
		temp[0] = 0
	} else {
		temp[0] = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			if s[i-1] == '0' || s[i-1] > '2' {
				return 0
			} else if s[i-1] < '3' {
				if i > 1 {
					temp[i] = temp[i-2]
				} else {
					temp[i] = 1
				}
			}
		} else if s[i] > '0' && s[i] < '7' {
			temp[i] = temp[i-1]
			if s[i-1] > '0' && s[i-1] < '3' {
				if i > 1 {
					temp[i] += temp[i-2]
				} else {
					temp[i] += 1
				}
			}
		} else {
			temp[i] = temp[i-1]
			if s[i-1] > '0' && s[i-1] < '2' {
				if i > 1 {
					temp[i] += temp[i-2]
				} else {
					temp[i] += 1
				}
			}
		}
	}
	return temp[n-1]
}

func main() {
	// s := "110"   // 1
	// s := "27"   // 1
	// s := "0" // 0
	// s := "12"   // 2
	// s := "226"
	fmt.Println(numDecodings(s))
}
