package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := ListNode{0, head}
	pre := head
	cur := head.Next
	for cur != nil {
		pre.Next = cur.Next
		cur.Next = temp.Next
		temp.Next = cur
		cur = pre.Next
	}
	return temp.Next
}

func main() {
}
