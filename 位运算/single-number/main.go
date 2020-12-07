package single_number

/**
136. 只出现一次的数字
https://leetcode-cn.com/problems/single-number/
 */

func singleNumber(nums []int) int {
	// 10 ^10 == 00
	// 两个数异或就变成0
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

