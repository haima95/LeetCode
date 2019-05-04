package main

import "fmt"

func myAtoi(str string) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -(1<<31)
	if len(str) <= 0 {
		return 0
	}
	var result int64 = 0
	isFirst := true
	isPositive := true
	for _,v := range []byte(str) {
		if isFirst {
			if v == ' '{
				continue
			}else if v == '-' {
				isFirst = false
				isPositive = false
			}else if v == '+' {
				isFirst = false
			}else if v > '9' || v < '0' {
				break
			}else {
				result = int64(v-'0')
				isFirst = false
			}
		}else if !isFirst {
			if v > '9' || v < '0' {
				break
			}else {
				result = result * 10 + int64(v-'0')
				if isPositive && result > INT_MAX {
					return INT_MAX
				}
				if !isPositive && -result < INT_MIN {
					return INT_MIN
				}
			}
		}
	}
	if !isPositive {
		result = -result
	}
	return int(result)
}

func main() {
	s := "9223372036854775808"
	fmt.Println(myAtoi(s))
}
