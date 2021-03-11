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
        while len(q):
            tmp_arr = collections.deque()
            for i in range(len(q)):
                node = q.popleft()
                if len(res) % 2:
                    tmp_arr.appendleft(node.val)  # 偶数层 -> 队列头部
                else:
                    tmp_arr.append(node.val)  # 奇数层 -> 队列尾部
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
            res.append(list(tmp_arr))
        return res
