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

func isValidBST(root *TreeNode) bool {
	return isValidNode(root, math.MinInt64, math.MaxInt64)
}

func isValidNode(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	v := root.Val
	return (v > min && v < max) && isValidNode(root.Left, min, v) && isValidNode(root.Right, v, max)
}

func main() {
	// false
	//e := TreeNode{6, nil, nil}
	//f := TreeNode{20, nil, nil}
	//a := TreeNode{5, nil, nil}
	//b := TreeNode{15, &e, &f}
	//c := TreeNode{10, &a, &b}

	// true
	//a := TreeNode{1, nil, nil}
	//b := TreeNode{3, nil, nil}
	//c := TreeNode{2, &a, &b}

	// false
	//e := TreeNode{3, nil, nil}
	//f := TreeNode{6, nil, nil}
	//a := TreeNode{1, nil, nil}
	//b := TreeNode{4, &e, &f}
	//c := TreeNode{5, &a, &b}

	// false
	//j := TreeNode{3, nil, nil}
	//g := TreeNode{4, nil, nil}
	//h := TreeNode{6, nil, nil}
	//e := TreeNode{0, nil, nil}
	//f := TreeNode{2, nil, &j}
	//a := TreeNode{1, &e, &f}
	//b := TreeNode{5, &g, &h}
	//c := TreeNode{3, &a, &b}

	c := TreeNode{2147483647, nil, nil}

	fmt.Println(isValidBST(&c))
}
