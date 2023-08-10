package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	i := 0
	//第一个大于根节点的元素，后面的部分为右子树部分
	for i < len(postorder)-1 {
		if postorder[i] > root {
			break
		}
		i++
	}
	//判断右子树部分是否都满足大于根元素的条件
	for j := i; j < len(postorder)-1; j++ {
		if postorder[j] < root {
			return false
		}
	}
	//判断左右子树是否满足条件
	left := true
	if i > 0 {
		left = verifyPostorder(postorder[:i])
	}
	right := true
	if i < len(postorder)-1 {
		right = verifyPostorder(postorder[i : len(postorder)-1])
	}
	return left && right

}

func main() {
	postorder := []int{4, 8, 6, 12, 16, 14, 10}
	fmt.Println(verifyPostorder(postorder))
}
