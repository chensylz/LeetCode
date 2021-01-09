package best_time_to_buy_and_sell_stock_with_transaction_fee

import "math"

/**
714. 买卖股票的最佳时机含手续费
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
 */

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return n
	}

	dI0 := 0
	dI1 := math.MinInt32

	for i := 0; i < n; i++ {
		temp := dI0
		dI0 = max(dI0, dI1 + prices[i])
		dI1 = max(dI1, temp - prices[i] - fee)
	}
	return dI0
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}