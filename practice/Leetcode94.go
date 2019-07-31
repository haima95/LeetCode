package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	seachTree(root, &result)
	return result
}

func seachTree(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	seachTree(root.Left, arr)
	*arr = append(*arr, root.Val)
	seachTree(root.Right, arr)
}

func main() {
	a := TreeNode{3, nil, nil}
	b := TreeNode{2, &a, nil}
	c := TreeNode{1, nil, &b}
	fmt.Println(inorderTraversal(&c))
}
