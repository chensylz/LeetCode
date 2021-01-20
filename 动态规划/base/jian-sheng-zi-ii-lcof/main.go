package main

import (
	"fmt"
	"math/big"
)

/**
剑指 Offer 14- II. 剪绳子 II
https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/
思路： 定义dp[i]为 长度为i时，能达到的最大乘积
 状态转移方程： dp[i] = max(dp[i], max(j * dp[i - j], j * (i - j)))
 */

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]*big.Int, n + 1)
	dp[0] =big.NewInt(0)
	dp[1] = big.NewInt(1)
	dp[2] = big.NewInt(2)
	dp[3] = big.NewInt(3)
	for i := 4; i <= n; i++ {
		dp[i] = big.NewInt(0)
		for j := 1; j <= i/2; j++ {
			d := big.NewInt(1)
			//fmt.Println(i,dp[i],dp[j],dp[i-j],11)
			dp[i] = max(dp[i], d.Mul(dp[j],dp[i-j]))
		}
	}
	d := dp[n].Mod(dp[n],big.NewInt(1000000007))
	return  int(d.Int64())
}

func max(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

func main() {
	fmt.Println(cuttingRope(10))
}