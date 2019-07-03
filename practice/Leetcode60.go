package main

import "fmt"

func getPermutation(n int, k int) string {

	temp := make([]int, 10)
	t := []byte{'.', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	temp[1] = 1
	temp[2] = 1
	for i := 3; i <= n; i++ {
		temp[i] = (i - 1) * temp[i-1]
	}
	var result []byte
	for i := n; i > 0; i-- {
		te := (k - 1) / temp[i]
		j := 1
		k1 := 0
		for ; j < n+1; j++ {
			if t[j] == '.' {
				continue
			}
			if k1 == te {
				break
			}
			k1++
		}
		result = append(result, t[j])
		t[j] = '.'
		k = k - te*temp[i]
	}
	return string(result)
}

func main() {
	n := 2
	for k := 1; k < 3; k++ {
		fmt.Println(getPermutation(n, k))
	}
}
