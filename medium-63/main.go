package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	//定义子问题
	//dp[i][j] 表示机器到(i,j)位置最多有多少种不同的路径
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	//边界条件
	for i := 0; i < m; i++ {
		if i == 0 {
			if obstacleGrid[i][0] == 1 {
				dp[i][0] = 0
				continue
			}
		} else {
			if obstacleGrid[i][0] == 1 || dp[i-1][0] == 0 {
				dp[i][0] = 0
				continue
			}
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if j == 0 {
			if obstacleGrid[0][j] == 1 {
				dp[0][j] = 0
				continue
			}
		} else {
			if obstacleGrid[0][j] == 1 || dp[0][j-1] == 0 {
				dp[0][j] = 0
				continue
			}

		}
		dp[0][j] = 1
	}
	//定义状态转移方程
	//点（i，j）可以由(i-1,j)向下移动达到，也可以由(i,j-1)向右移动得到
	//---dp[i][j] = dp[i-1][j]+dp[i][j-1]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{{0, 1, 0, 0, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	obstacles := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(obstacles)

}
