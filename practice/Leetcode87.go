package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)
	alpha := make([]int, 26)
	for i := 0; i < n; i++ {
		alpha[s1[i]-'a']++
	}
	for i := 0; i < n; i++ {
		alpha[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alpha[i] != 0 {
			return false
		}
	}
	for i := 1; i < n; i++ {
		if (isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])) || (isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i])) {
			return true
		}
	}
	return false
}

func main() {
	s1 := "abc"
	s2 := "bca"
	fmt.Println(isScramble(s1, s2))
}
