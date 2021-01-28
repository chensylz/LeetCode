from typing import List

"""
724. 寻找数组的中心索引
https://leetcode-cn.com/problems/find-pivot-index/
"""

class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 0:
            return -1
        total = sum(nums)
        sum_of_nums = 0
        for i, v in enumerate(nums):
            if 2 * sum_of_nums + v == total:
                return i
            sum_of_nums += v
        return -1


if __name__ == '__main__':
    print(Solution().pivotIndex([1, 7, 3, 6, 5, 6]))
