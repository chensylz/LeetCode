# coding=utf8

"""
334. 递增的三元子序列
https://leetcode-cn.com/problems/increasing-triplet-subsequence/
思路： 由于题目限制O(n), 和空间复杂度O(1)，这里就只能使用一遍遍历和两个变量来保存
此处使用small, mid来保存过程变量，
当有一个值小于等于small时，替换small，否则小于等于mid时，替换mid， 若大于mid，那么三元组找到
注：为什么在small被后序更小值替换后，其三元顺序也是正确的？ 原因是small就算被替换，被替换之前的值也比mid小，所以不存在顺序影响
"""
from typing import List


class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        if len(nums) < 3:
            return False
        answer = [float('inf'), float('inf')]
        for num in nums:
            if num > answer[1]:
                return True
            if num <= answer[0]:
                answer[0] = num
            elif num <= answer[1]:
                answer[1] = num
        return False


if __name__ == '__main__':
    print(Solution().increasingTriplet([1, 2, 3, 4, 5]))
