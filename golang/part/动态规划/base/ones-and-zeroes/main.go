package ones_and_zeroes

/**
474. 一和零
https://leetcode-cn.com/problems/ones-and-zeroes/

思路： 有两个状态需要转移，一个是m，一个是n
      那么设计状态转移方程: dp[i][j] = max(1 + dp[i - cost_zero[k]][j - cost_one[k]], dp[cost_zero[i], cost_one[j]])
 */


func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m + 1)
	for i := 0;  i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}
	for _, str := range strs {
		count := countZeroAndOne(str)
		for i := m; i >= count[0]; i-- {
			for j := n; j >= count[1]; j-- {
				dp[i][j] = max(1 + dp[i - count[0]][j - count[1]], dp[i][j])
			}
		}
	}
	return dp[m][n]
}

func countZeroAndOne(s string) []int {
	result := make([]int, 2)
	for _, c := range s {
		if c == '0' {
			result[0] += 1
		} else {
			result[1] += 1
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
