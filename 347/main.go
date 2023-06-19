package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	numCount := map[int]int{}
	resNums := []int{}
	for _, v := range nums {
		if numCount[v] == 0 {
			resNums = append(resNums, v)
		}
		numCount[v] = numCount[v] + 1
	}
	// 使用冒泡排序
	for i := 0; i < len(resNums); i++ {
		for j := i + 1; j < len(resNums); j++ {
			if numCount[resNums[j]] > numCount[resNums[i]] {
				resNums[j], resNums[i] = resNums[i], resNums[j]
			}
		}
	}
	return resNums[0:k]

}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	frequent := topKFrequent(nums, k)
	fmt.Println(frequent)
}
