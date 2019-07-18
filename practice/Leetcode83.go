package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := head
	val := head.Val
	for head.Next != nil {
		if val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			val = head.Next.Val
			head = head.Next
		}
	}
	return temp
}

func main() {
	a := ListNode{4, nil}
	b := ListNode{4, &a}
	c := ListNode{3, &b}
	d := ListNode{3, &c}
	e := ListNode{2, &d}
	f := ListNode{1, &e}
	g := ListNode{1, &f}
	head := deleteDuplicates(&g)
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
}
