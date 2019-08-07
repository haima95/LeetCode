package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	temp := []*TreeNode{root}
	n := 1
	s := 0
	last := 0
	for last < len(temp) {
		last = len(temp)
		var v []int
		var i int
		for i = 0; i < n; i++ {
			if temp[s+i] != nil {
				v = append(v, temp[s+i].Val)
				temp = append(temp, temp[s+i].Left)
				temp = append(temp, temp[s+i].Right)
			}
		}
		if len(v) != 0 {
			result = append(result, v)
		}
		s += i
		n = len(temp) - s
	}
	return result
}

func main() {
	h := TreeNode{2, nil, nil}
	//a := TreeNode{15, nil, nil}
	b := TreeNode{7, nil, nil}
	c := TreeNode{20, nil, &b}
	d := TreeNode{9, nil, &h}
	e := TreeNode{3, &d, &c}
	fmt.Println(levelOrder(&e))

}
