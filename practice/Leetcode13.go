package main

import "fmt"

func romanToInt(s string) int {
	result := 0
	len := len(s)
	for i:=0;i<len;i++ {
		switch s[i] {
		case 'I':
			result += 1
			if i+1 < len {
				if  s[i+1] == 'V' {
					result += 3
					i++
				}else if s[i+1] == 'X' {
					result += 8
					i++
				}
			}
		case 'V':
			result += 5
		case 'X':
			result += 10
			if i+1 < len {
				if  s[i+1] == 'L' {
					result += 30
					i++
				}else if s[i+1] == 'C' {
					result += 80
					i++
				}
			}
		case 'L':
			result += 50
		case 'C':
			result += 100
			if i+1 < len {
				if  s[i+1] == 'D' {
					result += 300
					i++
				}else if s[i+1] == 'M' {
					result += 800
					i++
				}
			}
		case 'D':
			result += 500
		case 'M':
			result += 1000
		}
	}	
	return result
}

func main() {
	s := "DCXXI"
	fmt.Println(romanToInt(s))
}
