package _035_search_insert_position

/**
题目:
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。


思路: 二分查找
 */

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target)
}

func binarySearch(nums []int, target int) int {
	start, end := 0, len(nums) - 1
	for start <= end {
		mid := (end - start) / 2 + start
		midValue := nums[mid]
		if target == midValue {
			return mid
		} else if target > midValue {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
