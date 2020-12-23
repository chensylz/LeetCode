package climbing_stairs

/**
70. 爬楼梯
https://leetcode-cn.com/problems/climbing-stairs/
 */

func climbStairs(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}