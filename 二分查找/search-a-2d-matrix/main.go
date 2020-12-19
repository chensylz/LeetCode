package search_a_2d_matrix

/**
74. 搜索二维矩阵
https://leetcode-cn.com/problems/search-a-2d-matrix/
 */


func searchMatrix(matrix [][]int, target int) bool {
	// 思路： 将2维数组转为1维数组
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix)
	col := len(matrix[0])

	start := 0
	end := row * col - 1

	for start + 1 < end {
		mid := start + (end - start) / 2
		// 获取1维数组对应的2维数组的坐标
		value := matrix[mid / col][mid % col]
		if value < target {
			start = mid
		} else if value > target {
			end = mid
		} else {
			return true
		}
	}

	if matrix[start / col][start % col] == target || matrix[end / col][end % col] == target {
		return true
	}
	return false
}