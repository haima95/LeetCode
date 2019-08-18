package main

import "fmt"

func isPalindrome(s string) bool {
	b := []byte(s)
	if len(b) < 2 {
		return true
	}
	f := 0

	e := len(b) - 1
	for f < e {
		if b[f] >= 'A' && b[f] <= 'Z' {
			b[f] += 'a' - 'A'
		}
		if b[e] >= 'A' && b[e] <= 'Z' {
			b[e] += 'a' - 'A'
		}
		if !(b[f] >= 'a' && b[f] <= 'z') && !(b[f] >= '0' && b[f] <= '9') {
			f++
			continue
		}
		if !(b[e] >= 'a' && b[e] <= 'z') && !(b[e] >= '0' && b[e] <= '9') {
			e--
			continue
		}
		if b[f] == b[e] {
			f++
			e--
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}
