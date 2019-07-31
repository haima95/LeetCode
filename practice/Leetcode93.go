package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	var result []string
	for i := 1; i < 4; i++ {
		if i > 1 && s[0] == '0' {
			return result
		}
		t, _ := strconv.Atoi(s[:i])
		if t > 255 {
			return result
		}
		if len(s)-i < 10 && len(s)-i > 2 {
			temp := getIpAddress(fmt.Sprintf("%s.%s", s[:i], s[i:]), i+1, 2)
			result = append(result, temp...)
		}
	}
	return result
}

func getIpAddress(s string, f, n int) (result []string) {
	if n == 4 {
		if f < len(s) {
			t, _ := strconv.Atoi(s[f:])
			if (s[f] == '0' && len(s[f:]) > 1) || t > 255 {
				return
			} else {
				result = append(result, s)
			}
		}
		return
	}
	for i := f + 1; i < f+4 && i < len(s); i++ {
		if i > f+1 && s[f] == '0' {
			return result
		}
		t, _ := strconv.Atoi(s[f:i])
		if t > 255 {
			return result
		}
		if len(s)-i < (4-n)*3+1 && len(s)-i > (4-n)*1-1 {
			temp := getIpAddress(fmt.Sprintf("%s.%s", s[:i], s[i:]), i+1, n+1)
			result = append(result, temp...)
		}
	}
	return result
}

func main() {
	//s := "25525511135"
	//s := "010010"
	s := "123456732"
	fmt.Println(restoreIpAddresses(s))
}
