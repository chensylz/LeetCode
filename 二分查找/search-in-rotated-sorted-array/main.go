package search_in_rotated_sorted_array

/**
33. 搜索旋转排序数组
https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
 */

func search(nums []int, target int) int {
	// 其实是两条上升直线，分别判断在两边的情况
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start + 1 < end {
		mid := start + (end - start) / 2
		if nums[mid] == target {
			return mid
		}

		// 判断在那条上升直线
		if nums[start] < nums[mid] {
			// 左边上升直线
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			// 右边上升直线
			if nums[end] >= target && nums[mid] <= target {
				start = mid
			} else {
				end = mid
			}
		}
	}

	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}
