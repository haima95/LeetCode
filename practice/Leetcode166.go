package main

import (
	"fmt"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return "invalid"
	}
	if numerator == 0 {
		return "0"
	}
	flag := true
	if numerator < 0 {
		flag = !flag
		numerator = -numerator
	}
	if denominator < 0 {
		flag = !flag
		denominator = -denominator
	}
	temp := make(map[string]int)
	var result []string
	local := -1
	for true {
		t := numerator / denominator
		numerator = numerator % denominator
		key := fmt.Sprintf("%d_%d", t, numerator)
		if v, ok := temp[key]; ok {
			local = v
			break
		} else {
			if len(result) == 0 {
				if numerator == 0 {
					result = append(result, fmt.Sprintf("%d", t))
					break
				}
				result = append(result, fmt.Sprintf("%d.", t))
			} else {
				temp[key] = len(result)
				result = append(result, fmt.Sprintf("%d", t))
			}
			if numerator == 0 {
				break
			}
		}
		numerator *= 10

	}
	r := ""
	if !flag {
		r = "-"
	}
	if numerator == 0 {
		return r + strings.Join(result, "")
	}
	return r + strings.Join(result[:local], "") + "(" + strings.Join(result[local:], "") + ")"

}

func main() {
	n := -100
	d := 12
	fmt.Println(fractionToDecimal(n, d))
}
