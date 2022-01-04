package main

import (
	"strconv"
)

/**
887. 鸡蛋掉落
https://leetcode-cn.com/problems/super-egg-drop/

定义 dp[i][j]
在i个鸡蛋做检测j层时，最坏情况下至少抛鸡蛋的次数
dp[i][j] = min(dp[i][j], max(dp[i - 1][k - 1], dp[i][j - k]) + 1) ,  1<=k<=j
此时， dp[i - 1][k - 1] 表示在j层时鸡蛋碎了，鸡蛋i -1, 向下查找，楼层数为 k - 1
      dp[i][k - i] 表示鸡蛋没碎，继续向上查询，剩余楼层 j - k
*/

func superEggDropv1(K int, N int) int {
	if K == 1 {
		return N
	}
	if K == 0 {
		return 0
	}
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, K+1)
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
				dp[i][j] = min(dp[i][j], max(dp[k-1][j-1], dp[i-k][j])+1)
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

//////////////////////////////       V2      ////////////////////////////////

func superEggDropV2(K int, N int) int {
	memo := make(map[string]int)
	return dp(memo, K, N)
}

func dpV2(memo map[string]int, K, N int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	if value, ok := memo[strconv.Itoa(K)+","+strconv.Itoa(N)]; ok {
		return value
	}
	var result int
	result = 2 << 31
	for i := 1; i <= N; i++ {
		result = min(result, max(
			dp(memo, K-1, i-1), // 碎了
			dp(memo, K, N-i),   // 没碎
		)+1)
	}
	memo[strconv.Itoa(K)+","+strconv.Itoa(N)] = result
	return result
}

//////////////////////////////       V3       ////////////////////////////////
func superEggDrop(K int, N int) int {
	memo := make(map[string]int)
	return dp(memo, K, N)
}

func dp(memo map[string]int, K, N int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	if value, ok := memo[strconv.Itoa(K)+","+strconv.Itoa(N)]; ok {
		return value
	}
	var result int
	result = 2 << 31
	//for i := 1; i <= N; i++ {
	//	result = min(result, max(
	//		dp(memo, K - 1, i - 1), // 碎了
	//		dp(memo, K, N - i), // 没碎
	//	) + 1)
	//}
	// 使用二分搜索
	left, right := 1, N
	for left <= right {
		mid := (left + right) / 2
		broken := dp(memo, K-1, mid-1)
		noBroken := dp(memo, K, N-mid)
		if broken > noBroken {
			right = mid - 1
			result = min(result, broken+1)
		} else {
			left = mid + 1
			result = min(result, noBroken+1)
		}
	}

	memo[strconv.Itoa(K)+","+strconv.Itoa(N)] = result
	return result
}
