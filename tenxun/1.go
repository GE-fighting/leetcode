package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func count(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 定义子问题
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(nums); i++ {
		add := 0
		for j := 1; j < i; j++ {
			if isPerfect(nums[i-1] * nums[j-1]) {
				add++
			}
		}
		dp[i] = dp[i-1] + add
	}
	return dp[len(nums)]

}

func isPerfect(num int) bool {
	if num <= 9 && num >= 1 {
		return true
	} else if num%10 == 0 {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	readString, _ := reader.ReadString(']')
	s := readString[1 : len(readString)-1]
	split := strings.Split(s, ",")
	nums := []int{}
	for _, v := range split {
		var num int
		fmt.Sscanf(v, "%d", &num)
		nums = append(nums, num)
	}
	fmt.Println(count(nums))
}
