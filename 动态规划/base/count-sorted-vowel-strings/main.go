package main

import "fmt"

/**
1641. 统计字典序元音字符串的数目
https://leetcode-cn.com/problems/count-sorted-vowel-strings/
思路： 二维数组 dp[i][j] 表示长度为i + 1的开头为j && 0 < j < 5的字典长度
 */

func countVowelStrings(n int) int {
	dp := make([][5]int, n)
	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			for z := j; z < 5; z++ {
				dp[i][j] += dp[i - 1][z]
			}
		}
		//dp[i][0] = dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4]
		//dp[i][1] = dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4]
		//dp[i][2] = dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4]
		//dp[i][3] = dp[i - 1][3] + dp[i - 1][4]
		//dp[i][4] = dp[i - 1][4]
	}
	sum := 0
	for i := 0; i < 5; i++ {
		sum += dp[n - 1][i]
	}
	return sum
}

func main() {
	fmt.Println(countVowelStrings(33))
}
