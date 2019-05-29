package main

import (
	"fmt"
)

//var result []string
func generateParenthesis(n int) []string {
	var result []string
	generate(&result, "", 0, 0, n)
	return result
}

func generate(result *[]string, temp string, open, close, max int) {
	if len(temp) == 2*max {
		*result = append(*result, temp)
		return
	}
	if open < max {
		generate(result, temp+"(", open+1, close, max)
	}
	if close < open {
		generate(result, temp+")", open, close+1, max)
	}
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))

}
