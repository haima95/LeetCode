package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := head.Next
	left := &ListNode{}
	left1 := left
	right := &ListNode{}
	right1 := right
	for temp != nil {
		if temp.Val < head.Val {
			left1.Next = temp
			left1 = left1.Next
		} else {
			right1.Next = temp
			right1 = right1.Next
		}
		temp = temp.Next
	}
	head.Next = nil
	left1.Next = head
	right1.Next = nil
	left = sortList(left.Next)
	right = sortList(right.Next)
	head.Next = right
	return left
}

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	rightHead := slow.Next
	slow.Next = nil

	return merge(sortList(head), sortList(rightHead))
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	current := result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return result.Next
}

func main() {
	a := ListNode{4, nil}
	b := ListNode{1, &a}
	d := ListNode{2, &b}
	e := ListNode{2, &d}
	c := sortList(&e)
	for c != nil {
		fmt.Printf("%d ", c.Val)
		c = c.Next
	}
}
