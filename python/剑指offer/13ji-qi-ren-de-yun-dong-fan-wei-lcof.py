# coding=utf8

"""
剑指 Offer 13. 机器人的运动范围
https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
思路：广度优先搜索
"""

class Solution:
    def movingCount(self, m: int, n: int, k: int) -> int:
        if k == 0:
            return 1

        from queue import Queue
        q = Queue()
        q.put((0, 0))
        s = set()
        while not q.empty():
            x, y = q.get()
            if (x, y) not in s and 0 <= x < m and 0 <= y < n and self.comp_number(x) + self.comp_number(y) <= k:
                s.add((x, y))
                for nx, ny in [(x + 1, y), (x, y + 1)]:
                    q.put((nx, ny))
        return len(s)

    def comp_number(self, x):
        total = 0
        while x:
            total += x % 10
            x //= 10
        return total
