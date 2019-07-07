package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	var pre byte
	var hasE bool
	//var hasSign bool
	var hasNum bool
	var hasDot bool
	s = strings.Trim(s, " ")
	for _, item := range []byte(s) {
		if item == ' ' {
			return false
		}
		if (item >= '0' && item <= '9') || item == 'e' || item == '.' || item == '-' || item == '+' {
			if pre == 0 {
				if item == '-' || item == '+' {
					//hasSign = true
				} else if item == '.' {
					hasDot = true
				} else if item == 'e' {
					return false
				} else if item >= '0' && item <= '9' {
					hasNum = true
				}
			} else {
				if item == '-' || item == '+' {
					if pre != 'e' {
						return false
					}
				} else if item == 'e' {
					if hasE || ((pre < '0' || pre > '9') && !(pre == '.' && hasNum)) {
						return false
					}
					hasE = true
				} else if item == '.' {
					if hasE || hasDot || pre == 'e' || pre == '.' {
						return false
					}
					hasDot = true
				} else if item >= '0' && item <= '9' {
					hasNum = true
				}
			}
			pre = item
		} else {
			return false
		}
	}
	if pre == '+' || pre == '-' || pre == 'e' || !hasNum {
		return false
	}
	return true
}

type result2 struct {
	data string
	re   bool
}

func main() {
	r := []result2{{"0", true}, {"0.1", true}, {"+.8", true}, {".e1", false},
		{"abc", false}, {"1 a", false}, {"2e10", true}, {"34.e8", true},
		{"-90e3", true}, {"1e", false}, {"e3", false}, {"82e1113e333", false},
		{"6e-1", true}, {"99e2.5", false}, {"53.5e93", true}, {"1. 1", false},
		{"--6", false}, {"-+3", false}, {"95a54e53", false}, {"1 ", true},
		{".1", true}, {".", false}, {"1.0", true}, {"1.", true}}
	///r := []result2{{"-90e3", true}, {".1", true}}
	for _, item := range r {
		if item.re == isNumber(item.data) {
			fmt.Printf("ok : %+v\n", item)
		} else {
			fmt.Printf("fail : %+v\n", item)
		}
	}
}
