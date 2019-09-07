package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fontier := &ListNode{0, head}
	pre := fontier
	cur := head
	var insertNode *ListNode
	for cur.Next != nil {
		insertNode = cur.Next
		if insertNode.Val < pre.Next.Val {
			cur.Next = insertNode.Next
			insertNode.Next = pre.Next
			pre.Next = insertNode
		} else if insertNode.Val >= cur.Val {
			cur = cur.Next
		} else {
			for pre != cur && pre.Next.Val < insertNode.Val {
				pre = pre.Next
			}
			cur.Next = insertNode.Next
			insertNode.Next = pre.Next
			pre.Next = insertNode
			pre = fontier
		}
	}
	return fontier.Next
}

func main() {
}
