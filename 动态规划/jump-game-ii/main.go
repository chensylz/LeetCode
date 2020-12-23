package jump_game_ii

/**
45. 跳跃游戏 II
https://leetcode-cn.com/problems/jump-game-ii/
 */

// v2 动态规划+贪心优化
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		// 取第一个能跳到当前位置的点即可
		// 因为跳跃次数的结果集是单调递增的，所以贪心思路是正确的
		idx := 0
		for idx < n && idx + nums[idx] < i {
			idx++
		}
		dp[i] = dp[idx] + 1
	}
	return dp[n - 1]
}


func jumpV3(nums []int) int {
	n := len(nums)
	steps := 0
	maxPosition := 0
	end := 0

	for i := 0; i < n - 1; i++ {
		maxPosition = max(maxPosition, nums[i] + i)
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


func jumpv1(nums []int) int {
	f := make([]int, len(nums))
	f[0] = 0
	for i := 1; i < len(nums); i++ {
		f[i] = i
		for j := 0; j < i; j++ {
			if nums[j] + j >= i {
				f[i] = min(f[j] + 1, f[i])
			}
		}
	}
	return f[len(nums) - 1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}