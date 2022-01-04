package _283_move_zeroes

/**
283. 移动零

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

example:
	输入: [0,1,0,3,12]
	输出: [1,3,12,0,0]
要求:
	1.必须在原数组上操作，不能拷贝额外的数组。
	2.尽量减少操作次数。
思路:
	双指针，左指针指向已排序的尾部，右指针寻找非零的值
 */

func moveZeroes(nums []int)  {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}
	for left, right := 0, 1; left < n && right < n; {
		if nums[left] == 0 && nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[left] != 0 {
			left++
		}
		right++
	}
}