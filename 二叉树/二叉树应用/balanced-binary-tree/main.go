package main

import "GoLeetCode/DataStructure"

/**
110. 平衡二叉树
https://leetcode-cn.com/problems/balanced-binary-tree/
 */


func isBalanced(root *DataStructure.TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

func maxDepth(root *DataStructure.TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// 左边不平衡，右边不平衡，左边比右边高，右边比左边高
	if left == -1 || right == -1 || left - right > 1 || right - left > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}
	return right + 1
}