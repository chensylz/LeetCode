package main

import "fmt"

func MaxProfitV1(prices []int, K int) int {
	// 一共三种状态：卖出、买入、持有
	// x表示天数， k表示最多交易次数， z表示当前持有状态(0代表卖出，1代表买入）
	// 状态转移： dp[x][k][0] = max(dp[x - 1][k][0], dp[x - 1][k][1] + prices[x])
	//                     解释： max(选择继续不买，  选择卖掉)
	//          dp[x][k][1] = max(dp[x - 1][k][1], dp[x - 1][k - 1][0] - prices[x])
	//                     解释: max(选择持有，      选择买入）
	// base case: dp[-1][k][0] = 0 因为 i 是从 0 开始的，所以 i = -1 意味着还没有开始，这时候的利润当然是 0
	//            dp[-1][k][1] = -2 << 31 还没开始的时候，是不可能持有股票的，⽤负⽆穷表⽰这种不可能
	n := len(prices)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, K + 1)
	}

	for i := 0; i < n; i++ {
		for k := 1; k <= K; k++ {
			if i - 1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				 continue
			}
			dp[i][k][0] = max(dp[i - 1][k][0], dp[i - 1][k][1] + prices[i])
			dp[i][k][1] = max(dp[i - 1][k][1], dp[i - 1][k - 1][0] - prices[i])
		}
	}

	return dp[n - 1][K][0]
}

func MaxProfit(prices []int, K int) int {
	// 一共三种状态：卖出、买入、持有
	// x表示天数， k表示最多交易次数， z表示当前持有状态(0代表卖出，1代表买入）
	// 状态转移： dp[x][k][0] = max(dp[x - 1][k][0], dp[x - 1][k][1] + prices[x])
	//                     解释： max(选择继续不买，  选择卖掉)
	//          dp[x][k][1] = max(dp[x - 1][k][1], dp[x - 1][k - 1][0] - prices[x])
	//                     解释: max(选择持有，      选择买入）
	// base case: dp[-1][k][0] = 0 因为 i 是从 0 开始的，所以 i = -1 意味着还没有开始，这时候的利润当然是 0
	//            dp[-1][k][1] = -2 << 31 还没开始的时候，是不可能持有股票的，⽤负⽆穷表⽰这种不可能
	n := len(prices)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, K + 1)
	}

	for i := 0; i < n; i++ {
		for k := 1; k <= K; k++ {
			if i - 1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i - 1][k][0], dp[i - 1][k][1] + prices[i])
			dp[i][k][1] = max(dp[i - 1][k][1], dp[i - 1][k - 1][0] - prices[i])
		}
	}

	return dp[n - 1][K][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(MaxProfit([]int{1,2,3,4,5}, 2))
}