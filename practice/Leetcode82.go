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
	temp := head
	pre := head
	if head == nil || head.Next == nil {
		return head
	}
	first := false
	val := temp.Val
	for temp.Next != nil {
		if val == temp.Next.Val {
			if val == head.Val {
				first = true
			}
			pre.Next = temp.Next.Next
			temp = pre
		} else {
			val = temp.Next.Val
			pre = temp
			temp = temp.Next
		}
	}
	if first {
		return head.Next
	}
	return head
}

func main() {
	a := ListNode{5, nil}
	b := ListNode{4, &a}
	c := ListNode{3, &b}
	d := ListNode{3, &c}
	e := ListNode{1, &d}
	f := ListNode{1, &e}
	g := ListNode{1, &f}
	head := deleteDuplicates(&g)
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
}
