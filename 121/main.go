package main

import "fmt"

func maxProfit(prices []int) int {
	f := [][]int{}
	maxSum := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > 0 {
				f[i][j] = profit
				if profit > maxSum {
					maxSum = profit
				}
			}
		}
	}
	return maxSum
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
