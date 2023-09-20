package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	text, _ := scanner.ReadString('}')
	strSlice := strings.Split(text[1:len(text)-1], ",")
	nums := []int{}
	for _, v := range strSlice {
		var num int
		if v == " #" {
			v = "-1"
		}
		fmt.Sscanf(v, "%d", &num)
		nums = append(nums, num)
	}
	bree := constructBree(nums)
	res := 0
	helper(bree, 0, 0, &res)
	fmt.Println(res)
}

func helper(root *TreeNode, one, zone int, ans *int) {
	if root.Val == 0 {
		zone += 1
	} else {
		one += 1
	}
	if root.Left != nil {
		helper(root.Left, one, zone, ans)
	}
	if root.Right != nil {
		helper(root.Right, one, zone, ans)
	}
	if one == zone+1 && root.Left == nil && root.Right == nil {
		*ans += 1
	}

}

// 广度优先搜索构建二叉树
func constructBree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	index := 1
	for index < len(nums) {
		node := queue[0]
		queue = queue[1:]
		if nums[index] != -1 {
			leftNode := &TreeNode{Val: nums[index]}
			queue = append(queue, leftNode)
			node.Left = leftNode
		}
		index++
		if index >= len(nums) {
			break
		}
		if nums[index] != -1 {
			rightNode := &TreeNode{Val: nums[index]}
			queue = append(queue, rightNode)
			node.Right = rightNode
		}
		index++
	}
	return root
}
