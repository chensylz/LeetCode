package binary_tree_inorder_traversal

/**
94. 二叉树的中序遍历
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 */

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}


func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 弹出
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}