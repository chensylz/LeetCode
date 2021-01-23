package number_of_1_bits

/**
191. 位1的个数
https://leetcode-cn.com/problems/number-of-1-bits/
 */

func hammingWeight(num uint32) int {
	var result = 0

	for num != 0 {
		num = num & (num - 1)
		result ++
	}
	return result
}