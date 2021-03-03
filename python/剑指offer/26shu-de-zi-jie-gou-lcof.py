# coding=utf8

"""
剑指 Offer 26. 树的子结构
https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
"""
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSubStructure(self, A: TreeNode, B: TreeNode) -> bool:
        def is_sub(tree_node_a, tree_node_b):
            if tree_node_b is None:
                return True
            if tree_node_a is None or tree_node_a.val != tree_node_b.val:
                return False
            return is_sub(tree_node_a.left, tree_node_b.left) and is_sub(tree_node_a.right, tree_node_b.right)

        if A is None or B is None:
            return False

        return is_sub(A, B) or self.isSubStructure(A.left, B) or self.isSubStructure(A.right, B)