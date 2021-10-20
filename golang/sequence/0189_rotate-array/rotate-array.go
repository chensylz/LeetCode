package _189_rotate_array

/**
题目: 189. 旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

进阶：
	尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
	你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

example1:
	输入: nums = [1,2,3,4,5,6,7], k = 3
	输出: [5,6,7,1,2,3,4]
	解释:
	向右旋转 1 步: [7,1,2,3,4,5,6]
	向右旋转 2 步: [6,7,1,2,3,4,5]
	向右旋转 3 步: [5,6,7,1,2,3,4]

思路1:
	开一个数组来存，最后赋值回去
思路2:
	翻转数组，正确结果必有 k % len(nums) 个元素在最开始，其余元素依次往后排
	原始: 1,2,3,4,5,6,7
	翻转: 7,6,5,4,3,2,1
	翻转前 k % len(nums)个: 4,5,6,7,3,2,1
	翻转后 len(nums) - (k % len(nums)): 4,5,6,7,1,2,3
 */

func rotate1(nums []int, k int)  {
	n := len(nums)
	res := make([]int, n)
	for index, value := range nums {
		res[(index + k) % n] = value
	}
	copy(nums, res)
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n / 2; i++ {
		nums[i], nums[n - 1 - i] = nums[n - 1 - i], nums[i]
	}
}
