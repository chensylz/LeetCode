package single_number_ii

/*
137. 只出现一次的数字 II
https://leetcode-cn.com/problems/single-number-ii/
 */

func singleNumber(nums []int) int {
	var result int
	// 计算64位中，1的个数
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 统计这个数有多少个1
			sum += (nums[j] >> i) & 1
		}
		result ^= (sum % 3) << i
	}
	return result
}
