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

func partition(head *ListNode, x int) *ListNode {
	temp := &ListNode{
		0, head,
	}
	cur := head
	pre := temp
	head = temp
	var pre1 *ListNode
	for cur != nil {
		if cur.Val >= x {
			if pre1 == nil {
				pre1 = pre
			}
			pre = cur
			cur = cur.Next
		} else {
			if pre1 != nil {
				pre.Next = cur.Next
				cur.Next = pre1.Next
				pre1.Next = cur
				cur = pre.Next
				pre1 = pre1.Next
			} else {
				pre = cur
				cur = cur.Next
			}
		}
	}
	return head.Next
}

func main() {
	a := ListNode{2, nil}
	b := ListNode{2, &a}
	c := ListNode{5, &b}
	d := ListNode{4, &c}
	e := ListNode{1, &d}
	f := ListNode{1, &e}

	head := partition(&f, 3)
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}

}
