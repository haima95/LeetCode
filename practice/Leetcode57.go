package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	var i int
	for i = 0; i < len(intervals) && intervals[i][0] < newInterval[0]; i++ {
	}
	var arrTemp [][]int
	if i > 0 {
		arrTemp = append(arrTemp, intervals[:i]...)
	}
	arrTemp = append(arrTemp, newInterval)

	if i < len(intervals) {
		arrTemp = append(arrTemp, intervals[i:]...)
	}
	var result [][]int
	temp := arrTemp[0]
	for i := 1; i < len(arrTemp); i++ {
		if arrTemp[i][0] > temp[1] {
			result = append(result, temp)
			temp = arrTemp[i]
		} else if arrTemp[i][0] <= temp[1] {
			if arrTemp[i][1] > temp[1] {
				temp[1] = arrTemp[i][1]
			}
		}
	}
	result = append(result, temp)
	return result
}

func main() {
	intervals := [][]int{}
	newIntervals := []int{14, 8}
	fmt.Println(insert(intervals, newIntervals))
}
