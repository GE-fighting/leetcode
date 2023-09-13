package main

import "fmt"

func main() {
	var nums []int

	nums = append(nums, 1, 2, 3)
	printSlice(nums)
	nums1 := make([]int, len(nums), 2*cap(nums))
	copy(nums1, nums)
	printSlice(nums1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
