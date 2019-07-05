package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	temp := head
	var tail *ListNode
	n := 0
	for temp != nil {
		n++
		tail = temp
		temp = temp.Next
	}
	if n == 0 || k%n == 0 {
		return head
	}
	n = n - k%n
	temp = head
	var t *ListNode
	for n > 0 {
		t = temp
		temp = temp.Next
		n--
	}
	t.Next = nil
	tail.Next = head
	head = temp
	return head
}

func main() {
	//e := ListNode{5, nil}
	//d := ListNode{4, &e}
	//c := ListNode{3, &d}
	//b := ListNode{2, nil}
	//a := ListNode{1, &b}
	//f := ListNode{0, &a}
	f := ListNode{1, nil}
	k := 0
	head := rotateRight(&f, k)
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}
