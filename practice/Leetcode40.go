package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	fmt.Println(candidates)
	temp := getCandidates(candidates, [][]int{{}}, target)
	temp2 := make(map[string][]int)
	for _, v := range temp {
		sort.Ints(v)
		temp2[fmt.Sprint(v)] = v
	}
	fmt.Println(temp)
	var result [][]int
	for _, v := range temp2 {
		result = append(result, v)
	}
	return result

}

func getCandidates(candidates []int, org [][]int, target int) (result [][]int) {
	if target <= 0 {
		return result
	}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] < target {
			var temp [][]int
			for _, v1 := range org {
				if len(v1) == 0 || v1[len(v1)-1] <= candidates[i] {
					item := append(v1, candidates[i])
					temp = append(temp, item)
				}
			}
			result = append(result, getCandidates(candidates[i+1:], temp, target-candidates[i])...)
		} else if candidates[i] == target {
			for _, v1 := range org {
				if len(v1) == 0 || v1[len(v1)-1] <= candidates[i] {
					item := append(v1, candidates[i])
					result = append(result, item)
				}
			}
		}
	}
	return
}

func main() {
	c := []int{10, 1, 2, 7, 6, 1, 5}
	t := 8
	fmt.Println(combinationSum2(c, t))
}
