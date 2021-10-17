package longest_increasing_subsequence

/**
300. 最长递增子序列
https://leetcode-cn.com/problems/longest-increasing-subsequence/
 */

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}
	result := dp[0]
	for i := 1; i < n; i++ {
		result = max(result, dp[i])
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}