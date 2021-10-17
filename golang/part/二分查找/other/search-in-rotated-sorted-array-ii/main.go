package search_in_rotated_sorted_array_ii

/**
81. 搜索旋转排序数组 II
https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
 */

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start + 1 < end {
		for start < end && nums[start] == nums[start + 1] {
			start += 1
		}

		for start < end && nums[end] == nums[end - 1] {
			end--
		}
		mid := start + (end - start) / 2
		if nums[mid] == target {
			return true
		}

		if nums[start] < nums[mid] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[end] >= target && target >= nums[mid] {
				start = mid
			} else {
				end = mid
			}
		}
	}

	if nums[start] == target {
		return true
	}

	if nums[end] == target {
		return true
	}

	return false
}
