package main


func superEggDrop(K int, N int) int {
	// 定义 dp[k][m] = n
	// 其中k 是鸡蛋个数， m是可抛鸡蛋次数，n是最坏情况下可以测试的楼层数
	// ⽆论你在哪层楼扔鸡蛋，鸡蛋只可能摔碎或者没摔碎，碎了的话就测楼下，没碎的话就测楼上。
	// 无论楼上楼下，总的楼层数 = 楼上的楼层 + 楼下的楼层 + 当前层1
	// 状态转移方程： dp[k][m] = dp[k - 1][m - 1] + dp[k][m - 1] + 1

	// 初始化
	dp := make([][]int, K + 1)
	for i := 0; i < K + 1; i++ {
		dp[i] = make([]int, N + 1)
	}

	m := 0
	for dp[K][m] < N {
		m++
		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k - 1][m - 1] + dp[k][m - 1] + 1
		}
	}
	return m
}
