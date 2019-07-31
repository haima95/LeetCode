package main

import (
	"fmt"
)

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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head.Next == nil {
		return head
	}
	temp := &ListNode{0, head}
	pre := temp
	cur := head
	i := 1
	for i < n {
		if i >= m && i < n {
			t := pre.Next
			pre.Next = cur.Next
			cur.Next = cur.Next.Next
			pre.Next.Next = t
		} else {
			pre = cur
			cur = cur.Next
		}
		i++
	}
	return temp.Next
}

func main() {
	//a := ListNode{5, nil}
	//b := ListNode{4, &a}
	//c := ListNode{3, &b}
	d := ListNode{2, nil}
	e := &ListNode{1, &d}
	head := e
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	head = reverseBetween(e, 1, 2)
	fmt.Println()
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
}
