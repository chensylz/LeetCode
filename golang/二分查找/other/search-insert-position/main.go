package search_insert_position

/**
35. 搜索插入位置
https://leetcode-cn.com/problems/search-insert-position/
 */

func searchInsert(nums []int, target int) int {
	// 思路：找到第一个 >= target 的元素位置
	start := 0
	end := len(nums) - 1
	for start + 1 < end {
		mid := start + (end - start) / 2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] < target {
		return end + 1
	}

	return 0
}
