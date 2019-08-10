package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	temp := head
	n := 0
	for temp != nil {
		temp = temp.Next
		n++
	}
	mid := n / 2
	temp = head
	for mid > 0 {
		temp = temp.Next
		mid--
	}
	root := &TreeNode{temp.Val, nil, nil}
	root.Left = sortedListToBST2(head, 0, n/2)
	root.Right = sortedListToBST2(temp.Next, 0, n-n/2-1)
	return root
}

func sortedListToBST2(head *ListNode, start, end int) *TreeNode {
	if head == nil || start >= end {
		return nil
	}
	temp := head
	mid := (start + end) / 2
	temp = head
	for mid > 0 {
		temp = temp.Next
		mid--
	}
	root := &TreeNode{temp.Val, nil, nil}
	root.Left = sortedListToBST2(head, start, (start+end)/2)
	root.Right = sortedListToBST2(temp.Next, 0, end-(start+end)/2-1)
	return root
}

func main() {
	a := ListNode{9, nil}
	b := ListNode{5, &a}
	c := ListNode{0, &b}
	d := ListNode{-3, &c}
	e := ListNode{-10, &d}
	root := sortedListToBST(&e)
	print(root)
}

func print(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	print(root.Left)
	print(root.Right)
}
