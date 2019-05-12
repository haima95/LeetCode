package main

import "fmt"

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}else {
			return false
		}
	}
	var first_match bool
	if len(s) > 0 {
		if s[0] == p[0] || p[0] == '.' {
			first_match = true
		}else {
			first_match = false
		}
	}

	if len(p) > 1 && p[1] == '*' {
		return isMatch(s,p[2:]) || first_match && isMatch(s[1:],p)
	}else {
		return first_match && isMatch(s[1:], p[1:])
	}
}

func main() {
	s := "aab"
	p := "c*a*b"
	fmt.Println(isMatch(s,p))
}
