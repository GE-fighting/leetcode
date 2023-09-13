package main

import (
	"fmt"
	"sort"
)

func findMinSum(n int, arr []int) int {
	if n == 1 {
		return -1
	}

	sort.Ints(arr)
	sum := 0
	for _, v := range arr {
		sum += v
	}

	if arr[0] == 1 {
		return -1
	} else {
		return sum - 2*arr[0]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(findMinSum(n, arr))
}
