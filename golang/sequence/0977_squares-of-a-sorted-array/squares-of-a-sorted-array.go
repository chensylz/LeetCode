package _977_squares_of_a_sorted_array

import "sort"

/**
题目: 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序

example1:
	输入：nums = [-4,-1,0,3,10]
	输出：[0,1,9,16,100]
	解释：平方后，数组变为 [16,1,0,9,100]
	排序后，数组变为 [0,1,9,16,100]

解法1:
	平方后排序，时间复杂度O(nlogn)
解法2:
	找到中间值，
 */

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	for index, value := range nums {
		result[index] = value * value
	}
	sort.Ints(result)
	return result
}