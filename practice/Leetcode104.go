package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var result, height int
	getDepth(root, height, &result)
	return result
}

func getDepth(root *TreeNode, h int, max *int) {
	if root != nil {
		h++
		if h > *max {
			*max = h
		}
		getDepth(root.Left, h, max)
		getDepth(root.Right, h, max)
	}
}

func main() {
	f := TreeNode{8, nil, nil}
	a := TreeNode{15, &f, nil}
	b := TreeNode{7, nil, nil}
	c := TreeNode{9, nil, nil}
	d := TreeNode{20, &a, &b}
	e := TreeNode{3, &c, &d}
	fmt.Println(maxDepth(&e))
}
