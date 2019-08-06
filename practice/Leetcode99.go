package main

import (
	"fmt"
	"sort"
)

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

func recoverTree(root *TreeNode) {
	var temp []*TreeNode
	insearch(root, &temp)
	var val []int
	for i := 0; i < len(temp); i++ {
		val = append(val, temp[i].Val)
	}
	sort.Ints(val)
	fmt.Println(val)
	for i := 0; i < len(temp); i++ {
		temp[i].Val = val[i]
	}
}

func insearch(root *TreeNode, temp *[]*TreeNode) {
	if root == nil {
		return
	}
	insearch(root.Left, temp)
	*temp = append(*temp, root)
	insearch(root.Right, temp)
}

func main() {
	a := TreeNode{2, nil, nil}
	b := TreeNode{1, nil, nil}
	c := TreeNode{4, &a, nil}
	d := &TreeNode{3, &b, &c}
	recoverTree(d)
	printTreeNode(d)
}

func printTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTreeNode(root.Left)
	printTreeNode(root.Right)
}
