package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := int(math.MaxInt32)
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, res)
	r := dfs(root.Right, res)
	if *res < root.Val+r+l {
		*res = root.Val + r + l
	}
	temp := 0
	if l >= r {
		if temp < root.Val+l {
			temp = root.Val + l
		}
	} else {
		if temp < root.Val+r {
			temp = root.Val + r
		}
	}
	return temp
}
func main() {
	b := TreeNode{-1, nil, nil}
	a := TreeNode{-2, &b, nil}
	fmt.Println(maxPathSum(&a))
}
