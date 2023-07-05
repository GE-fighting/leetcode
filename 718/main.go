package main

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	// 1、暴力破解
	//len1 := len(nums1)
	//len2 := len(nums2)
	//res := 0
	//for i := 0; i < len1; i++ {
	//	for j := 0; j < len2; j++ {
	//		if nums1[i] == nums2[j] {
	//			num := 0
	//			for i+num < len1 && j+num < len2 && nums1[i+num] == nums2[j+num] {
	//				num++
	//			}
	//			if num > res {
	//				res = num
	//			}
	//		}
	//
	//	}
	//}
	//return res
	//2、dp
	len1 := len(nums1)
	len2 := len(nums2)
	res := 0
	dp := [1000][1000]int{}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

func main() {
	num1 := []int{1, 2, 3, 2, 1}
	num2 := []int{3, 2, 1, 4, 7}
	length := findLength(num1, num2)
	fmt.Println(length)
}
