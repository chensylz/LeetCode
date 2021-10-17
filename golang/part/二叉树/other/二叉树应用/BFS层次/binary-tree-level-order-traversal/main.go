package binary_tree_level_order_traversal

/**
102. 二叉树的层序遍历
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([][]int, 0)

	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			// 出队列
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}
