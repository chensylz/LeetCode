package house_robber_ii

/**
213. 打家劫舍 II
https://leetcode-cn.com/problems/house-robber-ii/
 */

func rob(nums []int) int {
	// 要么开始不被抢，要么末尾那家不被抢
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robHelper(nums[0: n - 1]), robHelper(nums[1: n]))
}

func robHelper(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dI1 := 0
	dI2 := 0
	dI := 0
	for i := n - 1; i >= 0; i++ {
		dI = max(dI1, dI2 + nums[i])
		dI2 = dI1
		dI1 = dI
	}
	return dI
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}