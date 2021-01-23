package house_robber


/**
198. 打家劫舍
https://leetcode-cn.com/problems/house-robber/

状态转移方程： dp[i] = max(dp[i + 1], nums[i] + dp[i + 2])
 */

func robV1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n + 2)
	for i := n - 1; i >= 0; i-- {
		dp[i] = max(dp[i + 1], dp[i + 2] + nums[i])
	}
	return dp[0]
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dpI1 := 0
	dpI2 := 0
	result := 0
	for i := n - 1; i >= 0; i-- {
		result = max(dpI1, dpI2 + nums[i])
		dpI2 = dpI1
		dpI1 = result
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}