package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	result := 0
	n := 0
	for i:=1;i<len(s);i++ {
		for j:= n;j<i;j++ {
			if s[i] == s[j] {
				if i-n > result {
					result = i-n
				}
				n ++
				i--
			}
		}
	}
	fmt.Println(result)
	if len(s)-n > result {
		result = len(s)-n
	}
	return result
}

func main(){
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
