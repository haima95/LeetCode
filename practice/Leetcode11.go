package main

import "fmt"

func maxArea(height []int) int {
	result := 0
	for i,j := 0,len(height)-1;i<j; {
		temp := (j-i)*height[i]
		if height[i] > height[j] {
			temp = (j-i)*height[j]
			if temp > result {
				result = temp
			}
			j --
		}else {
			if temp > result {
				result = temp
			}
			i ++
		}
	}
	return result
}

func main() {
	a := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(a))
}
