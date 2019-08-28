package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) []string {
	return df(s, wordDict, map[string][]string{})
}

func df(s string, wordDict []string, te map[string][]string) []string {
	if res, ok := te[s]; ok {
		return res
	}
	if len(s) == 0 {
		return []string{""}
	}
	var res []string
	for i := 0; i < len(wordDict); i++ {
		if len(wordDict[i]) <= len(s) && wordDict[i] == s[:len(wordDict[i])] {
			for _, str := range df(s[len(wordDict[i]):], wordDict, te) {
				if len(str) == 0 {
					res = append(res, wordDict[i])
				} else {
					res = append(res, wordDict[i]+" "+str)
				}
			}
		}
	}
	te[s] = res
	return res
}

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	w := []string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa", "ba"}
	t := wordBreak(s, w)
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
}
