package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesHelper(1, n)
}

func generateTreesHelper(left, right int) []*TreeNode {
	var res []*TreeNode
	if left > right || left < 0 {
		res = append(res, nil)
		return res
	}
	for i := left; i <= right; i++ {
		leftBST := generateTreesHelper(left, i-1)
		rightBST := generateTreesHelper(i+1, right)
		for _, item := range leftBST {
			for _, item2 := range rightBST {
				root := &TreeNode{Val: i}
				root.Left = item
				root.Right = item2
				res = append(res, root)
			}
		}
	}
	return res
}

func main() {
	n := 4
	r := generateTrees(n)
	for _, v := range r {
		print(v)
		fmt.Println()
	}
}

func print(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	print(root.Left)
	print(root.Right)
}
