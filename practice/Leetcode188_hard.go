package main

import "fmt"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k == 0 || n < 2 {
		return 0
	}
	if k == 1 {
		sum := 0
		min := prices[0]
		for i := 1; i < n; i++ {
			temp := prices[i] - min
			if temp < 0 {
				min = prices[i]
			} else if temp > sum {
				sum = temp
			}
		}
		return sum
	}
	if k >= n/2 {
		sum := 0
		for i := 1; i < n; i++ {
			temp := prices[i] - prices[i-1]
			if temp > 0 {
				sum += temp
			}
		}
		return sum
	}
	t_k1, t_k0 := make([]int, k), make([]int, k)
	for i := 0; i < k; i++ {
		t_k1[i] = -prices[0]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if j > 0 {
				if t_k0[j-1]-prices[i] > t_k1[j] {
					t_k1[j] = t_k0[j-1] - prices[i]
				}
			} else {
				if -prices[i] > t_k1[j] {
					t_k1[j] = -prices[i]
				}
			}
			if t_k1[j]+prices[i] > t_k0[j] {
				t_k0[j] = t_k1[j] + prices[i]
			}
		}
	}
	return t_k0[k-1]
}

func main() {
	t := []int{3, 2, 6, 5, 0, 3}
	k := 2
	fmt.Println(maxProfit(k, t))
}
