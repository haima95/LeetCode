package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric2(root.Left, root.Right)
}

func isSymmetric2(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) || left.Val != right.Val {
		return false
	}
	if isSymmetric2(left.Left, right.Right) && isSymmetric2(left.Right, right.Left) {
		return true
	} else {
		return false
	}

}

func main() {
	a := TreeNode{6, nil, nil}
	b := TreeNode{6, nil, nil}
	c := TreeNode{5, nil, &b}
	d := TreeNode{5, &a, nil}
	e := TreeNode{4, &d, nil}
	f := TreeNode{4, nil, &c}
	g := TreeNode{3, &e, &f}
	fmt.Println(isSymmetric(&g))
}
