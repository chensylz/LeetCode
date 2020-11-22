package binary_tree_zigzag_level_order_traversal

/**
103. 二叉树的锯齿形层次遍历
https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([][]int, 0)
	lastVisitStatus := false

	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
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

		if lastVisitStatus {
			reverse(list)
		}
		lastVisitStatus = !lastVisitStatus

		result = append(result, list)
	}
	return result
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2;i++ {
		nums[i], nums[len(nums)- i - 1] = nums[len(nums)- i - 1], nums[i]
	}
}
