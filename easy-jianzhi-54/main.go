package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	values := []int{}
	inOrder(root, &values)
	return values[len(values)-k]
}

func inOrder(root *TreeNode, value *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, value)
	*value = append(*value, root.Val)
	inOrder(root.Right, value)
}
func main() {
	// 构建一个二叉树
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	largest := kthLargest(root, 1)
	fmt.Println(largest)
}
