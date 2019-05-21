package main

import "fmt"

func isValid(s string) bool {
	lefts := []byte(s)
	n := 0
	for _,v := range []byte(s) {
		switch v {
		case '(','[','{':
			lefts[n] = v
			n++
		case ')':
			if n < 1 || lefts[n-1] != '(' {
				return false
			}else if lefts[n-1] == '(' {
				n --
			}
		case ']':
			if n < 1 || lefts[n-1] != '[' {
				return false
			}else if lefts[n-1] == '[' {
				n --
			}
		case '}':
			if n < 1 || lefts[n-1] != '{' {
				return false
			}else if lefts[n-1] == '{' {
				n --
			}
		}
	}
	if n == 0 {
		return true
	}else {
		return false
	}
}

func main() {
	s := "[[()]]{[]([)}]"
	fmt.Println(isValid(s))
}
