package main

import "fmt"

/**
1301. 最大得分的路径数目
https://leetcode-cn.com/problems/number-of-paths-with-max-score/
*/

func pathsWithMaxScore(board []string) []int {
	n := len(board)
	modNum := 1000000007
	dpScore := make([][]int, n + 1)
	dpPath := make([][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		dpScore[i] = make([]int, n + 1)
		dpPath[i] = make([]int, n + 1)
	}
	dpPath[n - 1][n - 1] = 1
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >=0; j-- {
			// 判断是否走得通， 当dpPath[i + 1][j] == 0, dpPath[i+1][j+1] == 0, dpPath[i][j+1] == 0表示走不通
			if board[i][j] != 'X' && (dpPath[i + 1][j] != 0 || dpPath[i + 1][j + 1] != 0 ||
				dpPath[i][j + 1] != 0) {
				maxScore := max(max(dpScore[i + 1][j], dpScore[i + 1][j + 1]), dpScore[i][j + 1])
				if maxScore == dpScore[i + 1][j] {
					dpPath[i][j] = (dpPath[i][j] + dpPath[i + 1][j]) % modNum
				}
				if maxScore == dpScore[i + 1][j + 1] {
					dpPath[i][j] = (dpPath[i][j] + dpPath[i + 1][j + 1]) % modNum
				}
				if maxScore == dpScore[i][j + 1] {
					dpPath[i][j] = (dpPath[i][j] + dpPath[i][j + 1]) % modNum
				}
				dpScore[i][j] = maxScore + int(board[i][j] - '0')
			}
		}
	}
	maxScore := 0
	if dpScore[0][0] != 0 {
		maxScore = dpScore[0][0] - ('E' - '0')
	}
	return []int{maxScore, dpPath[0][0]}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(pathsWithMaxScore([]string{"E23", "2X2", "12S"}))
}
