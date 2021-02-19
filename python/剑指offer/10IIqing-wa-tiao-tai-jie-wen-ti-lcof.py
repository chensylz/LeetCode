# coding=utf8

"""
剑指 Offer 10- II. 青蛙跳台阶问题
https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

简单的动态规划: 状态转移方程: dp[n] = dp[n - 1] + dp[n - 2]
"""


class Solution:
    def numWays(self, n: int) -> int:
        if n == 0 or n == 1:
            return 1
        a, b = 1, 1
        res = 0
        for i in range(1, n):
            res = (a + b) % 1000000007
            a, b = b, res

        return res


if __name__ == '__main__':
    print(Solution().numWays(7))
