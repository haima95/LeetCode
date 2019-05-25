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
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	n := 0
	for cur != nil {
		n ++
		cur = cur.Next
	}
	if n < 2 || k < 2 {
		return head
	}
	s := n
	pre := head
	tail := head
	cur = head
	next := head.Next
	for n >= k {
		t := k
		for t>1 {
			tail.Next = next.Next
			next.Next = cur
			cur = next
			next = tail.Next
			t --
		}
		if s == n {
			head = cur
		}else {
			pre.Next = cur
		}
		if n-k < k {
			break
		}
		n -= k
		pre = tail
		cur = next
		tail = next
		next = next.Next
	}
	return head
}
func printList(head *ListNode){
	for head != nil {
		fmt.Printf("%d ",head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	n1 := &ListNode{1,nil}
	n1.Next = &ListNode{2,nil}
	n1.Next.Next = &ListNode{3,nil}
	n1.Next.Next.Next = &ListNode{4,nil}
	n1.Next.Next.Next.Next = &ListNode{5,nil}
	n1.Next.Next.Next.Next.Next = &ListNode{6,nil}
	n1.Next.Next.Next.Next.Next.Next = &ListNode{7,nil}

	printList(reverseKGroup(n1,4))
}
