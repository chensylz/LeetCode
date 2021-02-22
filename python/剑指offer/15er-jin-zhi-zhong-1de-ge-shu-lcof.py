# coding=utf8

"""
剑指 Offer 15. 二进制中1的个数
https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
"""

class Solution:
    def hammingWeight(self, n: int) -> int:
        count = 0
        while n:
            count += n & 1
            n >>=1
        return count


if __name__ == '__main__':
    print(Solution().hammingWeight('00000000000000000000000000001011'))