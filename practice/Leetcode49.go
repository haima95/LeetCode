package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	temp := make(map[string][]string)
	for _, str := range strs {
		it := []byte(str)
		sort.Slice(it, func(i, j int) bool {
			if it[i] <= it[j] {
				return true
			}
			return false
		})
		if _, ok := temp[string(it)]; ok {
			temp[string(it)] = append(temp[string(it)], str)
		} else {
			temp[string(it)] = []string{str}
		}
	}
	var result [][]string
	for _, v := range temp {
		k := []string{}
		for _, v1 := range v {
			k = append(k, v1)
		}
		result = append(result, k)
	}
	return result
}

func main() {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(s))
}
