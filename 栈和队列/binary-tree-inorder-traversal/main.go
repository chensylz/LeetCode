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
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈
		node := stack[len(stack) - 1]
		result = append(result, node.Val)
		stack = stack[:len(stack) - 1]
		root = node.Right
	}
	return result
}