package main

import "fmt"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	var result [][]int
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > temp[1] {
			result = append(result, temp)
			temp = intervals[i]
		} else if intervals[i][0] <= temp[1] {
			if intervals[i][1] > temp[1] {
				temp[1] = intervals[i][1]
			}
		}
	}
	result = append(result, temp)
	return result
}

func quickSort(intervals [][]int, start, end int) {
	if start >= end {
		return
	}
	temp := intervals[end]
	left := start
	right := end - 1
	for left <= right {
		if intervals[left][0] <= temp[0] {
			left++
			continue
		}
		if intervals[right][0] >= temp[0] {
			right--
			continue
		}
		intervals[left], intervals[right] = intervals[right], intervals[left]
	}
	intervals[left], intervals[end] = intervals[end], intervals[left]
	quickSort(intervals, start, left-1)
	quickSort(intervals, left+1, end)

}

func main() {
	input := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(input))
}
