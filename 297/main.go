package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 序列化二叉树为字符串
func serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	// 递归序列化左子树和右子树
	leftStr := serialize(root.Left)
	rightStr := serialize(root.Right)
	// 拼接字符串
	return fmt.Sprintf("%d,%s,%s", root.Val, leftStr, rightStr)
}

// 反序列化字符串为二叉树
func deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()

}

func main() {
	// 构建一个二叉树
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// 序列化二叉树
	data := serialize(root)
	fmt.Println("Serialized data:", data)

	// 反序列化字符串为二叉树
	newRoot := deserialize(data)

	// 验证反序列化结果与原二叉树是否相同
	fmt.Println("Is equal:", isEqual(root, newRoot))
}

// 判断两个二叉树是否相等
func isEqual(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isEqual(p.Left, q.Left) && isEqual(p.Right, q.Right)
}
