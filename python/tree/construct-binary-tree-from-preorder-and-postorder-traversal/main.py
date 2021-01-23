# coding=utf-8

"""
889. 根据前序和后序遍历构造二叉树
https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
思路： 假如原二叉树可以表示为 [1, 2, 3, 4, 5, 6, 7]
前序： [1] + [2, 4, 5] + [3, 6, 7]
后序： [4, 5, 2] + [6, 7, 3] + [1]
那么就可以知道左分支有多少个节点
我们知道左分支的头节点为 pre[1]，但它也出现在左分支的后序表示的最后
所以，左分支节点数量 就为： pre[1] = post[L - 1]， L = post.indexOf(pre[1]) + 1
"""
from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def constructFromPrePost(self, pre: List[int], post: List[int]) -> TreeNode:
        if not pre:
            return None
        root = TreeNode(pre[0])
        if len(pre) == 1:
            return root
        L = post.index(pre[1]) + 1
        root.left = self.constructFromPrePost(pre[1:L + 1], post[:L])
        root.right = self.constructFromPrePost(pre[L+1:], post[L:-1])
        return root


if __name__ == '__main__':
    Solution().constructFromPrePost([1,2,4,5,3,6,7], [4,5,2,6,7,3,1])