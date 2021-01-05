package stone_game

/**
877. 石子游戏
https://leetcode-cn.com/problems/stone-game/
 */
type node struct {
	fir int
	sec int
}

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]node, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]node, n)
		dp[i][i].fir = piles[i]
		dp[i][i].sec = 0
	}

	// 斜着遍历
	for l := 2; l <= n; l++ {
		for i := 0; i <= n - l; i++ {
			j := l + i - 1
			left := piles[i] + dp[i + 1][j].sec
			right := piles[j] + dp[i][j - 1].sec
			if left > right { // 判断选择前面还是后面的值比较合适
				dp[i][j].fir = left  // 选择前面
				dp[i][j].sec = dp[i + 1][j].fir  // 那么后手就没得选了，就还是之前先选的那个
			} else {
				dp[i][j].fir = right
				dp[i][j].sec = dp[i][j - 1].fir
			}
		}
	}
	result := dp[0][n - 1].fir > dp[0][n - 1].sec
	return result
}