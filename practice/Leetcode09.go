package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := 0
	y := x
	for x > 0 {
		temp = temp*10 + x%10
		x /= 10
	}
	if temp == y {
		return true
	}
	return false
}

func main(){
	fmt.Println(isPalindrome(121))
}
