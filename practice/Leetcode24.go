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

func swapPairs(head *ListNode) *ListNode {
	cur := head
	pre := head
	next := head
	for cur != nil && cur.Next != nil{
		next = cur.Next
		if cur == head {
			head = next
		}else {
			pre.Next = next
		}
		cur.Next = next.Next
		next.Next = cur
		pre = cur
		cur = cur.Next
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
	//n1.Next.Next.Next.Next.Next.Next = &ListNode{7,nil}


	printList(swapPairs(n1))
}
