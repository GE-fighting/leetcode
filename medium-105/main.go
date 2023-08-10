package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preNode []int, inNode []int) *TreeNode {
	if len(preNode) == 0 {
		return nil
	}
	root := &TreeNode{Val: preNode[0]}
	var k int
	for i, v := range inNode {
		if v == preNode[0] {
			k = i
			break
		}
	}
	if len(preNode) > 1 {
		root.Left = BuildTree(preNode[1:k+1], inNode[0:k])
	}
	if len(preNode) > k+1 {
		root.Right = BuildTree(preNode[k+1:], inNode[k+1:])
	}

	return root
}

func main() {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	tree := BuildTree(preOrder, inOrder)
	fmt.Println("%v", tree)
}
