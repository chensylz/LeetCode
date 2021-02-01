# coding=utf8

"""
剑指 Offer 04. 二维数组中的查找
https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
思路： 右上角开始遍历，如果当前的值大于搜索值，那么可以往左边移动
                  如果当前的值小于搜索值，那么可以往下面一行进行移动
"""
from typing import List


class Solution:
    def findNumberIn2DArray(self, matrix: List[List[int]], target: int) -> bool:
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return False

        col = 0
        row = len(matrix[0]) - 1
        cols = len(matrix)
        while col < cols and row >= 0:
            if matrix[col][row] == target:
                return True
            elif matrix[col][row] > target:
                row -= 1
            else:
                col += 1

        return False


if __name__ == '__main__':
    print(Solution().findNumberIn2DArray([
        [1, 4, 7, 11, 15],
        [2, 5, 8, 12, 19],
        [3, 6, 9, 16, 22],
        [10, 13, 14, 17, 24],
        [18, 21, 23, 26, 30]], 5))
