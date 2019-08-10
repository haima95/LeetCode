package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left) + 1
	r := minDepth(root.Right) + 1
	if l == 1 && r != 1 {
		return r
	}
	if l != 1 && r == 1 {
		return l
	}
	if l > r {
		return r
	} else {
		return l
	}
}

func main() {
}
