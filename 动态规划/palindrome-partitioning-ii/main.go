package palindrome_partitioning_ii

/**
132. 分割回文串 II
https://leetcode-cn.com/problems/palindrome-partitioning-ii/
 */

func minCut(s string) int {
	// 思路：最差的情况就是每个字符都组不成回文
	//      然后每个字符又是由它前面部分出现的回文+回文截止那个字符到它之间的次数
	//      推导公式： dp[i] = min(dp[j + 1]) && j < i && isPalindrome(j+1 ~ i)
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	dp := make([]int, len(s) + 1)
	dp[0] = -1

	for i := 1; i <= len(s); i++ {
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			if isPalindrome(s, j, i - 1) {
				dp[i] = min(dp[i], dp[j] + 1)
			}
		}
	}
	return dp[len(s)]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}


func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}