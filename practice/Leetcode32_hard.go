package main

import "fmt"

// 第一种方法，暴力法
func longestValidParentheses(s string) int {
	maxlen := 0
	n := len(s)
	for i := 0; i < n; i++ {
		for j := i + 2; j <= n; j += 2 {
			if isValid2(s[i:j]) {
				if maxlen < j-i {
					maxlen = j - i
				}
			}
		}
	}
	return maxlen
}

func isValid2(s string) bool {
	temp := []byte(s)
	n := 0
	for _, v := range []byte(s) {
		if v == '(' {
			temp[n] = v
			n++
		} else {
			if n > 0 && temp[n-1] == '(' {
				n--
			} else {
				return false
			}
		}
	}
	if n == 0 {
		return true
	} else {
		return false
	}
}

// 第二种方法，动态规划
func longestValidParentheses2(s string) int {
	maxlen := 0
	n := len(s)
	dp := []int{}
	for _, _ = range s {
		dp = append(dp, 0)
	}
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if dp[i] > maxlen {
			maxlen = dp[i]
		}
	}
	return maxlen
}

func main() {

	s := "(()())"
	fmt.Println(longestValidParentheses2(s))
}
