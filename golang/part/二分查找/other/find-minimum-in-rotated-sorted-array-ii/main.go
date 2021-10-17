package find_minimum_in_rotated_sorted_array_ii

/**
154. 寻找旋转排序数组中的最小值 II
https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
 */

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start + 1 < end {
		// 跳过相等的值
		for start < end && nums[end] == nums[end - 1] {
			end--
		}
		for start < end && nums[start] == nums[start + 1] {
			start++
		}
		mid := start + (end - start) / 2
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid
		}
	}

	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}