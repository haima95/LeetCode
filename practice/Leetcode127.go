package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	var i int
	for i = 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			break
		}
	}
	if i == len(wordList) {
		return 0
	}
	level := 1
	n := 0
	t := []string{beginWord}
	s := wordList
	for true {
		t1 := make([]string, 0)
		n = len(s)
		for i = 0; i < len(t); i++ {
			s1 := make([]string, 0)
			for j := 0; j < len(s); j++ {
				if t[i] == s[j] {
					continue
				}
				if isMostMatch(t[i], endWord) {
					level++
					return level
				}
				if isMostMatch(t[i], s[j]) {
					t1 = append(t1, s[j])
				} else {
					s1 = append(s1, s[j])
				}
			}
			s = s1
		}
		if n == len(s) {
			break
		}
		t = t1

		level++
	}
	return 0
}

func isMostMatch(target, resource string) bool {
	var diff int
	for i := 0; i < len(resource); i++ {
		if target[i] != resource[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}

func main() {
	//beginWord := "qa"
	//endWord := "sq"
	//wordList := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb",
	//	"sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba",
	//	"to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or",
	//	"rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb",
	//	"ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi",
	//	"qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne",
	//	"mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}

	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
