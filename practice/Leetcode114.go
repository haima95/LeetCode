package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var temp []*TreeNode
	fla(root, &temp)
	if len(temp) == 0 {
		return
	}
	root.Left = nil
	for i := 1; i < len(temp); i++ {
		if temp[i] != nil {
			temp[i].Left = nil
			root.Right = temp[i]
			root = root.Right
		}
	}
}

func fla(root *TreeNode, temp *[]*TreeNode) {
	if root == nil {
		return
	}
	*temp = append(*temp, root)
	fla(root.Left, temp)
	fla(root.Right, temp)
}

func main() {
	a := TreeNode{2, nil, nil}
	b := TreeNode{3, nil, nil}
	c := TreeNode{1, &a, &b}
	flatten(&c)
	print(&c)
}

func print(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			fmt.Println("have a left")
		}
		fmt.Println(root.Val)
		root = root.Right
	}
}
