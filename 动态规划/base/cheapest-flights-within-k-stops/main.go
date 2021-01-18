package cheapest_flights_within_k_stops

import "math"

/**
787. K 站中转内最便宜的航班
https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/
思路：
    edge = [src, dst, price]
	dp[dst][K] 表示从src经过k次中转到达dst的最便宜价格.
	状态转移：dp[dst][K] = min(dp[dst][K], dp[src][k - 1] price])
 */

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if n == 0 {
		return -1
	}
	// dp[i][k]是经过k个中转站后到达站 i 的最小费用
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, K + 1)
		for j := 0; j <= K; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	// 当前到src站，可以前往的站
	for i := 0; i < len(flights); i++ {
		if flights[i][0] == src {
			dp[flights[i][1]][0] = flights[i][2]
		}
	}

	// src到src
	for i := 0; i < K + 1; i++ {
		dp[src][i] = 0
	}


	for k := 1; k < K + 1; k++ {
		for _, flight := range flights {
			if dp[flight[0]][k - 1] != math.MaxInt32 {
				dp[flight[1]][k] = min(dp[flight[1]][k], dp[flight[0]][k - 1] + flight[2])
			}
		}
	}

	if dp[dst][K] == math.MaxInt32 {
		return -1
	}

	return dp[dst][K]
}


func min(x, y int) int  {
	if x < y {
		return x
	}
	return y
}