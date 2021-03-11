# coding=utf8

"""
剑指 Offer 32 - III. 从上到下打印二叉树 III
https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
"""
import collections
from typing import List

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        q = collections.deque()
        q.append(root)
        res = []
        change = 1
        while len(q):
            tmp_arr = []
            for i in range(len(q)):
                node = q.popleft()
                tmp_arr.append(node.val)
                left, right = node.left, node.right
                if change == 0:
                    if left:
                        q.append(left)
                    if right:
                        q.append(right)
                    change = 1
                else:
                    if right:
                        q.append(right)
                    if left:
                        q.append(left)
                    change = 0
            res.append(tmp_arr)
        return res