package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	m := len(words)
	result := []int{}
	if m == 0 {
		return result
	}
	n := len(words[0])
	match := make(map[string]int)
	for _, v := range words {
		if _, ok := match[v]; ok {
			match[v]++
		} else {
			match[v] = 1
		}
	}
	for i := 0; i < len(s)-m*n+1; i++ {
		temp := make(map[string]int)
		for j := i; j < m*n-n+1+i; j += n {
			v := s[j : j+n]
			if _, ok := temp[v]; ok {
				temp[v]++
			} else {
				temp[v] = 1
			}
		}
		flag := true
		for k, v := range match {
			if v != temp[k] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	//s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	//words := []string{"fooo", "barr", "wing", "ding", "wing"}
	//words := []string{"word", "good", "best", "word"}
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
}
