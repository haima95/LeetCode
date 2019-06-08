package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{{}}
	sort.Ints(candidates)
	return getCombination(candidates, target, result)
}

func getCombination(candidates []int, target int, org [][]int) (result [][]int) {
	if target <= 0 {
		return result
	}
	for _, v := range candidates {
		if v < target {
			temp := [][]int{}
			for _, v1 := range org {
				n := len(v1)
				if n == 0 || (n > 0 && v1[n-1] <= v) {
					temp2 := []int{}
					for i := 0; i < len(v1); i++ {
						temp2 = append(temp2, v1[i])
					}
					temp2 = append(temp2, v)
					temp = append(temp, temp2)
				}
			}
			result = append(result, getCombination(candidates, target-v, temp)...)
		} else if v == target {
			for _, v2 := range org {
				n := len(v2)
				if n == 0 || (n > 0 && v2[n-1] <= v) {
					temp := []int{}
					for i := 0; i < len(v2); i++ {
						temp = append(temp, v2[i])
					}
					temp = append(temp, v)
					result = append(result, temp)
				}
			}
		}
	}
	//fmt.Println(result)
	return result
}

func main() {
	c := []int{7, 3, 2}
	target := 18
	fmt.Printf("%v\n", combinationSum(c, target))
}
