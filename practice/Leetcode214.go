package main

import "fmt"

func shortestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	min := n - 1

	// 偶数
	even := n / 2
	for i := even; i > -1; i-- {
		l, r := i-1, i
		for l > -1 && s[l] == s[r] {
			l, r = l-1, r+1
		}
		if l == -1 && n-r < min {
			min = n - r
			break
		}
	}
	// 奇数
	odd := n/2 - 1
	for i := odd; i > -1; i-- {
		l, r := i, i+2
		for l > -1 && r < n && s[l] == s[r] {
			l, r = l-1, r+1
		}
		if l == -1 && n-r < min {
			min = n - r
			break
		}
	}
	result := make([]byte, min)
	for i := 0; i < min; i++ {
		result[i] = s[n-1-i]
	}
	result = append(result, s...)
	return string(result)
}

func main() {
	s := ""
	fmt.Println(shortestPalindrome(s))
}
