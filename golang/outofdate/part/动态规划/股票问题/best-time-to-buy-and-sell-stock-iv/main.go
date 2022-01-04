package best_time_to_buy_and_sell_stock_iv

/**
188. 买卖股票的最佳时机 IV
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
 */

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k + 1)
	}

	for i := 0; i < n; i++ {
		for z := 1; z <= k; z++ {
			if i == 0 {
				dp[i][z][0] = 0
				dp[i][z][1] = -prices[i]
				continue
			}
			dp[i][z][0] = max(dp[i - 1][z][0], dp[i - 1][z][1] + prices[i])
			dp[i][z][1] = max(dp[i - 1][z][1], dp[i - 1][z - 1][0] - prices[i])
		}
	}
	return dp[n - 1][k][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}