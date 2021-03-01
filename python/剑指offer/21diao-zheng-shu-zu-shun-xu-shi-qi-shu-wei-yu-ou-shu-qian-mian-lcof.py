# coding=utf8

"""
剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
"""
from typing import List


class Solution:
    def exchange(self, nums: List[int]) -> List[int]:
        left, right = 0, len(nums) - 1
        while left < right:
            if nums[left] & 1 != 0:
                left += 1
                continue
            if nums[right] & 1 != 1:
                right -= 1
                continue
            nums[left], nums[right] = nums[right], nums[left]
        return nums
