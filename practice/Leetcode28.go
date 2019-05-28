package main

import "fmt"

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}
	if n < m {
		return -1
	}
	for i:=0;i<=n-m;i++ {
		flag := true
		for j:=0;j<m;j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

func main() {
	s := "ahello"
	n := "hello"
	fmt.Println(strStr(s,n))
}
