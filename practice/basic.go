package main

import "fmt"

func printIntArray(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%2d ", arr[i][j])
		}
		fmt.Println()
	}
}
