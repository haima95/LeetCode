package main

import "fmt"

func multiply(num1 string, num2 string) string {

	if len(num1) < 1 || len(num2) < 1 || num1 == "0" || num2 == "0" {
		return "0"
	}
	n := len(num1)
	m := len(num2)
	var temp uint8

	dig := make([]uint8, n+m)
	for i := 0; i < n; i++ {
		temp = 0
		x := num1[n-i-1] - '0'
		for j := 0; j < m; j++ {
			y := num2[m-j-1] - '0'
			t := x*y + temp
			dig[i+j] += t % 10
			t1 := i + j
			for dig[t1] > 9 {
				temp1 := dig[t1]
				dig[t1] = temp1 % 10
				dig[t1+1] += temp1 / 10
				t1++
			}
			temp = t / 10
		}
		if temp > 0 {
			dig[i+m] += temp
		}
	}
	var r []byte

	for i := len(dig) - 1; i > -1; i-- {
		if i == len(dig)-1 && dig[i] == 0 {
			continue
		}
		r = append(r, '0'+dig[i])
	}
	return string(r)
}

func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}
