# coding=utf8

"""
剑指 Offer 32 - I. 从上到下打印二叉树
https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
思路： 广度优先
"""
from queue import Queue
from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def levelOrder(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        q = Queue()
        res = []
        q.put(root)
        while not q.empty():
            node = q.get()
            res.append(node.val)
            if node.left:
                q.put(node.left)
            if node.right:
                q.put(node.right)
        return res


if __name__ == '__main__':
    r = TreeNode(3)
    r.left = TreeNode(9)
    r.right = TreeNode(20)
    r.right.left = TreeNode(15)
    r.right.right = TreeNode(7)
    print(Solution().levelOrder(r))