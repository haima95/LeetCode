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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val > l2.Val {
		head = l2
		l2 = l2.Next
	}else {
		head = l1
		l1 = l1.Next
	}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			temp.Next = l2
			l2 = l2.Next
		}else {
			temp.Next = l1
			l1 = l1.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	}else if l2 != nil {
		temp.Next = l2
	}
	return head
}
func printNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ",head.Val)
		head = head.Next
	}
}

func main() {
	n1 := &ListNode{4,nil}
	n1.Next = &ListNode{5,nil}
	n1.Next.Next = &ListNode{7,nil}

	n2 := &ListNode{1,nil}
	n2.Next = &ListNode{2,nil}
	n2.Next.Next = &ListNode{3,nil}

	head := mergeTwoLists(n1,n2)
	printNode(head)
}
