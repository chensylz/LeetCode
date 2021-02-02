# coding=utf8

"""
剑指 Offer 05. 替换空格
https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
"""


class Solution:
    def replaceSpace(self, s: str) -> str:
        s_re = ""
        for c in s:
            if c == " ":
                s_re += "%20"
            else:
                s_re += c
        return s_re


if __name__ == '__main__':
    print(Solution().replaceSpace("  a s 1  2"))