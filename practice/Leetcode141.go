package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	n := 0
	h := head
	for h != nil {
		n++
		h = h.Next
	}
	temp := make([]*ListNode, n)
	n = 0
	h = head
	for h != nil {
		temp[n] = h
		h = h.Next
		n++
	}
	i := 0
	for ; i < len(temp)/2; i++ {
		temp[i].Next = temp[len(temp)-1-i]
		temp[len(temp)-1-i].Next = temp[i+1]
	}
	temp[i].Next = nil
}

func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	// find the middle
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse the right part
	head2 := slow.Next
	slow.Next = nil
	var prev *ListNode
	for head2 != nil {
		temp := head2.Next
		head2.Next = prev
		prev = head2
		head2 = temp
	}
	cur1, cur2 := head, prev
	for cur2 != nil {
		next1, next2 := cur1.Next, cur2.Next
		cur1.Next, cur2.Next = cur2, next1
		cur1, cur2 = next1, next2
	}
}

func main() {
	e := ListNode{5, nil}
	d := ListNode{4, &e}
	a := ListNode{3, &d}
	b := ListNode{2, &a}
	c := &ListNode{1, &b}
	reorderList2(c)
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}

// 1->2
// 1->3->2
// 1->4->2->3
// 1->5->2->4->3
