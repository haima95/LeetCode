package main

import "fmt"

func myPow(x float64, n int) float64 {
	var result float64 = 1
	if n == 0 {
		return 1.0000
	} else if n > 0 {
		temp := map[int]float64{1: x}
		i := 1
		for i <= n {
			result *= temp[i]
			temp[i*2] = temp[i] * temp[i]
			n -= i
			i = i * 2
		}
		for n != 0 {
			if n < i {
				i /= 2
			} else {
				result *= temp[i]
				n -= i
			}
		}
	} else {
		temp := map[int]float64{-1: 1 / x}
		i := -1
		for n <= i {
			result *= temp[i]
			temp[i*2] = temp[i] * temp[i]
			n -= i
			i = i * 2
		}
		for n != 0 {
			if n > i {
				i /= 2
			} else {
				result *= temp[i]
				n -= i
			}
		}
	}
	return result
}

func main() {
	t := 2.0
	for i := -10; i < 0; i++ {
		fmt.Println(t, "的", i, "次方为", myPow(t, i))
	}
}
