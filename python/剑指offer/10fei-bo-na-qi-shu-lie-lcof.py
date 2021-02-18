# coding=utf8

"""
剑指 Offer 10- I. 斐波那契数列
https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
"""


class Solution:
    def fib(self, n: int) -> int:
        if n == 0:
            return 0
        if n == 1:
            return 1
        a, b = 0, 1
        res = -1
        for i in range(1, n):
            res = (a + b) % 1000000007
            a, b = b, a + b
        return res


if __name__ == '__main__':
    print(Solution().fib(5))
