package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pArr := []*TreeNode{p}
	qArr := []*TreeNode{q}
	n := 0
	for n < len(pArr) && n < len(qArr) && len(pArr) == len(qArr) {
		if pArr[n] == nil && qArr[n] == nil {
			n++
			continue
		} else if pArr[n] != nil && qArr[n] != nil && pArr[n].Val == qArr[n].Val {
			pArr = append(pArr, pArr[n].Left)
			pArr = append(pArr, pArr[n].Right)
			qArr = append(qArr, qArr[n].Left)
			qArr = append(qArr, qArr[n].Right)
			n++
		} else {
			return false
		}
	}
	if len(pArr) != len(qArr) {
		return false
	}
	return true
}

func main() {
	a := TreeNode{2, nil, nil}
	b := TreeNode{3, nil, nil}
	c := TreeNode{1, &a, &b}

	d := TreeNode{2, nil, nil}
	e := TreeNode{3, nil, nil}
	f := TreeNode{1, &d, &e}

	fmt.Println(isSameTree(&c, &f))
}
