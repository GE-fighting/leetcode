package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	nodeList := []*ListNode{}
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	start := &ListNode{}
	res := start
	for j := len(nodeList) - 1; j >= 0; j-- {
		nodeList[j].Next = nil
		start.Next = nodeList[j]
		start = start.Next
	}
	return res.Next
}

func NewList(nums []int) *ListNode {
	dummy := &ListNode{} // 创建虚拟头节点
	curr := dummy        // 创建当前节点指针
	for _, num := range nums {
		curr.Next = &ListNode{Val: num} // 创建新节点
		curr = curr.Next                // 将当前节点指向新节点
	}
	return dummy.Next // 返回真实头节点
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := NewList(nums)
	list := reverseList(head)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

}
