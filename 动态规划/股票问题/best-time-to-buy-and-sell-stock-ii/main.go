package best_time_to_buy_and_sell_stock_ii

import "math"

/**
122. 买卖股票的最佳时机 II
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
思路： 由于是更多的交易，那么 k与k-1毫无差异
转移方程： 	 dp[x][k][0] = max(dp[x - 1][k][0], dp[x - 1][k][1] + prices[x])
//                     解释： max(选择继续不买，  选择卖掉)
//          dp[x][k][1] = max(dp[x - 1][k][1], dp[x - 1][k - 1][0] - prices[x])
//                      = max(dp[x - 1][k][1], dp[x - 1][k][0] - prices[x])
//                     解释: max(选择持有，      选择买入）
 */

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dI0 := 0
	dI1 := math.MinInt32
	for i := 0; i < n; i++ {
		temp := dI0
		dI0 = max(dI0, dI1 + prices[i])
		dI1 = max(dI1, temp - prices[i])
	}
	return dI0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}