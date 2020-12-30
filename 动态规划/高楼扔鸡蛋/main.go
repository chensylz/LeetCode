package main

/**
887. 鸡蛋掉落
https://leetcode-cn.com/problems/super-egg-drop/
 */

func superEggDrop(K int, N int) int {
	if K == 1  {
		return N
	}
	if K == 0 {
		return 0
	}
	dp := make([][]int, N + 1)
	for i := 0; i < N + 1; i++ {
		dp[i] = make([]int, K + 1)
	}

	// 第0行与第一行
	for i := 0; i <= K; i++ {
		dp[0][i] = 0
		dp[1][i] = 1
	}
	dp[1][0] = 0
	// 第0列与第一列
	for i := 0; i <= N; i++ {
		dp[i][0] = 0
		dp[i][1] = i
	}

	// 第二行第二列开始
	for i := 2; i <= N; i++ {
		for j := 2; j <= K; j++ {
			dp[i][j] = 2 << 31
			for k := 1; k <= i; k++ {
				// 蛋没碎和蛋碎了
				// 碎了，就需要往低层继续扔：层数少 1 ，鸡蛋也少 1
				// 不碎，就需要往高层继续扔：层数是当前层到最高层的距离差，鸡蛋数量不少
				// 两种情况都做了一次尝试，所以加 1
				dp[i][j] = min(dp[i][j], max(dp[k - 1][j - 1], dp[i - k][j]) + 1)
			}
		}
	}
	return dp[N][K]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}