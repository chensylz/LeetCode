package main

import "fmt"

/**
123. 买卖股票的最佳时机 III
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

转移方程：  dp[i][k][0] = max(dp[i-1][k][0], dp[i - 1][k][1] + prices[i])
	      dp[i][k][1] = max(dp[i-1][k][1], dp[i - 1][k - 1][0] - prices[i])
 */

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][3][2]int, n)
	for i := 0; i < n; i++ {
		for k := 1; k <= 2; k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i - 1][k][0], dp[i - 1][k][1] + prices[i])
			dp[i][k][1] = max(dp[i - 1][k][1], dp[i - 1][k - 1][0] - prices[i])
		}
	}
	return dp[n - 1][2][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
}