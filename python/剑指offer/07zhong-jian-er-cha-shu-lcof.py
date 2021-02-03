# coding=utf8

"""
剑指 Offer 07. 重建二叉树
https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
"""
from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if not preorder:
            return None
        root = TreeNode(preorder[0])
        if len(preorder) == 1:
            return root
        L = inorder.index(preorder[0]) + 1
        root.left = self.buildTree(preorder[1: L], inorder[:L])
        root.right = self.buildTree(preorder[L:], inorder[L:])
        return root

