# coding=utf8
"""
剑指 Offer 03. 数组中重复的数字
https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
"""
from typing import List


class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        arr = [0] * len(nums)
        for num in nums:
            arr[num] += 1
            if arr[num] > 1:
                return num
        return -1


if __name__ == '__main__':
    print(Solution().findRepeatNumber([2, 3, 1, 0, 2, 5, 3]))
