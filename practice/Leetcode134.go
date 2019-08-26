package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		if gas[i] >= cost[i] {
			start := gas[i]
			j := 1
			for j < n {
				start += gas[(i+j)%n] - cost[(i+j-1)%n]
				if start < cost[(i+j)%n] {
					break
				}
				j++
			}
			if j == n {
				return i
			}
		}
	}
	return -1
}

func main() {
	gas := []int{2}
	cost := []int{3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
