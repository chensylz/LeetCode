package number_of_islands

/**
200. 岛屿数量
https://leetcode-cn.com/problems/number-of-islands/
 */

func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}

	if grid[i][j] == '1' {
		// 标记为已访问
		grid[i][j] = 0
		return dfs(grid, i - 1, j) + dfs(grid, i, j - 1) + dfs(grid, i + 1, j) + dfs(grid, i, j + 1) + 1
	}
	return 0
}