package main

import "fmt"

func twoSum(nums []int,target int) []int{
	temp := make(map[int]int)
	for i,v := range nums {
		c := target - v
		if j,ok := temp[c];ok {
			return []int{j,i}
		}
		temp[v] = i
	}
	return []int{}
}
func main(){
	nums := []int{2,7,11,15}
	target := 9
	fmt.Printf("%v\n",twoSum(nums,target))
}


