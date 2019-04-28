package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {   // my function
	var result *ListNode
	var head *ListNode
	add := 0
	less := 0
	for ; l1!=nil && l2!=nil; {
		less = (l1.Val + l2.Val + add)%10
		add = (l1.Val + l2.Val + add)/10
		temp := &ListNode{
			Val:less,
		}
		if result == nil {
			result = temp
			head = temp
		}else{
			result.Next = temp
			result = temp

		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for ; l1!=nil; {
		less = (l1.Val + add)%10
		add = (l1.Val + add)/10
		temp := &ListNode{
			Val:less,
		}
		if result == nil {
			result = temp
			head = temp
		}else{
			result.Next = temp
			result = temp

		}
		l1 = l1.Next
	}
	for ; l2!=nil; {
		less = (l2.Val + add)%10
		add = (l2.Val + add)/10
		temp := &ListNode{
			Val:less,
		}
		if result == nil {
			result = temp
			head = temp
		}else{
			result.Next = temp
			result = temp

		}
		l2 = l2.Next
	}
	if add != 0 {
		temp := &ListNode{
			Val:add,
		}
		result.Next = temp
	}
	return head
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {   // other function
	result := &ListNode{
		Val:0,
	}
	head := result
	add := 0
	less := 0
	for ; l1!=nil || l2!=nil; {
		x := add
		if l1 != nil {
			x += l1.Val
		}
		if l2 != nil {
			x += l2.Val
		}
		less = x%10
		add = x/10
		temp := &ListNode{
			Val:less,
		}
		result.Next = temp
		result = temp
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if add != 0 {
		temp := &ListNode{
			Val:add,
		}
		result.Next = temp
	}
	return head.Next
}