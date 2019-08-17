package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	max := 0
	result := 0
	for i := 1; i < n; i++ {
		if prices[i]-prices[0] > max {
			max = prices[i] - prices[0]
			t := maxProfit2(prices[i+1:])
			if max+t > result {
				result = max + t
			}
		} else if prices[0] > prices[i] {
			prices[0] = prices[i]
		}
	}
	return result
}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	max := 0
	temp := prices[0]
	for i := 1; i < n; i++ {
		if prices[i]-temp > max {
			max = prices[i] - temp
		} else if temp > prices[i] {
			temp = prices[i]
		}
	}
	return max
}

func main() {
	t := []int{2, 1, 4, 5, 2, 9, 7}
	fmt.Println(maxProfit(t))
}
