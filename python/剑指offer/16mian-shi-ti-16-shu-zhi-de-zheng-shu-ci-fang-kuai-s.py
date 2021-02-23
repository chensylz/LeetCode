# coding=utf8

"""
剑指 Offer 16. 数值的整数次方
https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/

思路：
最简单的方法: x重复乘n次， 时间复杂度O(n)
使用快速幂运算，时间复杂度O(logn)

x^(-n) = 1/(x^n)
"""

class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 1 or x == 0:
            return x
        res = 1
        if n < 0:
            x = 1 / x
            n = -n
        while n:
            if n & 1 == 1:
                res *= x
            x *= x
            n >>= 1

        return res

