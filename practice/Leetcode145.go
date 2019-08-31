package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	temp := []*TreeNode{root}
	for len(temp) > 0 {
		it := temp[len(temp)-1]
		if it != nil {
			if it.Left != nil {
				temp = append(temp, it.Left)
				it.Left = nil
			} else if it.Right != nil {
				temp = append(temp, it.Right)
				it.Right = nil
			} else {
				result = append(result, it.Val)
				temp = temp[:len(temp)-1]
			}

		}
	}
	return result
}

func main() {
	a := TreeNode{3, nil, nil}
	b := TreeNode{2, &a, nil}
	c := TreeNode{1, nil, &b}
	fmt.Println(postorderTraversal(&c))
}
