package bitwise_and_of_numbers_range

/**
201. 数字范围按位与
https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/
 */

func rangeBitwiseAnd(m int, n int) int {
	var count int
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}