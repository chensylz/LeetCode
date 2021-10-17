package jump_game

/**
55. 跳跃游戏
https://leetcode-cn.com/problems/jump-game/
 */

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && nums[j] + j >= i {
				dp[i] = true
			}
		}
	}
	return dp[len(nums) - 1]
}
