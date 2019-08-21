package main

import "fmt"

func partition(s string) [][]string {
	var result [][]string
	if len(s) == 0 {
		return result
	}
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			temp := make([]string, 1)
			temp[0] = s[:i]
			partition2(&result, temp, s[i:])
		}
	}
	return result
}

func partition2(result *[][]string, s1 []string, s2 string) {
	if len(s2) == 0 {
		*result = append(*result, s1)
		return
	}
	for i := 1; i <= len(s2); i++ {
		if isPalindrome(s2[:i]) {
			temp := make([]string, len(s1)+1)
			copy(temp, s1)
			temp[len(s1)] = s2[:i]
			partition2(result, temp, s2[i:])
		}
	}
}

func isPalindrome(s string) bool {
	i, n := 0, len(s)-1
	for i < n {
		if s[i] != s[n] {
			return false
		}
		i++
		n--
	}
	return true
}

func main() {
	s := "aabaa"
	fmt.Println(partition(s))
}
