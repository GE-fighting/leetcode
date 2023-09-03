package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	start := []int{}
	backtrack(nums, start, &res)
	return res
}

func backtrack(start []int, path []int, res *[][]int) {
	if len(start) == 0 {
		*res = append(*res, path)
		return
	}
	for i := 0; i < len(start); i++ {
		//做出选择
		num := start[i]
		newNums := append([]int{}, start...)
		newPath := append(path, num)
		newNums = append(newNums[:i], newNums[i+1:]...)
		backtrack(newNums, newPath, res)
	}
}
func main() {
	nums := []int{1, 2, 3}
	i := permute(nums)
	fmt.Println(i)
}
