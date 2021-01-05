package main

/**
516. 最长回文子序列
https://leetcode-cn.com/problems/longest-palindromic-subsequence/
思路； 1. 两个字符串，可划分为动态规划二维表
      2. 行为子序列，列表示原序列
      3. 转移状态方程:            dp[i + 1][j - 1] + 2, s[i + 1] == s[j - 1]
                    dp[i][j] = 丨                                           i < j
                                max(dp[i + 1][j], dp[i][j - 1])
      4. 因为状态转移需要当前格的左边，下边，下左三个方向，所以需要方向遍历， 同时每个子序列最差的情况即
         自己为一个序列，所以在初始化时，i, j相等为1， 对于i > j的位置，不存在子序列，为0
*/


func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i + 1][j - 1] + 2
			} else {
				dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[0][n - 1]
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	longestPalindromeSubseq("bbbab")
}