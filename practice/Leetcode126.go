package main

import (
	"fmt"
	"math"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	words := make(map[string]struct{})
	for _, w := range wordList {
		words[w] = struct{}{}
	}
	if _, ok := words[endWord]; !ok {
		return [][]string{}
	}

	result := make([][]string, 0)

	level := 1
	length := math.MaxInt32
	visited := make(map[string]struct{})
	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})
	for len(queue) > 0 {
		ladder := queue[0]
		queue = queue[1:]

		if len(ladder) > level {
			for w := range visited {
				delete(words, w)
			}
			visited = make(map[string]struct{})
		}
		lastWord := ladder[len(ladder)-1]
		if len(ladder) > length {
			break
		} else {
			level = len(ladder)
			if lastWord == endWord {
				result = append(result, ladder)
				length = len(ladder)
			}
		}

		for i := 0; i < len(lastWord); i++ {
			for b := 'a'; b <= 'z'; b++ {
				newWord := lastWord[:i] + string(b) + lastWord[i+1:]
				if _, ok := words[newWord]; !ok || newWord == lastWord {
					continue
				}
				visited[newWord] = struct{}{}
				copyPath := make([]string, 0, len(ladder)+1)
				copyPath = append(copyPath, ladder...)
				copyPath = append(copyPath, newWord)
				queue = append(queue, copyPath)
			}
		}
	}
	return result
}

func main() {
	//beginWord := "a"
	//endWord := "c"
	//wordList := []string{"a", "b", "c"}

	beginWord := "qa"
	endWord := "sq"
	wordList := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb",
		"sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba",
		"to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or",
		"rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb",
		"ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi",
		"qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne",
		"mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}
