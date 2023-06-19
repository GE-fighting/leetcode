package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for true {
		pos := Partition(nums, left, right)
		if pos == k-1 {
			return nums[pos]
		}
		if pos > k-1 {
			right = pos - 1
		} else {
			left = pos + 1
		}
	}
	return 0
}

func Partition(value []int, left, right int) int {
	head, tail := left+1, right
	for head <= tail {
		for head <= tail && value[head] >= value[left] {
			head++
		}
		for head <= tail && value[tail] <= value[left] {
			tail--
		}
		if head <= tail {
			value[head], value[tail] = value[tail], value[head]
		}
	}
	value[left], value[tail] = value[tail], value[left]
	return tail
}
func main() {
	num := []int{3, 2, 1, 5, 6, 4}
	k := 2
	largest := findKthLargest(num, k)
	fmt.Println(largest)

}
