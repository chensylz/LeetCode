package main

import "fmt"

/**
1314. 矩阵区域和
https://leetcode-cn.com/problems/matrix-block-sum/
 */

func matrixBlockSum(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])
	if m == 0 || n == 0 {
		return nil
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for z := i - K; z <= i + K; z++ {
				for h := j - K; h <= j + K; h++ {
					if z >=0 && h >=0 && z < m && h < n {
						result[i][j] += mat[z][h]
					}
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(matrixBlockSum([][]int{{67,64,78},{99,98,38},{82,46,46},{6,52,55},{55,99,45}}, 3))
}

