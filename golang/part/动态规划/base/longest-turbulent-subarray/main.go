package main

import "fmt"

/**
978. 最长湍流子数组
https://leetcode-cn.com/problems/longest-turbulent-subarray/
思路: dp[i][0]表示以元素i结尾，且是下降的最长长度  nums[i] < nums[i - 1]
     dp[i][1]表示以元素i结尾，且是上升的最长长度  nums[i] > nums[i - 1]
     状态转移方程: dp[i][n] = dp[i - 1][n] + 1
 */

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 1
	}
	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 1, 1
	result := 0
	for i := 1; i < n; i++ {
		dp[i][0], dp[i][1] = 1, 1
		if arr[i] > arr[i - 1] {
			dp[i][1] = dp[i - 1][0] + 1
		} else if arr[i] < arr[i - 1] {
			dp[i][0] = dp[i - 1][1] + 1
		}
		result = max(result, dp[i][0], dp[i][1])
	}
	return result
}


func max(x... int) int {
	result := x[0]
	for _, v := range x {
		if v > result {
			result = v
		}
	}
	return result
}

func main() {
	fmt.Println(maxTurbulenceSize([]int{9,4,2,10,7,8,8,1,9}))
}