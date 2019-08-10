package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if getHeight(root) == -1 {
		return false
	}
	return true
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getHeight(root.Left)
	if l == -1 {
		return l
	}
	r := getHeight(root.Right)
	if r == -1 {
		return -1
	}
	if l >= r && l <= r+1 {
		return l + 1
	} else if r >= l && r <= l+1 {
		return r + 1
	} else {
		return -1
	}
}

func main() {
	a := TreeNode{4, nil, nil}
	b := TreeNode{4, nil, nil}
	c := TreeNode{3, &a, &b}
	d := TreeNode{3, nil, nil}
	e := TreeNode{2, &c, &d}
	f := TreeNode{2, nil, nil}
	g := TreeNode{1, &e, &f}
	fmt.Println(isBalanced(&g))
}
