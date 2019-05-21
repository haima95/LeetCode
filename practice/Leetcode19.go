package main

import "fmt"

type ListNode struct {
	Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	m := 0
	for temp != nil {
		temp = temp.Next
		m ++
	}
	n = m-n
	if n == 0 {
		head = head.Next
	}else {
		temp = head
		for n > 1 {
			temp = temp.Next
			n--
		}
		temp.Next = temp.Next.Next
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
	h1 := &ListNode{1,nil}
	//h1.Next = &ListNode{2,nil}
	//h1.Next.Next =&ListNode{3,nil}
	//h1.Next.Next.Next =&ListNode{4,nil}
	//h1.Next.Next.Next.Next =&ListNode{5,nil}
	printNode(removeNthFromEnd(h1,1))
}
