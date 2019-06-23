package main

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	jump := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		for j := 1; j <= nums[i-1]; j++ {
			t := j + i
			if t > n {
				t = n
			}
			if jump[t] == 0 {
				jump[t] = jump[i] + 1
			} else if jump[t] > jump[i]+1 {
				jump[t] = jump[i] + 1
			}

		}
	}
	return jump[n]
}

func jump2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	i := 0
	jump := 0
	maxJump := jump
	for {
		jump++
		if i+nums[i] >= len(nums)-1 {
			return jump
		}
		maxCover := -1
		for j := 1; j <= nums[i]; j++ {
			if j+i+nums[j+i] > maxCover {
				maxCover = j + i + nums[j+i]
				maxJump = j + i
			}
		}
		i = maxJump
	}
}

type result struct {
	data []int
	re   int
}

func main() {
	r := []result{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{1}, 0},
		{[]int{0}, 0},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1, 3, 5, 3, 4, 7, 1, 1, 1, 1, 1}, 4},
	}
	for _, v := range r {
		if jump2(v.data) == v.re {
			fmt.Println("ok:", v)
		} else {
			fmt.Println("fail:", jump2(v.data))
		}

	}
}
