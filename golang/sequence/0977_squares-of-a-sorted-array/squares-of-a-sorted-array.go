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
	找到中间值，然后中间值的两侧平方后都是递增的，判断两侧值依次插入结果数组即可
 */

func sortedSquares1(nums []int) []int {
	result := make([]int, len(nums))
	for index, value := range nums {
		result[index] = value * value
	}
	sort.Ints(result)
	return result
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	midIndex := -1
	for i := 0; i < n && nums[i] < 0; i++ {
		midIndex = i
	}
	ans := make([]int, 0, n)
	for i, j := midIndex, midIndex + 1; i >= 0 || j < n; {
		if i < 0 {
			ans = append(ans, nums[j] * nums[j])
			j++
		} else if j == n {
			ans = append(ans, nums[i] * nums[i])
			i--
		} else if nums[i] * nums[i] < nums[j] * nums[j] {
			ans = append(ans, nums[i] * nums[i])
			i--
		} else {
			ans = append(ans, nums[j] * nums[j])
			j++
		}
	}
	return ans
}
