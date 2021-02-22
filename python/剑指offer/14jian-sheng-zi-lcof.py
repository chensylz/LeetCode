# coding=utf8

"""
剑指 Offer 14- I. 剪绳子
https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
"""


class Solution:
    def cuttingRope(self, n: int) -> int:
        dp = [0] * (n + 1)
        for i in range(2, n + 1):
            for j in range(1, i):
                dp[i] = max(dp[i], max(j * (i - j), j * dp[i - j]))
        return dp[n]
