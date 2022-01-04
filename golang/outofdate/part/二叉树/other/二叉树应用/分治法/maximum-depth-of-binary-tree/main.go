package main

import (
	"GoLeetCode/outofdate/DataStructure"
)

/**
104. 二叉树的最大深度
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
*/

func maxDepth(root *DataStructure.TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}
