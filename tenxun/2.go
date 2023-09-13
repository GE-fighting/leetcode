package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1e9 + 7

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(str[:len(str)-1])

	dp := make([][][][]int64, n+1)
	for i := range dp {
		dp[i] = make([][][]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = make([][]int64, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int64, n+1)
			}
		}
	}

	dp[0][0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k <= i; k++ {
				for l := 0; l <= i; l++ {
					dp[i+1][j+1][k][l] = (dp[i+1][j+1][k][l] + dp[i][j][k][l]) % MOD // 添加 'r'
					if j > k {
						dp[i+1][j][k+1][l] = (dp[i+1][j][k+1][l] + dp[i][j][k][l]) % MOD // 添加 'e'
					}
					if k > l {
						dp[i+1][j][k][l+1] = (dp[i+1][j][k][l+1] + dp[i][j][k][l]) % MOD // 添加 'd'
					}
				}
			}
		}
	}

	var ans int64 = 0
	for j := 0; j <= n; j++ {
		for k := 0; k <= n; k++ {
			for l := 0; l <= n; l++ {
				ans = (ans + dp[n][j][k][l]) % MOD
			}
		}
	}

	fmt.Println(ans)
}
