package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	temp := []*TreeNode{root}
	n := 1
	s := 0
	last := 0
	left := true
	for last < len(temp) {
		last = len(temp)
		var v []int
		var i int
		for i = n - 1; i > -1; i-- {
			if temp[s+i] != nil {
				v = append(v, temp[s+i].Val)
				if !left {
					temp = append(temp, temp[s+i].Right)
					temp = append(temp, temp[s+i].Left)
				} else {
					temp = append(temp, temp[s+i].Left)
					temp = append(temp, temp[s+i].Right)
				}
			}
		}
		left = !left
		if len(v) != 0 {
			result = append(result, v)
		}
		s += n
		n = len(temp) - s
	}
	return result
}

func main() {
	//h := TreeNode{2, nil, nil}
	a := TreeNode{5, nil, nil}
	b := TreeNode{4, nil, nil}
	c := TreeNode{3, nil, &a}
	d := TreeNode{2, &b, nil}
	e := TreeNode{1, &d, &c}
	fmt.Println(zigzagLevelOrder(&e))
}
