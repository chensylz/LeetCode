package coin_change

/**
322. 零钱兑换
https://leetcode-cn.com/problems/coin-change/
 */

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount + 1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin] + 1)
			}
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
