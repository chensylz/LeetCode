package unique_paths_ii

/**
63. 不同路径 II
https://leetcode-cn.com/problems/unique-paths-ii/
 */

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, cols)
			}
			dp[i][j] = 1
		}
	}

	// 第一列， 只能从上边来，判断第一列情况
	for i := 1; i < rows; i++ {
		if obstacleGrid[i][0] == 1 || dp[i - 1][0] == 0 {
			dp[i][0] = 0
		}
	}
	// 第一行，只能从左边来，判断第一行的情况
	for j := 1; j < cols; j++ {
		if obstacleGrid[0][j] == 1 || dp[0][j - 1] == 0 {
			dp[0][j] = 0
		}
	}

	// 计算
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			}
		}
	}

	return dp[rows - 1][cols - 1]
}

