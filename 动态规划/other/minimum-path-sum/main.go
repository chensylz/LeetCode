package minimum_path_sum

/**
64. 最小路径和
https://leetcode-cn.com/problems/minimum-path-sum/
 */

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]
	// 第一列，只能从上面一步走来
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i - 1][0] + grid[i][0]
	}
	// 第一行，只能从左边走来
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j - 1] + grid[0][j]
	}
	// 不是第一行和第一列的，可以从上或者左边来
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
		}
	}
	return dp[rows - 1][cols - 1]
}


func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}