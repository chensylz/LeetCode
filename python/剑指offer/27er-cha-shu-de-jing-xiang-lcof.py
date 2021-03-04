# coding=utf8

"""
剑指 Offer 27. 二叉树的镜像
https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
"""

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def mirrorTree(self, root: TreeNode) -> TreeNode:
        def mirror(root_node):
            if root_node is None:
                return root_node
            res = TreeNode(root.val)
            res.left = mirror(root_node.right)
            res.right = mirror(root_node.left)
            return res
        return mirror(root)


if __name__ == '__main__':
    root_t = TreeNode(4)
    root_t.left = TreeNode(2)
    root_t.right = TreeNode(7)
    Solution().mirrorTree(root_t)
