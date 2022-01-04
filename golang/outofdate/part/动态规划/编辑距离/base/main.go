package main

import "fmt"

func minDistance(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	dp := make([][]int, m + 1)
	for i := 0; i < m + 1; i++ {
		dp[i] = make([]int, n + 1)
	}
	// base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i - 1] == s2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = min(
					dp[i - 1][j] + 1, // 删除
					dp[i][j - 1] + 1, // 插入
					dp[i - 1][j - 1] + 1, // 替换
				)
			}
		}
	}
	return dp[m - 1][n - 1]
}
func min(x, y, z int) int {
	minValue := 2 << 31
	values := []int{x, y, z}
	for i := range values {
		if minValue > values[i] {
			minValue = values[i]
		}
	}
	return minValue
}

func main() {
	fmt.Print(minDistance("intention", "execution"))
}