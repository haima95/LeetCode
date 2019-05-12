package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	temp := make(map[byte]int)
	flag := true
	for flag {
		if i >= len(strs[0]) {
			break
		}
		temp[strs[0][i]] = 1
		for _,v := range strs[1:] {
			if i >= len(v) {
				break
			}
			if _,ok := temp[v[i]];ok {
				temp[v[i]]++
			}
		}
		if c,_ :=temp[strs[0][i]];c != len(strs){
			flag = false
		}else {
			i ++
		}
	}
	if i == 0 {
		return ""
	}
	return strs[0][:i]
}

func main() {
	s := []string{"1","","d"}
	fmt.Println(longestCommonPrefix(s))
}
