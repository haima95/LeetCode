package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	temp := []*TreeNode{root}
	n, i := 0, 0
	for n < len(temp) {
		n = len(temp)
		var val []int
		for i < n {
			if temp[i] != nil {
				val = append(val, temp[i].Val)
				temp = append(temp, temp[i].Left)
				temp = append(temp, temp[i].Right)
			}
			i++
		}
		i = n
		if len(val) > 0 {
			result = append([][]int{val}, result...)
		}
	}
	return result
}

func main() {
	a := TreeNode{15, nil, nil}
	b := TreeNode{7, nil, nil}
	c := TreeNode{20, &a, &b}
	d := TreeNode{9, nil, nil}
	e := TreeNode{3, &d, &c}
	fmt.Println(levelOrderBottom(&e))
}
