/*
题目描述
小红想知道n个“red”字符可以构成多少不同长度为3n的字符串，答案请对10E9 + 7取模，输入一个n表示几个“red”串，输出所有的结果。
输入输出示例
输入2
输出5
解释
redred
rreedd
reredd
rreded
rerded
*/

package main

import "fmt"

const MOD = int(1e9) + 7

func countString(n int) int {
	//定义问题
	//dp[i][j][k],代表放置i个r，j个e，k个d的有多少中不同的字符串
	dp := make([][][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([][]int, n+1)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	//边界条件
	dp[0][0][0] = 1
	/*
		状态转移方程
	*/
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k <= j; k++ {
				if i+1 <= n {
					dp[i+1][j][k] = (dp[i+1][j][k] + dp[i][j][k]) % MOD
				}
				if j+1 <= n {
					dp[i][j+1][k] = (dp[i][j+1][k] + dp[i][j][k]) % MOD
				}
				if k+1 <= n {
					dp[i][j][k+1] = (dp[i][j][k+1] + dp[i][j][k]) % MOD
				}
			}
		}
	}
	return dp[n][n][n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(countString(n))
}
