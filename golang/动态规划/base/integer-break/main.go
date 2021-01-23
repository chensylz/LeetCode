package main

import "fmt"

/**
343. 整数拆分
https://leetcode-cn.com/problems/integer-break/
 */

func integerBreak(n int) int {
	dp := make([]int, n + 1)
	dp[0], dp[1] = 0, 0

	for i := 2; i <= n; i++ {
		maxValue := 0
		for j := i; j > 0; j-- {
			maxValue = max(maxValue, max(j * (i - j), j * dp[i - j]))
		}
		dp[i] = maxValue
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(integerBreak(10))
}