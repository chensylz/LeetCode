# coding=utf8

"""
剑指 Offer 38. 字符串的排列
https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
"""
from typing import List


class Solution:
    def permutation(self, s: str) -> List[str]:
        c, res = list(s), []

        def dfs(x):
            if x == len(s) - 1:
                res.append(''.join(c))
                return
            dic = set()
            for i in range(x, len(c)):
                if c[i] in dic:
                    continue
                dic.add(c[i])
                c[i], c[x] = c[x], c[i]  # 交换
                dfs(x + 1)
                c[i], c[x] = c[x], c[i]  # 交换回

        dfs(0)
        return res