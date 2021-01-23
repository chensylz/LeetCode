package main

import "fmt"

/**
837. 新21点
https://leetcode-cn.com/problems/new-21-game/
思路： dp[i] = sum(dp[i + 1], dp[i + 2] + ... + dp[i + w])
      base: dp[K+i - 1] = 1.0, k + i < N, i < W
 */

func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W)

	s := 0.0
	// 当K + i < N时， 已经获胜
	for i := 0; i < K + W; i++ {
		if K + i <= N {
			dp[K + i] = 1.0
			s += 1.0
		}
	}

	for i := K - 1; i >= 0; i-- {
		dp[i] = s / float64(W)
		s = s - dp[i + W] + dp[i]
	}

	return dp[0]
}

func main() {
	fmt.Println(new21Game(6, 1, 10))
}