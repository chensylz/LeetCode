# coding=utf8

"""
剑指 Offer 12. 矩阵中的路径
https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
思路： 深度优先搜索
"""
from typing import List


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        def dfs(i, j, k):
            if not 0 <= i < len(board) or not 0 <= j < len(board[0]) or board[i][j] != word[k]:
                return False
            if k == len(word) - 1:
                return True
            board[i][j] = '/'
            res = dfs(i + 1, j, k + 1) or dfs(i - 1, j, k + 1) or dfs(i, j + 1, k + 1) or dfs(i, j - 1, k + 1)
            board[i][j] = word[k]
            return res

        for x in range(len(board)):
            for y in range(len(board[0])):
                if dfs(x, y, 0):
                    return True
        return False
