package _704_binary_search

/**
给定一个n个元素有序的（升序）整型数组nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

example1:
	输入: nums = [-1,0,3,5,9,12], target = 9
	输出: 4
	解释: 9 出现在 nums 中并且下标为 4

example2:
	输入: nums = [-1,0,3,5,9,12], target = 2
	输出: -1
	解释: 2 不存在 nums 中因此返回 -1

思路:
	二分查找
 */

func search(nums []int, target int) int {
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
	return -1
}
