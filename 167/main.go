package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	low := 0
	high := len(numbers) - 1
	for high > low {
		if numbers[high]+numbers[low] > target {
			high--
		}
		if numbers[high]+numbers[low] < target {
			low++
		}
		if numbers[high]+numbers[low] == target {
			break
		}
	}
	return []int{low + 1, high + 1}
}

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(twoSum(nums, 6))

}
