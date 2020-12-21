package triangle

/**
120. 三角形最小路径和
https://leetcode-cn.com/problems/triangle/
 */

func minimumTotal(triangle [][]int) int {
	return recursion(triangle, 0, 0)
}

func recursion(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}

	return min(recursion(triangle, i + 1, j), recursion(triangle, i + 1, j + 1)) + triangle[i][j]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}