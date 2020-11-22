package binary_tree_maximum_path_sum

/**
124. 二叉树中的最大路径和
https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type ResultType struct {
	SinglePath int // 单边
	MaxPath int // 最大值(单边或者单边+根节点)
}

func maxPathSum(root *TreeNode) int {
	return helper(root).MaxPath
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath: -(1 << 31),
		}
	}

	left := helper(root.Left)
	right := helper(root.Right)
	result := ResultType{}

	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath + root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath + root.Val, 0)
	}

	maxPath := max(left.MaxPath, right.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath + right.SinglePath + root.Val)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}