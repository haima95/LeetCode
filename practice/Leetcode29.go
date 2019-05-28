package main

import "fmt"

func divide(dividend int, divisor int) int {
	result := 0
	if dividend < 0 {
		if divisor < 0 {
			for dividend <= divisor {
				result ++
				dividend -= divisor
			}
		}else {
			for -dividend >= divisor {
				result ++
				dividend += divisor
			}
			result = -result
		}
	}else {
		if divisor < 0 {
			for dividend >= -divisor {
				result ++
				dividend += divisor
			}
			result = -result
		}else {
			for dividend >= divisor {
				result ++
				dividend -= divisor
			}
		}
	}
	if result > 1<<31-1 {
		result = 1<<31-1
	}else if result < -1<<31 {
		result = -1<<31
	}
	return result
}

func divide2(dividend int, divisor int) int {
	flag := true
	if dividend < 0 {
		dividend = -dividend
		flag = false
	}
	if divisor < 0 {
		divisor = -divisor
		flag = !flag
	}
	result := 0
	for i:=31;i>=0;i-- {
		temp := divisor << uint(i)
		for dividend >= temp {
			dividend -= temp
			result += 1<<uint(i)
		}
	}
	if flag {
		if result > 2147483647 {
			result = 2147483647
		}
	}else {
		if result > 2147483648 {
			result = -2147483648
		}else {
			result = -result
		}
	}
	return result
}

func main() {
	d1 := 7
	d2 := -3
	fmt.Println(divide2(d1,d2))
}
