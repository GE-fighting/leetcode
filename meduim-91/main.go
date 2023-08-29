package main

import "fmt"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	//定义子问题
	//dp[i] 表示前i+1个字符的解码方法
	length := len(s)
	dp := make([]int, length)

	//边界条件
	dp[0] = 1
	//状态转移方程
	//这里分两种情况
	//s[i]为0时
	//--- 0<s[i-1]<3,dp[i] = dp[i-2]+1
	//--- s[i-1]不等于1或2的话,dp[i] = 0
	//s[i]不等于0
	//--- dp[i] = dp[i-1]+1
	//---if s[i-1] in [1,2] => dp[i-2]+1;else => dp[i] = dp[i-1]+1
	for i := 1; i < length; i++ {
		if s[i] == '0' {
			if s[i-1] > '0' && s[i-1] < '3' {
				if i < 2 {
					dp[i] = 1
				} else {
					dp[i] = dp[i-2] + 1
				}
			}
		} else {
			dp[i] = dp[i-1] + 1
			if s[i-1] > '0' && s[i-1] < '3' {
				if i < 2 {
					dp[i] = max(dp[i], 1)
				} else {
					dp[i] = max(dp[i], dp[i-2]+1)
				}

			}
		}
	}
	return dp[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "10"
	decodings := numDecodings(s)
	fmt.Println(decodings)
}
