package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	left := preorderTraversal(root.Left)
	if len(left) > 0 {
		result = append(result, left...)
	}
	right := preorderTraversal(root.Right)
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	temp := []*TreeNode{root}
	for len(temp) > 0 {
		it := temp[len(temp)-1]
		temp = temp[:len(temp)-1]
		if it != nil {
			result = append(result, it.Val)
			if it.Right != nil {
				temp = append(temp, it.Right)
			}
			if it.Left != nil {
				temp = append(temp, it.Left)
			}
		}
	}
	return result
}

func main() {
}
