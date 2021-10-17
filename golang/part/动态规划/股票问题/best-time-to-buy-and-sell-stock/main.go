package best_time_to_buy_and_sell_stock

import "math"

/**
121. 买卖股票的最佳时机
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
思路： 由于K = 1， 所以不需要保存K值

 */

func maxProfitV1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = -prices[0]
			continue
		}
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		dp[i][1] = max(dp[i - 1][1], -prices[i])
	}
	return dp[n - 1][0]
}


func maxProfit(prices []int) int {
	// 简化DP，因为只与相邻存储有关，那么只需要保存相邻的即可
	n := len(prices)
	if n == 0 {
		return 0
	}
	dpI0 := 0
	dpI1 := math.MinInt32
	for i := 0; i < n; i++ {
		dpI0 = max(dpI0, dpI1 + prices[i])
		dpI1 = max(dpI1, -prices[i])
	}
	return dpI0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}