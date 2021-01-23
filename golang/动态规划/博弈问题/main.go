package main

type node struct {
	fir int
	sec int
}

/**
石子游戏
你和你的朋友⾯前有⼀排⽯头堆，⽤⼀个数组 piles 表⽰，piles[i] 表⽰第 i
堆⽯⼦有多少个。你们轮流拿⽯头，⼀次拿⼀堆，但是只能拿⾛最左边或者
最右边的⽯头堆。所有⽯头被拿完后，谁拥有的⽯头多，谁获胜。
⽯头的堆数可以是任意正整数，⽯头的总数也可以是任意正整数，这样就能
打破先⼿必胜的局⾯了。⽐如有三堆⽯头  piles = [1, 100, 3]  ，先⼿不管
拿 1 还是 3，能够决定胜负的 100 都会被后⼿拿⾛，后⼿会获胜。

假设两⼈都很聪明，请你设计⼀个算法，返回先⼿和后⼿的最后得分（⽯头
总数）之差。⽐如上⾯那个例⼦，先⼿能获得 4 分，后⼿会获得 100 分，你
的算法应该返回 -96。


思路：
	1. 定义状态转移数组
                 piles = [3, 9, 1, 2]
				 | 0     | 1    | 2    | 3    |
               0 | (3,0) |(3,9) |(4,9) |(11,4)|
               1 |       |(9,0) |(9,1) |(10,2)|
               2 |       |      |(1,0) |(2,1) |
               3 |       |      |      |(2,0) |
    2. 明显可得状态转移方程: dp[i][j].fir = max(piles[i] + dp[i+1][j].sec, dp[i][j-1].sec)
                         if 先手
                            dp[i][j].sec = dp[i + 1][j].fir
                         else:
                            dp[i][j].sec = dp[i][j - 1].fir
 */

func PilesGame(piles []int) int {
	n := len(piles)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return piles[0]
	}
    dp := make([][]node, n)
    for i := 0; i < n; i++ {
    	dp[i] = make([]node, n)
    	// base case
    	dp[i][i].fir = piles[i]
    	dp[i][i].sec = 0
	}

	// 斜着遍历
	for l := 2; l <= n; l++ {
		for i := 0; i <= n - l; i++ {
			j := l + i - 1
			left := piles[i] + dp[i + 1][j].sec // 先手
			right := piles[j] + dp[i][j - 1].sec // 后手
			if left > right {  // 如果先手大于后手的值
				dp[i][j].fir = left // 那么选择先手
				dp[i][j].sec = dp[i+1][j].fir
			} else { // 先手小于后手
				dp[i][j].fir = right // 那么选择后手
				dp[i][j].sec = dp[i][j - 1].fir
			}
		}
	}
	result := dp[0][n - 1]
	return result.fir - result.sec
}
