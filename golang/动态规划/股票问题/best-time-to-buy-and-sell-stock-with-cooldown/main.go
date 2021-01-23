package best_time_to_buy_and_sell_stock_with_cooldown

import "math"

/**
309. 最佳买卖股票时机含冷冻期
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

转移方程  dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
         dp[i][1] = max(dp[i - 1][1], dp[i - 2][0] - prices[i])
 */


func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dI0 := 0
	dI1 := math.MinInt32
	dPre0 := 0
	for i := 0; i < n; i++ {
		temp := dI0
		dI0 = max(dI0, dI1 + prices[i])
		dI1 = max(dI1, dPre0 - prices[i])
		dPre0 = temp
	}
	return dI0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}