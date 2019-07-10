package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	t := x / 2
	for {
		if t*t <= x && (t+1)*(t+1) > x {
			return t
		} else if t*t > x {
			t /= 2
		} else if (t+1)*(t+1) <= x {
			t += 1
		}
	}
}

func mySqrt2(x int) int {
	l, h := 0, x
	for l < h {
		mid := (l+h)/2 + 1
		if x/mid >= mid {
			l = mid
		} else {
			h = mid - 1
		}
	}
	return l
}

func main() {
	x := 100
	fmt.Println(mySqrt2(x))
}
