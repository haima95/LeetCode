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

func mergeKLists(lists []*ListNode) *ListNode {
	var head,temp *ListNode
	n := len(lists)
	var min,key,flag int
	for {
		min = 1<<31
		key = 0
		flag = 0
		for i:=0;i<n;i++ {
			if lists[i] == nil {
				flag ++
				continue
			}
			if lists[i].Val < min{
				min = lists[i].Val
				key = i
			}
		}
		if flag == n {
			break
		}
		if head == nil {
			head = lists[key]
			lists[key] = lists[key].Next
			temp = head
		}else {
			temp.Next = lists[key]
			lists[key] = lists[key].Next
			temp = temp.Next
		}
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
	n1 := &ListNode{6,nil}
	n1.Next = &ListNode{9,nil}
	n1.Next.Next = &ListNode{10,nil}

	n2 := &ListNode{6,nil}
	n2.Next = &ListNode{8,nil}
	n2.Next.Next = &ListNode{9,nil}

	n3 := &ListNode{1,nil}
	n3.Next = &ListNode{5,nil}

	arr := []*ListNode{n1,n2,n3}
	printList(mergeKLists(arr))
}
