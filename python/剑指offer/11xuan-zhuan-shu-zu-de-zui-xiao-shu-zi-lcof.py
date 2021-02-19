# coding=utf8

"""
剑指 Offer 11. 旋转数组的最小数字
https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
"""
from typing import List


class Solution:
    def minArray(self, numbers: List[int]) -> int:
        if len(numbers) == 0:
            return 0
        pre = numbers[0]
        for i in numbers[1:]:
            if pre > i:
                return i
            pre = i
        return numbers[0]


if __name__ == '__main__':
    print(Solution().minArray([2,2,2,0,1]))